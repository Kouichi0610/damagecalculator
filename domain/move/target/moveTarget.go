package target

// わざの対象
type MoveTarget int

const (
	Select       MoveTarget = iota // 対象1体
	User                           // 自身
	AllOpponents                   // 全ての相手
	AllOther                       // 自分以外全て
)

// 複数対象可能
func (t MoveTarget) EnableMultiTarget() bool {
	if t == AllOpponents || t == AllOther {
		return true
	}
	return false
}
