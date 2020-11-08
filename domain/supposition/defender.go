package supposition

import (
	"damagecalculator/domain/ability"
	_ "damagecalculator/domain/condition"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	_ "damagecalculator/domain/stats"
	"damagecalculator/domain/types"
	_ "damagecalculator/domain/types"
)

/*
	攻撃側の名前と技から
	防御側の仮想敵一覧を作成

	TODO:種族値合計などから良い感じのリストを取る(とくせい、もちもの)
	TODO:か、データベースに登録
*/

// TODO:PokeParamsに置き換え
type (
	DefendersMaker interface {
		Defenders(attacker string, move string) ([]situation.PokeParams, error)
	}
)

func NewDefendersMaker(nm pokenames.Repository, sp species.Repository, mv move.Repository, ab ability.Repository, it item.Repository) DefendersMaker {
	return &dmaker{
		nm: nm,
		sp: sp,
		mv: mv,
		ab: ab,
		it: it,
	}
}

type dmaker struct {
	nm pokenames.Repository
	sp species.Repository
	mv move.Repository
	ab ability.Repository
	it item.Repository
}

func (d *dmaker) Defenders(attacker string, move string) ([]situation.PokeParams, error) {
	res := make([]situation.PokeParams, 0)

	mv, err := d.mv.Get(move)
	if err != nil {
		return nil, err
	}

	for _, df := range defenders {
		// わざが有効なタイプならリストに追加
		sp, err := d.sp.Get(df.Name)
		if err != nil {
			return nil, err
		}
		dTypes := types.NewTypes(sp.Types...)
		aTypes := types.NewTypes(mv.Type)
		ef := aTypes.Magnification(dTypes)
		if ef.IsNoEffective() || ef.IsNotVery() {
			continue
		}
		data := df
		data.Individuals = "max"
		data.Ranks = []int{0, 0, 0, 0, 0}
		data.Item = ""
		data.Condition = ""
		res = append(res, data)
	}
	return res, nil
}

var defenders []situation.PokeParams

func init() {
	defenders = []situation.PokeParams{
		{Name: "カビゴン", BasePoints: []uint{252, 0, 252, 0, 0, 0}, Nature: "ずぶとい", Ability: "あついしぼう"},
		{Name: "ハピナス", BasePoints: []uint{252, 0, 6, 0, 252, 0}, Nature: "おだやか", Ability: "てんのめぐみ"},
		{Name: "シャンデラ", BasePoints: []uint{252, 0, 6, 252, 0, 0}, Nature: "ひかえめ", Ability: "もらいび"},
		{Name: "ドヒドイデ", BasePoints: []uint{252, 0, 252, 0, 6, 0}, Nature: "わんぱく", Ability: "さいせいりょく"},
		{Name: "ウォッシュロトム", BasePoints: []uint{252, 0, 0, 252, 6, 0}, Nature: "ひかえめ", Ability: "ふゆう"},
		{Name: "フシギバナ", BasePoints: []uint{252, 0, 0, 252, 6, 0}, Nature: "ひかえめ", Ability: "しんりょく"},
		{Name: "ユキノオー", BasePoints: []uint{252, 252, 0, 0, 6, 0}, Nature: "いじっぱり", Ability: "ゆきふらし"},
		{Name: "ジャラランガ", BasePoints: []uint{252, 252, 0, 0, 6, 0}, Nature: "いじっぱり", Ability: "ぼうだん"},
		{Name: "マタドガス", BasePoints: []uint{252, 0, 252, 0, 6, 0}, Nature: "わんぱく", Ability: "ふゆう"},
		{Name: "カバルドン", BasePoints: []uint{252, 0, 252, 0, 6, 0}, Nature: "わんぱく", Ability: "すなおこし"},
		{Name: "マンタイン", BasePoints: []uint{252, 0, 6, 0, 252, 0}, Nature: "おだやか", Ability: "ちょすい"},
		{Name: "クレセリア", BasePoints: []uint{252, 0, 6, 252, 0, 0}, Nature: "ひかえめ", Ability: "ふゆう"},
		{Name: "ストライク", BasePoints: []uint{6, 252, 0, 0, 0, 252}, Nature: "ようき", Ability: "むしのしらせ"},
		{Name: "ギガイアス", BasePoints: []uint{252, 252, 6, 0, 0, 0}, Nature: "ゆうかん", Ability: "がんじょう"},
		{Name: "ゲンガー", BasePoints: []uint{6, 0, 0, 252, 0, 252}, Nature: "おくびょう", Ability: "のろわれボディ"},
		{Name: "ガブリアス", BasePoints: []uint{6, 252, 0, 0, 0, 252}, Nature: "ようき", Ability: "すながくれ"},
		{Name: "バンギラス", BasePoints: []uint{252, 252, 6, 0, 0, 0}, Nature: "ゆうかん", Ability: "すなおこし"},
		{Name: "ジバコイル", BasePoints: []uint{252, 0, 6, 252, 0, 0}, Nature: "ひかえめ", Ability: "がんじょう"},
		{Name: "ニンフィア", BasePoints: []uint{252, 0, 6, 252, 0, 0}, Nature: "ひかえめ", Ability: "フェアリースキン"},
	}
}
