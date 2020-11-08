package server

import (
	"damagecalculator/domain/stats"
	"damagecalculator/usecase/statspattern"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *serverImpl) statsPattern(c *gin.Context) {
	type query struct {
		Level      int
		Name       string
		Nature     string
		Individual string
	}
	var q query
	c.BindQuery(&q)

	individual := stats.ToIndividualType(q.Individual)

	ld := statspattern.NewLoader(s.s)
	r, err := ld.Get(q.Level, q.Name, q.Nature, individual)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	type result struct {
		HP        []uint `json:"hp"`
		Attack    []uint `json:"attack"`
		Defense   []uint `json:"defense"`
		SpAttack  []uint `json:"sp_attack"`
		SpDefense []uint `json:"sp_defense"`
		Speed     []uint `json:"speed"`
	}

	res := result{
		HP:        r.HP(),
		Attack:    r.Attack(),
		Defense:   r.Defense(),
		SpAttack:  r.SpAttack(),
		SpDefense: r.SpDefense(),
		Speed:     r.Speed(),
	}

	c.JSON(http.StatusOK, res)
}
