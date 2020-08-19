package accuracy

// 命中率(%)
type Accuracy uint

func NewAccuracy(a uint) Accuracy {
	if a > 100 {
		a = 100
	}
	return Accuracy(a)
}
