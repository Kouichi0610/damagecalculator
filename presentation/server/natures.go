package server

import (
	"damagecalculator/domain/stats"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 性格一覧の取得
func (s *serverImpl) natureList(c *gin.Context) {
	list := stats.NatureDescriptions()

	type result struct {
		Name string `json:"name"`
		Desc string `json:"description"`
	}
	res := make([]result, 0)
	for _, n := range list {
		res = append(res, result{Name: n.Name, Desc: n.Desc})
	}
	c.JSON(http.StatusOK, res)
}
