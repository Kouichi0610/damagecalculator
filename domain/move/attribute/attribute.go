package attribute

// パンチ、牙、波動、追加効果などわざの行動属性の判定
type Attribute interface {
	IsContact() bool    // 接触わざ
	IsKnuckle() bool    // パンチ系
	IsFang() bool       // 牙系
	IsWaveMotion() bool // 波動系
	IsSound() bool      // 音系

	HasAppendEffect() bool // 追加効果あり
	HasRecoil() bool       // 反動
}

type Action uint
type Effect uint

func NewAttribute(a Action, e Effect) Attribute {
	return &attribute{
		a: a,
		e: e,
	}
}

type attribute struct {
	a Action
	e Effect
}

const (
	Remote     Action = iota // 遠隔
	Contact                  // 接触
	Knuckle                  // パンチ系
	Fang                     // 牙
	WaveMotion               // 波動
	Sound                    // 音系
)
const (
	NoAttribute  Effect = iota
	AppendEffect        // 追加効果あり
	Recoil              // 反動
)

func (a *attribute) IsContact() bool {
	if a.a == Remote || a.a == WaveMotion || a.a == Sound {
		return false
	}
	return true
}
func (a *attribute) IsKnuckle() bool {
	return a.a == Knuckle
}
func (a *attribute) IsFang() bool {
	return a.a == Fang
}
func (a *attribute) IsWaveMotion() bool {
	return a.a == WaveMotion
}
func (a *attribute) IsSound() bool {
	return a.a == Sound
}
func (a *attribute) HasAppendEffect() bool {
	return a.e == AppendEffect
}
func (a *attribute) HasRecoil() bool {
	return a.e == Recoil
}
