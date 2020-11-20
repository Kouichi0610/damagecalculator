package server

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/species"
	"net/http"
	"strconv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Serve() (http.Handler, error)
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

func (s *serverImpl) Serve() (http.Handler, error) {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	router.GET("/get_names", s.getNames)
	router.GET("/get_sample", s.getSample)
	router.POST("/post_sample", s.postSample)
	router.GET("/filtered_list", s.filteredList)
	router.GET("/speed_list", s.speedList)
	router.GET("/get_species", s.getSpecies)
	router.GET("/nature_list", s.natureList)
	router.GET("/stats_pattern", s.statsPattern)
	router.GET("/ability_owner_speed", s.abilitiesOwnerSpeedEffect)
	router.GET("/ability_other_speed", s.abilitiesOtherSpeedEffect)
	router.GET("/defender_damages", s.defenderDamages)
	router.GET("/attacker_damages", s.attackerDamages)
	router.GET("/moves", s.moves)

	return router, router.Run(":8080")
}

func (s *serverImpl) getNames(c *gin.Context) {
	n := s.n.Get()
	c.JSON(http.StatusOK, n)
}

type Sample struct {
	Name      string `json:"name"`
	HP        uint   `json:"hp"`
	Attack    uint   `json:"attack"`
	Defense   uint   `json:"defense"`
	SpAttack  uint   `json:"sp_attack"`
	SpDefense uint   `json:"sp_defense"`
	Speed     uint   `json:"speed"`
}

// テスト
func (s *serverImpl) getSample(c *gin.Context) {
	type query struct {
		Name string
	}
	//name := "ツンデツンデ"
	var q query
	c.BindQuery(&q)
	sp, _ := s.s.Get(q.Name)
	res := make([]Sample, 0)
	res = append(res, Sample{
		Name:      q.Name,
		HP:        sp.HP,
		Attack:    sp.Attack,
		Defense:   sp.Defense,
		SpAttack:  sp.SpAttack,
		SpDefense: sp.SpDefense,
		Speed:     sp.Speed,
	})

	c.JSON(http.StatusOK, res)
}

// POSTは文字列型しか受け付けない？
type postStruct struct {
	X string `json:"x"`
	Y string `json:"y"`
}

func (s *serverImpl) postSample(c *gin.Context) {
	var request postStruct
	c.BindJSON(&request)

	x, err := strconv.Atoi(request.X)
	if err != nil {
		x = 0
	}
	y, err := strconv.Atoi(request.Y)
	if err != nil {
		y = 0
	}

	c.JSON(http.StatusOK, x+y)
}
