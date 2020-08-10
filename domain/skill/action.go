package skill

/*
	わざで使う部位
	特性で使用する
*/
type Action uint
type Attribute uint

const (
	Remote     Action = iota // 遠隔
	Contact                  // 接触
	Knuckle                  // パンチ系
	Fang                     // 牙
	WaveMotion               // 波動
	Sound                    // 音系
)

const (
	NoAttribute  Attribute = iota
	AppendEffect           // 追加効果あり
	Recoil                 // 反動
)

func (p Action) IsContact() bool {
	if p == Remote {
		return false
	}
	if p == WaveMotion {
		return false
	}
	return true
}
