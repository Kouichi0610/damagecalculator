package basePoints

/*
	基礎ポイント
	・HP、こうげき、ぼうぎょ、とくこう、とくぼう、すばやさそれぞれに設定
	・個別の範囲は0～252
	(合計値は508以内であること)

	TODO:Statsから独立
*/
type (
	BasePoint  uint
	BasePoints interface {
		HP() BasePoint
		Attack() BasePoint
		Defense() BasePoint
		SpAttack() BasePoint
		SpDefense() BasePoint
		Speed() BasePoint
		Info() string
	}
)

func NewBasePoint(v uint) BasePoint {
	return clamp(v)
}

func New(hp, at, df, sa, sd, sp uint) BasePoints {
	return newBasePoints(hp, at, df, sa, sd, sp, "")
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
func clamp(s uint) BasePoint {
	const max = 252
	if s > max {
		return max
	}
	return BasePoint(s)
}

type basePoints struct {
	hp, at, df, sa, sd, sp BasePoint
	info                   string
}

func (b *basePoints) HP() BasePoint {
	return b.hp
}
func (b *basePoints) Attack() BasePoint {
	return b.at
}
func (b *basePoints) Defense() BasePoint {
	return b.df
}
func (b *basePoints) SpAttack() BasePoint {
	return b.sa
}
func (b *basePoints) SpDefense() BasePoint {
	return b.sd
}
func (b *basePoints) Speed() BasePoint {
	return b.sp
}
func (b *basePoints) Info() string {
	return b.info
}
