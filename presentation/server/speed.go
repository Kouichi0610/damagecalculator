package server

import (
	"damagecalculator/domain/ability"
	DOMAIN "damagecalculator/domain/speed"
	"damagecalculator/usecase/speed"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *serverImpl) getSpeed(c *gin.Context) {
	/*
		Name        string
		Level       uint
		Individuals string
		BasePoint   uint
		Ability     string
		Nature      string
		Item        string
		Condition   string
		Weather     string
		Field       string
	*/
	var q DOMAIN.ServiceArgs
	c.BindQuery(&q)

	sv := DOMAIN.NewService(s.s, s.a, s.m, s.i)
	speed := sv.Speed(&q)
	type result struct {
		Speed uint `json:"speed"`
	}
	res := &result{Speed: speed}
	c.JSON(http.StatusOK, res)
}

// 種族速度一覧の取得
func (s *serverImpl) speedList(c *gin.Context) {
	type query struct {
		Level string
	}
	var q query
	c.BindQuery(&q)

	lv, err := strconv.Atoi(q.Level)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	gen := speed.NewGenerator()
	list := gen.Generate(uint(lv), s.s)

	res := make([]speedResult, 0)
	for _, sp := range list {
		info := speedResult{
			Info:  sp.Info(),
			Speed: sp.Speed(),
		}
		res = append(res, info)
	}
	c.JSON(http.StatusOK, res)
}

// とくせいによる自身のすばやさ補正
// TODO:getSpeedに委譲するためこちらは削除
func (s *serverImpl) abilitiesOwnerSpeedEffect(c *gin.Context) {
	type query struct {
		Ability string
	}
	var q query
	c.BindQuery(&q)

	gen := ability.NewOwnerSpeedCorrectorGenerator()
	res := gen.Create(q.Ability)
	c.JSON(http.StatusOK, res)
}

// とくせいによる他のポケモンへのすばやさ補正
func (s *serverImpl) abilitiesOtherSpeedEffect(c *gin.Context) {
	type query struct {
		Ability string
	}
	var q query
	c.BindQuery(&q)

	gen := ability.NewOtherSpeedCorrectorGenerator()
	res := gen.Create(q.Ability)
	c.JSON(http.StatusOK, res)
}

type speedResult struct {
	Info  string `json:"info"`
	Speed uint   `json:"speed"`
}
