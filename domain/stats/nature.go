package stats

import (
	"damagecalculator/domain/basepoints"
	"damagecalculator/domain/individuals"
)

type NatureType uint

const (
	Bashful NatureType = iota
	Lonely
	Adamant
	Naughty
	Brave
	Bold
	Impish
	Lax
	Relaxed
	Modest
	Mild
	Rash
	Quiet
	Calm
	Gentle
	Careful
	Sassy
	Timid
	Hasty
	Jolly
	Naive
)

type Nature struct {
	hp, at, df, sa, sd, sp statsCalculator
}

type NatureDesc struct {
	Name, Desc string
}

func NatureDescriptions() []NatureDesc {
	res := make([]NatureDesc, 0)
	list := []NatureType{
		Bashful,
		Lonely,
		Adamant,
		Naughty,
		Brave,
		Bold,
		Impish,
		Lax,
		Relaxed,
		Modest,
		Mild,
		Rash,
		Quiet,
		Calm,
		Gentle,
		Careful,
		Sassy,
		Timid,
		Hasty,
		Jolly,
		Naive,
	}
	for _, n := range list {
		name := nameMap[n]
		desc := descMap[n]
		res = append(res, NatureDesc{Name: name, Desc: desc})
	}
	return res
}

func (n *Nature) calcHP(l Level, s Species, i individuals.Individual, b basePoints.BasePoint) uint {
	return n.hp(l, s, i, b)
}
func (n *Nature) calcAttack(l Level, s Species, i individuals.Individual, b basePoints.BasePoint) uint {
	return n.at(l, s, i, b)
}
func (n *Nature) calcDefense(l Level, s Species, i individuals.Individual, b basePoints.BasePoint) uint {
	return n.df(l, s, i, b)
}
func (n *Nature) calcSpAttack(l Level, s Species, i individuals.Individual, b basePoints.BasePoint) uint {
	return n.sa(l, s, i, b)
}
func (n *Nature) calcSpDefense(l Level, s Species, i individuals.Individual, b basePoints.BasePoint) uint {
	return n.sd(l, s, i, b)
}
func (n *Nature) calcSpeed(l Level, s Species, i individuals.Individual, b basePoints.BasePoint) uint {
	return n.sp(l, s, i, b)
}

func NameToNature(name string) *Nature {
	for k, v := range nameMap {
		if v == name {
			return NewNature(k)
		}
	}
	return NewNature(Bashful)
}
func NameToNatureType(name string) NatureType {
	for k, v := range nameMap {
		if v == name {
			return k
		}
	}
	return Bashful
}

func NewNature(n NatureType) *Nature {
	return newNatureFuncs[n]()
}
func (n *Nature) ToShadinja() {
	n.hp = calcShadinjaHP
}

type newNatureFunc func() *Nature

var newNatureFuncs map[NatureType]newNatureFunc

var nameMap map[NatureType]string
var descMap map[NatureType]string

func init() {
	descMap = map[NatureType]string{
		Bashful: "補正なし",
		Lonely:  "攻撃1.1 防御0.9",
		Adamant: "攻撃1.1 特攻0.9",
		Naughty: "攻撃1.1 特防0.9",
		Brave:   "攻撃1.1 素早さ0.9",
		Bold:    "防御1.1 攻撃0.9",
		Impish:  "防御1.1 特攻0.9",
		Lax:     "防御1.1 特防0.9",
		Relaxed: "防御1.1 素早さ0.9",
		Modest:  "特攻1.1 攻撃0.9",
		Mild:    "特攻1.1 防御0.9",
		Rash:    "特攻1.1 特防0.9",
		Quiet:   "特攻1.1 素早さ0.9",
		Calm:    "特防1.1　攻撃0.9",
		Gentle:  "特防1.1 防御0.9",
		Careful: "特防1.1 特攻0.9",
		Sassy:   "特防1.1　素早さ0.9",
		Timid:   "素早さ1.1 攻撃0.9",
		Hasty:   "素早さ1.1 防御0.9",
		Jolly:   "素早さ1.1 特攻0.9",
		Naive:   "素早さ1.1 特防0.9",
	}
	nameMap = map[NatureType]string{
		Bashful: "てれや",
		Lonely:  "さみしがり",
		Adamant: "いじっぱり",
		Naughty: "やんちゃ",
		Brave:   "ゆうかん",
		Bold:    "ずぶとい",
		Impish:  "わんぱく",
		Lax:     "のうてんき",
		Relaxed: "のんき",
		Modest:  "ひかえめ",
		Mild:    "おっとり",
		Rash:    "うっかりや",
		Quiet:   "れいせい",
		Calm:    "おだやか",
		Gentle:  "おとなしい",
		Careful: "しんちょう",
		Sassy:   "なまいき",
		Timid:   "おくびょう",
		Hasty:   "せっかち",
		Jolly:   "ようき",
		Naive:   "むじゃき",
	}
	newNatureFuncs = map[NatureType]newNatureFunc{
		Bashful: bashful,
		Lonely:  lonely,
		Adamant: adamant,
		Naughty: naughty,
		Brave:   brave,
		Bold:    bold,
		Impish:  impish,
		Lax:     lax,
		Relaxed: relaxed,
		Modest:  modest,
		Mild:    mild,
		Rash:    rash,
		Quiet:   quiet,
		Calm:    calm,
		Gentle:  gentle,
		Careful: careful,
		Sassy:   sassy,
		Timid:   timid,
		Hasty:   hasty,
		Jolly:   jolly,
		Naive:   naive,
	}
}

func bashful() *Nature {
	res := new(Nature)
	res.hp = calcHP
	res.at = calcFlat
	res.df = calcFlat
	res.sa = calcFlat
	res.sd = calcFlat
	res.sp = calcFlat
	return res
}
func lonely() *Nature {
	res := bashful()
	res.at = calcUpper
	res.df = calcLower
	return res
}
func adamant() *Nature {
	res := bashful()
	res.at = calcUpper
	res.sa = calcLower
	return res
}
func naughty() *Nature {
	res := bashful()
	res.at = calcUpper
	res.sd = calcLower
	return res
}
func brave() *Nature {
	res := bashful()
	res.at = calcUpper
	res.sp = calcLower
	return res
}
func bold() *Nature {
	res := bashful()
	res.df = calcUpper
	res.at = calcLower
	return res
}
func impish() *Nature {
	res := bashful()
	res.df = calcUpper
	res.sa = calcLower
	return res
}
func lax() *Nature {
	res := bashful()
	res.df = calcUpper
	res.sd = calcLower
	return res
}
func relaxed() *Nature {
	res := bashful()
	res.df = calcUpper
	res.sp = calcLower
	return res
}
func modest() *Nature {
	res := bashful()
	res.sa = calcUpper
	res.at = calcLower
	return res
}
func mild() *Nature {
	res := bashful()
	res.sa = calcUpper
	res.df = calcLower
	return res
}
func rash() *Nature {
	res := bashful()
	res.sa = calcUpper
	res.sd = calcLower
	return res
}
func quiet() *Nature {
	res := bashful()
	res.sa = calcUpper
	res.sp = calcLower
	return res
}
func calm() *Nature {
	res := bashful()
	res.sd = calcUpper
	res.at = calcLower
	return res
}
func gentle() *Nature {
	res := bashful()
	res.sd = calcUpper
	res.df = calcLower
	return res
}
func careful() *Nature {
	res := bashful()
	res.sd = calcUpper
	res.sa = calcLower
	return res
}
func sassy() *Nature {
	res := bashful()
	res.sd = calcUpper
	res.sp = calcLower
	return res
}
func timid() *Nature {
	res := bashful()
	res.sp = calcUpper
	res.at = calcLower
	return res
}
func hasty() *Nature {
	res := bashful()
	res.sp = calcUpper
	res.df = calcLower
	return res
}
func jolly() *Nature {
	res := bashful()
	res.sp = calcUpper
	res.sa = calcLower
	return res
}
func naive() *Nature {
	res := bashful()
	res.sp = calcUpper
	res.sd = calcLower
	return res
}
