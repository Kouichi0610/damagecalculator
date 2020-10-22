package server

import (
	"damagecalculator/usecase/speed"
	"fmt"
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
	fmt.Printf("Level:%d\n", lv)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	gen := speed.NewGenerator()
	list := gen.Generate(uint(lv), s.s)

	res := make([]result, 0)
	for _, sp := range list {
		info := result{
			Info:  sp.Info(),
			Speed: sp.Speed(),
		}
		res = append(res, info)
	}
	c.JSON(http.StatusOK, res)
}

type result struct {
	Info  string `json:"info"`
	Speed uint   `json:"speed"`
}
