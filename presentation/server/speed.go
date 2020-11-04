package server

import (
	"damagecalculator/domain/ability"
	"damagecalculator/usecase/speed"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
func (s *serverImpl) abilitiesOwnerSpeedEffect(c *gin.Context) {
	type query struct {
		ability string
	}
	var q query
	c.BindQuery(&q)

	gen := ability.NewOwnerSpeedCorrectorGenerator()
	res := gen.Create(q.ability)
	c.JSON(http.StatusOK, res)
}

// とくせいによる他のポケモンへのすばやさ補正
func (s *serverImpl) abilitiesOtherSpeedEffect(c *gin.Context) {
	type query struct {
		ability string
	}
	var q query
	c.BindQuery(&q)

	gen := ability.NewOtherSpeedCorrectorGenerator()
	res := gen.Create(q.ability)
	c.JSON(http.StatusOK, res)
}

type speedResult struct {
	Info  string `json:"info"`
	Speed uint   `json:"speed"`
}
