package stats

// レベル(1～100)
type Level uint

var LevelMin uint = 1
var LevelMax uint = 100

func NewLevel(value uint) Level {
	if value < LevelMin {
		value = LevelMin
	}
	if value > LevelMax {
		value = LevelMax
	}
	res := Level(value)
	return res
}
