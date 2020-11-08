package condition

// 状態異常
type ConditionType uint

const (
	None ConditionType = iota
	Poison
	Burn
	Paralysis
	Sleep
	Ice
)

func FromString(s string) ConditionType {
	for k, v := range names {
		if v == s {
			return k
		}
	}
	return None
}

var names map[ConditionType]string

func init() {
	names = make(map[ConditionType]string, 0)
	names[None] = "なし"
	names[Poison] = "どく"
	names[Burn] = "やけど"
	names[Paralysis] = "まひ"
	names[Sleep] = "ねむり"
	names[Ice] = "こおり"
}
