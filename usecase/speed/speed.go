/*
	速度調整リスト作成
	速度調整のための仮想敵リストを作成
	TODO:手動でデータベースにしたほうが良いと思う
*/
package speed

import (
	"damagecalculator/domain/species"
	"sort"
)

/*
	種族値上限下限フィルタ(きりが無いので)
	素早さ調整リスト(補正あり、無振り、+4)
		種族値60以上なら
		未満なら無振り

	TODO:最速補正あり 無振り　+4
	　　　　(60以上)

	TODO:すいすい、ようりょくそを考慮していない
*/
type (
	// 種族値が同じポケモンをまとめて
	SpeedInfo interface {
		Info() string // names, 種族値, 基礎ポイント
		Speed() uint  // 素早さ実数
	}
	SpeedInfos []SpeedInfo

	ListGenerator interface {
		Generate(level uint, s species.Repository) SpeedInfos
	}
)

func NewGenerator() ListGenerator {
	return new(generatorImpl)
}

type generatorImpl struct {
}

/*

	作るだけで結構重い-> Vue側で起動時に回しておくか
	トリックルーム検証必要
*/
func (g *generatorImpl) Generate(level uint, s species.Repository) SpeedInfos {
	grp := newSpeedGroups(s)
	res := grp.createInfos(level)
	sort.SliceStable(res, func(i, j int) bool { return res[i].Speed() > res[j].Speed() })
	return res
}
