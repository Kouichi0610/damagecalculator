package item

import (
	"damagecalculator/domain/types"
)

/*
	持ち主の能力に影響を与える
	・でんきだま、しんかのきせきなど特定のポケモンに影響を与えるものは
	　UI側で判定する

	・特定のタイプの威力に補正をかける
	・すべてのタイプの威力に補正を掛ける
	・こうかばつぐんの時に威力補正を掛ける
	・能力値に補正を掛ける
	・重さに補正を掛ける
	・持ち主のタイプを変える(ひのたまプレート、エレキメモリ) 特定のポケモンに
	・晴れ、雨の影響を受けない(使う可能性低いのでひとまず除外する)

	あやしいおこう	エスパー威力1.2
	うしおのおこう	水威力1.2
	おはなのおこう	草威力1.2
	かたいいし		いわ威力1.2
	がんせきおこう	いわ威力1.2
	いのちのたま	威力1.3 (HP１/10減る)
	xxプレート		対応するタイプの威力1.2倍、アルセウスが対応するタイプに

	かるいし	重さ半分
	くろいてっきゅう	すばやさ半分、浮遊、飛行、電磁浮遊に地面タイプが当たる

	たつじんのおび	効果抜群のわざの威力1.2倍

	ちからのハチマキ		物理技の威力1.1倍

	でんきだま	ピカチュウの攻撃特攻2倍

	とつげきチョッキ	特防1.5倍(変化技が出せない)

	ノーマルジュエル		ノーマルタイプの威力1.3倍(1度使うとなくなる)

	しんかのきせき	最終進化前に持たせると防御特防1.5倍

	ばんのうがさ	はれ、雨の影響を受けない
	ふといホネ		カラカラガラガラのこうげき2倍

*/

type ItemCreator interface {
	Create(isAttacker bool) Item
}

func (s *StatsCorrectData) Create(isAttacker bool) Item {
	return nil
}

func (s *TypeCorrectData) Create(isAttacker bool) Item {
	return nil
}
func (s *PowerCorrectData) Create(isAttacker bool) Item {
	return nil
}
func (s *WeightCorrectData) Create(isAttacker bool) Item {
	return nil
}
func (s *SuperEffectiveCorrectData) Create(isAttacker bool) Item {
	return nil
}

// 能力値補正
type StatsCorrectData struct {
	Attack, Defense, SpAttack, SpDefense, Speed float64
}

// 特定のタイプの威力に補正
type TypeCorrectData struct {
	Type  types.Type
	Scale float64
}

// 効果抜群の威力補正
type SuperEffectiveCorrectData struct {
	Scale float64
}

// 威力に補正
type PowerCorrectData struct {
	Scale float64
}

// 重さ補正
type WeightCorrectData struct {
	Scale float64
}
