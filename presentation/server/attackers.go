package server

import (
	"damagecalculator/domain/situation"
	"damagecalculator/domain/stats"
	"damagecalculator/usecase/attackers"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	仮想敵から受けるダメージ一覧を取得
*/
func (s *serverImpl) attackerDamages(c *gin.Context) {
	type query struct {
		Level         uint
		BaseHP        uint
		BaseAttack    uint
		BaseDefense   uint
		BaseSpAttack  uint
		BaseSpDefense uint
		BaseSpeed     uint
		Individuals   string
		Name          string
		Move          string
		Ability       string
		Nature        string
		Item          string
		Condition     string
		Weather       string
		Field         string
	}
	var q query
	c.BindQuery(&q)

	service := attackers.NewService(s.n, s.s, s.m, s.a, s.i)
	lv := stats.NewLevel(q.Level)

	defender := &situation.PokeParams{
		Name:        q.Name,
		Individuals: q.Individuals,
		BasePoints:  []uint{q.BaseHP, q.BaseAttack, q.BaseDefense, q.BaseSpAttack, q.BaseSpDefense, q.BaseSpeed},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     q.Ability,
		Item:        q.Item,
		Nature:      q.Nature,
		Condition:   q.Condition,
	}
	conditions := &situation.FieldCondition{
		Weather:      q.Weather,
		Field:        q.Field,
		HasReflector: false,
	}

	damages := service.Create(lv, defender, conditions)

	type result struct {
		Target         string  `json:"target"`
		Move           string  `json:"move"`
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
			Move:           damages.Move(),
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
