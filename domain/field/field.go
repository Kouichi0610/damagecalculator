// 天候、フィールドの状態
package field

import (
	"damagecalculator/domain/corrector"
	"damagecalculator/domain/status"
	"damagecalculator/domain/types"
)

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
	すなあらし	岩タイプに特防補正

	エレキフィールド	電気
	サイコフィールド	エスパー
	グラスフィールド	草
	ミストフィールド	フェアリー

	補正計算式はcorrectorに持たせる
*/

type Fields struct {
	f Field
	w Weather
}

// Skill -> []Corrector

func NewFields(f Field, w Weather) *Fields {
	return &Fields{
		f: f,
		w: w,
	}
}

func (f *Fields) Correctors(skillType *types.Types) []corrector.Corrector {
	res := make([]corrector.Corrector, 0)

	if f.f.hasPlus(skillType) {
		res = append(res, corrector.NewDamage(13, 10))
	}
	if f.w.hasPlus(skillType) {
		res = append(res, corrector.NewDamage(15, 10))
	}
	if f.w.hasMinus(skillType) {
		res = append(res, corrector.NewDamage(5, 10))
	}
	return res
}

// 岩タイプならとくぼう1.5倍
func (f *Fields) StatusCorrector(at, df *types.Types) (ac, dc *status.StatsCorrectors) {
	ac = status.NewStatsCorrectors()
	dc = status.NewStatsCorrectors()
	if f.w == SandStorm {
		if at.Has(types.Rock) {
			ac.SpDefense(3, 2)
		}
		if df.Has(types.Rock) {
			dc.SpDefense(3, 2)
		}
		return
	}
	return
}

func (f *Fields) HasWeather(w Weather) bool {
	return f.w == w
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
