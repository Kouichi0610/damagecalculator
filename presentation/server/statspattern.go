package server

// TODO:statslist api

import (
	"damagecalculator/usecase/statspattern"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO:実装(ORMできないか確認)

// level int, name string, nature string, hp int, at int, df int, sa int, sd int, sp int
func (s *serverImpl) statsPattern(c *gin.Context) {
	type query struct {
		Level     int
		Name      string
		Nature    string
		HP        int
		Attack    int
		Defense   int
		SpAttack  int
		SpDefense int
		Speed     int
	}
	var q query
	c.BindQuery(&q)

	ld := statspattern.NewLoader(s.s)
	r, err := ld.Get(q.Level, q.Name, q.Nature, q.HP, q.Attack, q.Defense, q.SpAttack, q.SpDefense, q.Speed)
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
