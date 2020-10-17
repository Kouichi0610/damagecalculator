package server

import (
	"damagecalculator/usecase/filter"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO:テスト環境作る

/*
	params
	filtered_list?types="ほのお"&total="600

	returns
	json:[{"name":"カイリュー","types":"ドラゴン/ひこう","hp":91,"attack":134,"defense":95,"sp_attack":100,"sp_defense":100,"speed":80}, ...]
*/
func (s *serverImpl) filteredList(c *gin.Context) {
	type query struct {
		Types string
		Total string
	}
	var q query
	c.BindQuery(&q)
	f := filter.NewFilter(strToFilterType(q.Types), q.Total)
	list := f.Filter(s.s)

	// TODO:json 返していない
	res := make([]filterTypeResult, 0)
	for _, r := range list {
		d := filterTypeResult{
			Name:      r.Name(),
			Types:     r.Types(),
			HP:        r.HP(),
			Attack:    r.Attack(),
			Defense:   r.Defense(),
			SpAttack:  r.SpAttack(),
			SpDefense: r.SpDefense(),
			Speed:     r.Speed(),
		}
		res = append(res, d)
	}
	c.JSON(http.StatusOK, res)
}

type filterTypeResult struct {
	Name      string `json:"name"`
	Types     string `json:"types"`
	HP        uint   `json:"hp"`
	Attack    uint   `json:"attack"`
	Defense   uint   `json:"defense"`
	SpAttack  uint   `json:"sp_attack"`
	SpDefense uint   `json:"sp_defense"`
	Speed     uint   `json:"speed"`
}

func strToFilterType(s string) filter.Type {
	res, ok := typesMap[s]
	if !ok {
		return filter.All
	}
	return res
}

var typesMap map[string]filter.Type

func init() {
	typesMap = make(map[string]filter.Type, 0)
	typesMap["ノーマル"] = filter.Normal
	typesMap["ほのお"] = filter.Fire
	typesMap["みず"] = filter.Water
	typesMap["でんき"] = filter.Electric
	typesMap["くさ"] = filter.Grass
	typesMap["こおり"] = filter.Ice
	typesMap["かくとう"] = filter.Fighting
	typesMap["どく"] = filter.Poison
	typesMap["じめん"] = filter.Ground
	typesMap["ひこう"] = filter.Flying
	typesMap["エスパー"] = filter.Psychic
	typesMap["むし"] = filter.Bug
	typesMap["いわ"] = filter.Rock
	typesMap["ゴースト"] = filter.Ghost
	typesMap["ドラゴン"] = filter.Dragon
	typesMap["あく"] = filter.Dark
	typesMap["はがね"] = filter.Steel
	typesMap["フェアリー"] = filter.Fairy
}
