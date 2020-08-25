package usecase

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"fmt"
)

/*
	ダメージ計算手続き
	process
	TODO:名称

	攻撃側名前
		とくせい
		わざ
	防御側名前
		とくせい

	もちもの、状況


	設定必要なもの

*/
type CalculateProcessor interface {
	AttackerName(string)
	AttackerAbilities() []string
	Moves() []string
	DefenderName(string)
	DefenderAbilities() []string
}

type calculatorProcess struct {
	situation.SituationData

	move    move.Repository
	ability ability.Repository
	species species.Repository
	item    item.Repository
}

func NewCalculatorProcess(mv move.Repository, sp species.Repository, ab ability.Repository, it item.Repository) CalculateProcessor {
	return &calculatorProcess{
		move:    mv,
		species: sp,
		ability: ab,
		item:    it,
	}
}

func (c *calculatorProcess) AttackerName(name string) {
	c.Attacker.Name = name
}
func (c *calculatorProcess) AttackerAbilities() []string {
	sp, err := c.species.Get(c.Attacker.Name)
	if err != nil {
		return nil
	}
	res := make([]string, 0)
	res = append(res, sp.Abilities...)
	return res
}
func (c *calculatorProcess) Moves() []string {
	sp, err := c.species.Get(c.Attacker.Name)
	if err != nil {
		return nil
	}
	res := make([]string, 0)

	for _, move := range sp.Moves {
		m, err := c.move.Get(move)
		if err != nil {
			fmt.Printf("error %s\n", err)
			continue
		}
		res = append(res, m.Name)
	}
	return res
}
func (c *calculatorProcess) DefenderName(name string) {
	c.Defender.Name = name
}
func (c *calculatorProcess) DefenderAbilities() []string {
	sp, err := c.species.Get(c.Defender.Name)
	if err != nil {
		return nil
	}
	res := make([]string, 0)
	res = append(res, sp.Abilities...)
	return res
}
