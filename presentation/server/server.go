package server

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/species"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Serve() error
}

func NewServer(n pokenames.Repository, s species.Repository, m move.Repository, a ability.Repository, i item.Repository) Server {
	return &serverImpl{
		n: n,
		s: s,
		m: m,
		a: a,
		i: i,
	}
}

type serverImpl struct {
	n pokenames.Repository
	s species.Repository
	m move.Repository
	a ability.Repository
	i item.Repository
}

func (s *serverImpl) Serve() error {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	router.GET("/names", s.names)

	return router.Run(":8080")
}

/*
	名前一覧を返す
	TODO:入力から候補一覧を取得
		ex)f -> フシギダネ、フシギソウ、フシギバナ...
*/
func (s *serverImpl) names(c *gin.Context) {
	n := s.n.Get()

	c.JSON(200, n)
}
