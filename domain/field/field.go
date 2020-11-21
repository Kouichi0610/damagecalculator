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

func ToWeather(name string) Weather {
	for k, v := range weatherNames {
		if v == name {
			return k
		}
	}
	return NoWeather
}

func ToField(name string) Field {
	for k, v := range fieldNames {
		if v == name {
			return k
		}
	}
	return NoField
}

func WeatherNames() []string {
	res := make([]string, 0)
	for _, n := range weatherNames {
		res = append(res, n)
	}
	return res
}

func FieldNames() []string {
	res := make([]string, 0)
	for _, n := range fieldNames {
		res = append(res, n)
	}
	return res
}

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

	TODO:グラスフィールド じしん、じならし、マグニチュード半減
	TODO:ミストフィールド ドラゴンタイプ半減
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
		res = append(res, corrector.NewDamage(1.3))
	}
	if f.w.hasPlus(skillType) {
		res = append(res, corrector.NewDamage(1.5))
	}
	if f.w.hasMinus(skillType) {
		res = append(res, corrector.NewDamage(0.5))
	}
	return res
}

// 岩タイプならとくぼう1.5倍
func (f *Fields) StatusCorrector(at, df *types.Types) (ac, dc *status.StatsCorrectors) {
	ac = status.NewStatsCorrectors()
	dc = status.NewStatsCorrectors()
	if f.w == SandStorm {
		if at.Has(types.Rock) {
			ac.SpDefense(1.5)
		}
		if df.Has(types.Rock) {
			dc.SpDefense(1.5)
		}
		return
	}
	return
}

func (f *Fields) HasWeather(w Weather) bool {
	return f.w == w
}
func (f *Fields) HasField(fl Field) bool {
	return f.f == fl
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

var weatherNames map[Weather]string
var fieldNames map[Field]string

func init() {
	weatherNames = make(map[Weather]string)
	fieldNames = make(map[Field]string)
	weatherNames[NoWeather] = "なし"
	weatherNames[Sunny] = "はれ"
	weatherNames[Rainy] = "あめ"
	weatherNames[Snow] = "あられ"
	weatherNames[SandStorm] = "すなあらし"

	fieldNames[NoField] = "なし"
	fieldNames[ElectricField] = "エレキフィールド"
	fieldNames[PsychoField] = "サイコフィールド"
	fieldNames[GrassField] = "グラスフィールド"
	fieldNames[MystField] = "ミストフィールド"
}
