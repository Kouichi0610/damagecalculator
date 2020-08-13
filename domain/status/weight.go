package status

// 重さ(kg)
type Weight float64

func NewWeight(w float64) Weight {
	if w < 0 {
		w = 0
	}
	return Weight(w)
}
