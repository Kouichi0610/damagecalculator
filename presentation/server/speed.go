package server

import (
	"damagecalculator/usecase/speed"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 種族速度一覧の取得
func (s *serverImpl) speedList(c *gin.Context) {
	gen := speed.NewGenerator()
	list := gen.Generate(s.s)

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
