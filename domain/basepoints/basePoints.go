package basePoints

/*
	基礎ポイント
	・HP、こうげき、ぼうぎょ、とくこう、とくぼう、すばやさそれぞれに設定
	・個別の範囲は0～252
	(合計値は508以内であること)

	TODO:Statsから独立
*/
type (
	BasePoints interface {
		HP() uint
		Attack() uint
		Defense() uint
		SpAttack() uint
		SpDefense() uint
		Speed() uint
		Info() string
	}
)

func New(hp, at, df, sa, sd, sp uint) BasePoints {
	return newBasePoints(hp, at, df, sa, sd, sp, "")
}
func newBasePoints(hp, at, df, sa, sd, sp uint, info string) BasePoints {
	return &basePoints{
		hp:   clamp(hp),
		at:   clamp(at),
		df:   clamp(df),
		sa:   clamp(sa),
		sd:   clamp(sd),
		sp:   clamp(sp),
		info: info,
	}
}

func NewFullAttack() BasePoints {
	return newBasePoints(0, 252, 0, 252, 0, 0, "攻撃極振り")
}
func NewHP6() BasePoints {
	return newBasePoints(6, 0, 0, 0, 0, 0, "HP6")
}
func NewFullHP() BasePoints {
	return newBasePoints(252, 0, 0, 0, 0, 0, "HP252")
}
func NewFullGuard() BasePoints {
	return newBasePoints(252, 0, 252, 0, 0, 0, "HP防御252")
}
func NewFullSpGuard() BasePoints {
	return newBasePoints(252, 0, 0, 0, 252, 0, "HP特防252")
}

func clamp(s uint) uint {
	const max = 252
	if s > max {
		return max
	}
	return s
}

type basePoints struct {
	hp, at, df, sa, sd, sp uint
	info                   string
}

func (b *basePoints) HP() uint {
	return b.hp
}
func (b *basePoints) Attack() uint {
	return b.at
}
func (b *basePoints) Defense() uint {
	return b.df
}
func (b *basePoints) SpAttack() uint {
	return b.sa
}
func (b *basePoints) SpDefense() uint {
	return b.sd
}
func (b *basePoints) Speed() uint {
	return b.sp
}
func (b *basePoints) Info() string {
	return b.info
}
