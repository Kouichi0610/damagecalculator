package skill

/*
	わざで使う部位
	特性で使用する
*/
type Part uint

const (
	Remote     Part = iota // 遠隔
	Contact                // 接触
	Knuckle                // パンチ系
	Fang                   // 牙
	WaveMotion             // 波動
	Sound                  // 音系
)

// TODO:反動技、追加効果

func (p Part) IsContact() bool {
	if p == Remote {
		return false
	}
	if p == WaveMotion {
		return false
	}
	return true
}
