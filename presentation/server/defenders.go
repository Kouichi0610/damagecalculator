package server

import (
	"damagecalculator/domain/situation"
	"damagecalculator/domain/stats"
	"damagecalculator/usecase/defenders"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	ダメージ計算

*/
func (s *serverImpl) defenderDamages(c *gin.Context) {
	type query struct {
		Level       uint
		BasePoints  []uint
		Individuals string
		Name        string
		Move        string
		Ability     string
		Nature      string
		Item        string
		Condition   string
	}
	var q query
	c.BindQuery(&q)

	service := defenders.NewService(s.n, s.s, s.m, s.a, s.i)
	lv := stats.NewLevel(q.Level)

	attacker := &situation.PokeParams{
		Name:        q.Name,
		Individuals: q.Individuals,
		BasePoints:  q.BasePoints,
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     q.Ability,
		Item:        q.Item,
		Nature:      q.Nature,
		Condition:   q.Condition,
	}

	damages := service.Create(lv, attacker, q.Move, nil)

	type result struct {
		Target         string  `json:"target"`
		DamageMin      uint    `json:"damage_min"`
		DamageMax      uint    `json:"damage_max"`
		RateMin        float64 `json:"rate_min"`
		RateMax        float64 `json:"rate_max"`
		DetermineCount uint    `json:"determine_count"`
	}
	res := make([]result, 0)

	for _, damages := range damages {
		r := result{
			Target:         damages.Target(),
			DamageMin:      damages.DamageMin(),
			DamageMax:      damages.DamageMax(),
			RateMin:        damages.RateMin(),
			RateMax:        damages.RateMax(),
			DetermineCount: damages.DetermineCount(),
		}
		res = append(res, r)
	}

	c.JSON(http.StatusOK, res)
}
