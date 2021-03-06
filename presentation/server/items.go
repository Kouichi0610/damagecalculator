package server

import (
	"damagecalculator/domain/usableItem"
	"net/http"

	"github.com/gin-gonic/gin"
)

// もちもの一覧の取得
func (s *serverImpl) usableItems(c *gin.Context) {
	type query struct {
		Name string
	}
	var q query
	c.BindQuery(&q)

	type result struct {
		Name string `json:"name"`
		Desc string `json:"description"`
	}
	// TODO:
	sv := usableItem.NewService(s.i)
	list := sv.List(q.Name)

	res := make([]*result, 0)
	for _, item := range list {
		res = append(res, &result{Name: item.Name(), Desc: item.Description()})
	}

	c.JSON(http.StatusOK, res)
}
