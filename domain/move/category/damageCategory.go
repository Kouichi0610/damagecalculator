package category

import (
	"damagecalculator/domain/status"
	"fmt"
)

// 物理、特殊などダメージ計算の攻防パラメータを引き出す
type DamageCategory uint

const (
	Physical DamageCategory = iota
	Special
	PsycoShock
	BodyPress
	FoulPlay
	ShellArms
)

func List() []DamageCategory {
	return []DamageCategory{
		Physical,
		Special,
		PsycoShock,
		BodyPress,
		FoulPlay,
		ShellArms,
	}
}

func (d DamageCategory) IsPhysical() bool {
	switch d {
	case Physical:
		return true
	case BodyPress:
		return true
	case FoulPlay:
		return true
	}
	return false
}

func (d DamageCategory) String() string {
	switch d {
	case Physical:
		return "ぶつり"
	case BodyPress:
		return "ぶつり"
	case FoulPlay:
		return "ぶつり"
	}
	return "とくしゅ"
}

type CategoryFunc func(attacker, defender status.StatusChecker) (at, df *status.RankedValue)

func NewCategory(c DamageCategory) (CategoryFunc, error) {
	switch c {
	case Physical:
		return physical, nil
	case Special:
		return special, nil
	case PsycoShock:
		return psycoShock, nil
	case BodyPress:
		return bodyPress, nil
	case FoulPlay:
		return foulPlay, nil
	case ShellArms:
		return shellArms, nil
	}
	return nil, fmt.Errorf("illegal args %d", c)
}

// 物理
func physical(attacker, defender status.StatusChecker) (at, df *status.RankedValue) {
	return attacker.Attack(), defender.Defense()
}

// 特殊
func special(attacker, defender status.StatusChecker) (at, df *status.RankedValue) {
	return attacker.SpAttack(), defender.SpDefense()
}

// サイコショック
func psycoShock(attacker, defender status.StatusChecker) (at, df *status.RankedValue) {
	return attacker.SpAttack(), defender.Defense()
}

// ボディプレス
func bodyPress(attacker, defender status.StatusChecker) (at, df *status.RankedValue) {
	return attacker.Defense(), defender.Defense()
}

// イカサマ
func foulPlay(attacker, defender status.StatusChecker) (at, df *status.RankedValue) {
	return defender.Attack(), defender.Defense()
}

// シェルアームズ(物理、特殊ダメージ大きいほうで)
func shellArms(attacker, defender status.StatusChecker) (at, df *status.RankedValue) {
	pa := float32(attacker.Attack().Value())
	pd := float32(defender.Defense().Value())
	sa := float32(attacker.SpAttack().Value())
	sd := float32(defender.SpDefense().Value())
	prate := pa / pd
	srate := sa / sd

	if prate > srate {
		return physical(attacker, defender)
	}
	return special(attacker, defender)

}
