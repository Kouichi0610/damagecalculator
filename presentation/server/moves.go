package server

import (
	"damagecalculator/usecase/moves"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *serverImpl) moves(c *gin.Context) {
	type query struct {
		Name string
	}
	var q query
	c.BindQuery(&q)

	loader := moves.NewLoader(s.s, s.m)
	res := loader.Moves(q.Name)
	c.JSON(http.StatusOK, res)
}
