// 天候、フィールドの状態
// 基本判定のみ。補正値は使用側に任せる
package field

import "damagecalculator/domain/types"

type Weather uint
type Field uint

const (
	NoWeather Weather = iota
	Sunny
	Rainy
	Snow
	SandStorm
)
const (
	NoField Field = iota
	ElectricField
	PsychoField
	GrassField
	MystField
)

// 技のタイプから補正がかかっているか確認

/*
	影響を与えるもの
	・ダメージ(タイプによる)
	・岩タイプ特防補正

*/

/*
	天候、フィールドはそれぞれ重複なし
	天候とフィールドは併用可

	晴れ		炎+ 水-
	雨		水+ 炎-
	ゆき		補正なし
	すなあらし	岩タイプ特防補正

	エレキフィールド	電気
	サイコフィールド	エスパー
	グラスフィールド	草
	ミスとフィールド	フェアリー


	補正計算式自体はcorrectorに持たせる
*/

type Fields struct {
	f Field
	w Weather
	// ダメージ補正
	// hasWeatherPlus
	// hasWeatherMinus
	// hasFieldPlus

	// 天候によるステータス補正
}

func NewFields(f Field, w Weather) *Fields {
	return &Fields{
		f: f,
		w: w,
	}
}

func (f *Fields) HasWeather(w Weather) bool {
	return f.w == w
}
func (f *Fields) HasField(fl Field) bool {
	return f.f == fl
}

func (f *Fields) HasWeatherPlus(s *types.Types) bool {
	return f.w.hasPlus(s)
}
func (f *Fields) HasWeatherMinus(s *types.Types) bool {
	return f.w.hasMinus(s)
}
func (f *Fields) HasFieldPlus(s *types.Types) bool {
	return f.f.hasPlus(s)
}
func (f *Fields) HasSpDefensePlus(t *types.Types) bool {
	return t.Has(types.Rock) && f.HasWeather(SandStorm)
}

func (w Weather) hasPlus(s *types.Types) bool {
	if w == Sunny {
		return s.Has(types.Fire)
	}
	if w == Rainy {
		return s.Has(types.Water)
	}
	return false
}

func (w Weather) hasMinus(s *types.Types) bool {
	if w == Sunny {
		return s.Has(types.Water)
	}
	if w == Rainy {
		return s.Has(types.Fire)
	}
	return false
}
func (f Field) hasPlus(s *types.Types) bool {
	switch f {
	case ElectricField:
		return s.Has(types.Electric)
	case PsychoField:
		return s.Has(types.Psychic)
	case MystField:
		return s.Has(types.Fairy)
	case GrassField:
		return s.Has(types.Grass)
	}
	return false
}
