package pokeapi

import (
	"damagecalculator/domain/gender"
	"damagecalculator/domain/move"
	"damagecalculator/domain/move/accuracy"
	"damagecalculator/domain/move/attribute"
	"damagecalculator/domain/move/category"
	"damagecalculator/domain/move/count"
	"damagecalculator/domain/move/detail"
	"damagecalculator/domain/move/power"
	"damagecalculator/domain/move/target"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/types"
	"testing"

	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

func Test_Moves(t *testing.T) {
	mv := new(movesRepository)
	mf, err := mv.Get("rock-blast")
	//mf, err := mv.Get("hyper-voice")
	//mf, err := mv.Get("swallow")
	if err != nil {
		t.Error()
	}

	cnt, _ := count.NewAttackCount(2, 5)
	expect := &move.MoveFactory{
		Name:      "ロックブラスト",
		Power:     power.NewPower(25),
		Type:      types.Rock,
		Accuracy:  accuracy.NewAccuracy(90),
		Category:  category.Physical,
		Count:     cnt,
		Target:    target.Select,
		Attribute: attribute.NewAttribute(attribute.Remote, attribute.NoAttribute),
		Detail:    detail.Default,
	}

	if !reflect.DeepEqual(mf, expect) {
		t.Errorf("%v", mf)
	}
}

// 任意のポケモンから種族データを取得できること
func Test_Species(t *testing.T) {
	sp := new(speciesRepository)

	poke, err := sp.Get("ツンデツンデ")
	if err != nil {
		t.Errorf(err.Error())
	}
	expect := &species.Species{
		Name:      "ツンデツンデ",
		Stats:     stats.NewSpeciesStats(61, 131, 211, 53, 101, 13),
		Weight:    820.0,
		Gender:    gender.Unknown,
		Types:     []types.Type{types.Rock, types.Steel},
		Abilities: []string{"ビーストブースト"},
		Moves: []string{
			"tackle",
			"take-down",
			"double-edge",
			"rock-throw",
			"earthquake",
			"bide",
			"rock-slide",
			"return",
			"frustration",
			"hidden-power",
			"facade",
			"rock-tomb",
			"rock-blast",
			"gyro-ball",
			"giga-impact",
			"flash-cannon",
			"iron-head",
			"stone-edge",
			"smack-down",
			"round",
			"bulldoze",
			"infestation",
			"brutal-swing",
		},
	}

	if !reflect.DeepEqual(expect, poke) {
		t.Errorf("%v", poke)
	}

	// 存在しない
	poke, err = sp.Get("x")
	if err == nil {
		t.Error()
	}
}

func SampleItems(st, ed int) []string {
	correctCategories := []string{
		"species-specific", // 特定のポケモンに効果
		"jewels",           // ノーマルジュエルのみ
		"held-items",       // もちもの
		"type-enhancement", // おこう
		"plates",           // プレート
		"choice",           // こだわりメガネ わざの選択できなくなる系
		"bad-held-items",   // くっつきバリなどデメリットのある持ち物
		"type-protection",  // 半減実
		"in-a-pinch",       // カムラなど
	}
	// TODO:generations

	res := make([]string, 0)
	for i := st; i <= ed; i++ {
		id := strconv.Itoa(i)
		p, err := pokeapi.Item(id)
		if err != nil {
			break
		}
		check := false
		for _, s := range correctCategories {
			if s == p.Category.Name {
				check = true
				break
			}
		}
		if !check {
			continue
		}

		s := new(strings.Builder)

		s.WriteString(fmt.Sprintf("Name:%s ", p.Names[0].Name))
		// 用途？
		s.WriteString("Attributes:[")
		for _, a := range p.Attributes {
			s.WriteString(fmt.Sprintf("%s ", a.Name))
		}
		s.WriteString("] ")
		s.WriteString(fmt.Sprintf("Baby:[%v] ", p.BabyTriggerFor))
		s.WriteString(fmt.Sprintf("Category:[%s] ", p.Category.Name))

		/*
			s.WriteString("EffectEntries:[")
			for i, a := range p.EffectEntries {
				s.WriteString(fmt.Sprintf("%d %s ", i, a.Effect))
			}
			s.WriteString("] ")
		*/

		// どのバージョンでつかわれているか
		s.WriteString("GameIndices[")
		for _, a := range p.GameIndices {
			s.WriteString(fmt.Sprintf("%s ", a.Generation.Name))
		}
		s.WriteString("]")
		// 不明。データなし
		s.WriteString("HeldBy[")
		for _, a := range p.HeldByPokemon {
			fmt.Sprintf("%v ", a)
		}
		s.WriteString("] ")

		res = append(res, s.String())
	}
	return res
}
