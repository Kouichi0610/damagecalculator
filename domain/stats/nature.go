package stats

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

func (n *Nature) calcHP(l Level, s Species, i Individual, b BasePoint) uint {
	return n.hp(l, s, i, b)
}
func (n *Nature) calcAttack(l Level, s Species, i Individual, b BasePoint) uint {
	return n.at(l, s, i, b)
}
func (n *Nature) calcDefense(l Level, s Species, i Individual, b BasePoint) uint {
	return n.df(l, s, i, b)
}
func (n *Nature) calcSpAttack(l Level, s Species, i Individual, b BasePoint) uint {
	return n.sa(l, s, i, b)
}
func (n *Nature) calcSpDefense(l Level, s Species, i Individual, b BasePoint) uint {
	return n.sd(l, s, i, b)
}
func (n *Nature) calcSpeed(l Level, s Species, i Individual, b BasePoint) uint {
	return n.sp(l, s, i, b)
}

func NewNature(n NatureType) *Nature {
	return newNatureFuncs[n]()
}
func (n *Nature) ToShadinja() {
	n.hp = calcShadinjaHP
}

type newNatureFunc func() *Nature

var newNatureFuncs map[NatureType]newNatureFunc

func init() {
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
