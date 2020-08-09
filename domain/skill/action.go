package skill

type Part uint

const (
	Remote     Part = iota // 遠隔
	Contact                // 接触
	Knuckle                // パンチ系
	Fang                   // 牙
	WaveMotion             // 波動
)

func (p Part) IsContact() bool {
	if p == Remote {
		return false
	}
	if p == WaveMotion {
		return false
	}
	return true
}
