package server

import (
	"damagecalculator/usecase/species"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *serverImpl) getSpecies(c *gin.Context) {
	type query struct {
		Name string
	}
	var q query
	c.BindQuery(&q)

	loader := species.NewLoader(s.s)
	sp, err := loader.Get(q.Name)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	type result struct {
		Name      string   `json:name`
		Types     []string `json:types`
		Species   []uint   `json:species`
		Weight    float64  `json:weight`
		Abilities []string `json:abilities`
	}
	res := result{
		Name:      sp.Name(),
		Types:     sp.Types(),
		Species:   sp.Species(),
		Weight:    sp.Weight(),
		Abilities: sp.Abilities(),
	}
	c.JSON(http.StatusOK, res)
}
