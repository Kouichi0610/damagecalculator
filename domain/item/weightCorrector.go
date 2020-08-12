package item

// 重さ補正
// TODO:statusの重さ定義
type weightCorrector interface {
	Correct(uint) uint
}
