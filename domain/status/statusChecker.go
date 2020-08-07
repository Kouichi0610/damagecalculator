package status

import (
	"damagecalculator/domain/stats"
	"damagecalculator/domain/types"
)

/*
	ポケモンの状態
	・タイプ
	・Stats(能力値&ランク)
	・重さ
	TODO:
	・持ち物
	・特性
*/
type StatusChecker interface {
	Level() stats.Level
	Types() *types.Types
	HP() *HP
	Attack() *RankedValue
	Defense() *RankedValue
	SpAttack() *RankedValue
	SpDefense() *RankedValue
	Speed() *RankedValue
}
