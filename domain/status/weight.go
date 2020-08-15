package status

// 重さ(kg)
type Weight float64

// 重さを生成する。0にならないことを保証する。
func NewWeight(w float64) Weight {
	if w < 0 {
		w = 0.01
	}
	return Weight(w)
}
