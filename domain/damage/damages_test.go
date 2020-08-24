package damage

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"damagecalculator/domain/types"
	"damagecalculator/infra/local"

	"testing"
)

/*
	とくせい、もちもの、わざの最終チェックテストコード

*/

// デフォルトの急所補正は通常の1.5倍となること
func Test_急所補正(t *testing.T) {
	a := defaultSituation()
	d := calcDamage(a)
	defaultMin := d.Min()

	a.IsCritical = true
	d = calcDamage(a)
	criticalMin := d.Min()

	if criticalMin != uint(float64(defaultMin)*1.5) {
		t.Error()
	}
}
func Test_やけど補正(t *testing.T) {
	// TODO:
}
func Test_リフレクター補正(t *testing.T) {
	// TODO:
}
func Test_ひかりのかべ補正(t *testing.T) {
	// TODO:
}

func Test_ウェザーボール(t *testing.T) {
	a := defaultSituation()
	a.Move = "ウェザーボール"
	typesCheck := func(w field.Weather) *types.Types {
		a.Weather = w
		st := toSituation(a)
		return st.MoveTypes()
	}
	typeMap := map[field.Weather]*types.Types{
		field.NoWeather: types.NewTypes(types.Normal),
		field.Sunny:     types.NewTypes(types.Fire),
		field.Rainy:     types.NewTypes(types.Water),
		field.Snow:      types.NewTypes(types.Ice),
		field.SandStorm: types.NewTypes(types.Rock),
	}
	for w, ex := range typeMap {
		ac := typesCheck(w)
		if !ex.Equal(ac) {
			t.Errorf("%d %s", w, ac.String())
		}
	}

	// 天候化では威力が倍になること
	a.Attacker.Name = "バタフリー"
	a.Defender.Name = "ケンタロス"

	powerCheck := func(w field.Weather) uint {
		a.Weather = w
		st := toSituation(a)
		co := st.Correctors()
		res := uint(100)
		for _, c := range co {
			res = c.Correct(res)
		}
		return res
	}
	powerMap := map[field.Weather]uint{
		field.NoWeather: 100,
		field.Sunny:     300, // 威力補正2.0 * 天候補正1.5
		field.Rainy:     300, // 威力補正2.0 * 天候補正1.5
		field.Snow:      200, // 威力補正2.0
		field.SandStorm: 200, // 威力補正2.0
	}
	// 天候補正が入る(晴れ->ほのお1.5倍ダメージ、雨->水1.5倍ダメージ)
	for w, ex := range powerMap {
		ac := powerCheck(w)
		if ac != ex {
			t.Error()
		}
	}
}

func Test_ちきゅうなげ(t *testing.T) {
	a := defaultSituation()
	a.Move = "ちきゅうなげ"
	a.Attacker.Level = 85
	// こうかはばつぐんでもいまひとつでもダメージに変化ない事
	a.Defender.Name = "ラッキー"
	d := calcDamage(a)
	if !(d.Min() == 85 && d.Max() == 85) {
		t.Errorf("%d", d.Min())
	}
	a.Defender.Name = "スピアー"
	d = calcDamage(a)
	if !(d.Min() == 85 && d.Max() == 85) {
		t.Errorf("%d", d.Min())
	}
	// 無効にするタイプ(ゴースト)には効かないこと
	a.Defender.Name = "ゲンガー"
	d = calcDamage(a)
	if !(d.Min() == 0 && d.Max() == 0) {
		t.Error()
	}
}

func Test_ジャイロボール_くろいてっきゅう(t *testing.T) {
}

// フィールド時、1.3倍
func Test_フィールド補正(t *testing.T) {
	a := defaultSituation()
	d := calcDamage(a)
	dmin := d.Min()

	a.Field = field.ElectricField
	d = calcDamage(a)
	fmin := d.Min()

	if fmin != uint(float64(dmin)*1.3) {
		t.Error()
	}
}

// 天候補正時、1.5倍
func Test_天候補正(t *testing.T) {
	a := defaultSituation()
	a.Move = "ほのおのパンチ"
	d := calcDamage(a)
	dmin := d.Min()

	a.Weather = field.Sunny
	d = calcDamage(a)
	fmin := d.Min()

	if fmin != uint(float64(dmin)*1.5) {
		t.Error()
	}
}

func Test_タイプ一致補正(t *testing.T) {

}

/*
func Test_重さ(t *testing.T) {
	a := defaultArgs()
	a.AttackerItem = &item.WeightCorrectData{2.0}
	a.DefenderItem = &item.WeightCorrectData{0.5}
	st, _ := a.Create()
	atWeight := st.Attacker().Weight()
	dfWeight := st.Defender().Weight()
	if atWeight != 200.0 {
		t.Errorf("%f", atWeight)
	}
	if dfWeight != 50.0 {
		t.Errorf("%f", dfWeight)
	}
}
func Test_ヘビーボンバー(t *testing.T) {
	d := NewDamageCalculator()
	a := defaultArgs()
	a.Move = &move.MoveFactory{
		Name:     "unknown",
		Power:    1,
		Type:     types.Steel,
		Accuracy: 100,
		Category: category.Physical,
		CountMin: 1,
		CountMax: 1,
		Target:   target.Select,
		Action:   attribute.Contact,
		Effect:   attribute.NoAttribute,
		Detail:   detail.HeavySlam,
	}
	st, _ := a.Create()
	dmgs := d.CreateDamage(st)
	if dmgs.Min() != 33 {
		t.Errorf("%v", dmgs)
	}
	// 重さがダメージに影響を与えること
	a.AttackerItem = &item.WeightCorrectData{5.0}
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 96 {
		t.Errorf("%v", dmgs)
	}
}

func Test_もちもの補正(t *testing.T) {
	d := NewDamageCalculator()
	a := defaultArgs()
	a.Attacker.Types = []types.Type{types.Fire}
	a.Defender.Types = []types.Type{types.Bug}
	a.Move.Type = types.Normal
	a.AttackerItem = &item.TypeCorrectData{types.Bug, 1.5}
	st, _ := a.Create()
	dmgs := d.CreateDamage(st)
	if dmgs.Min() != 64 {
		t.Errorf("%v", dmgs)
	}
	// 補正を掛けること
	a.AttackerItem = &item.TypeCorrectData{types.Normal, 1.5}
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 96 {
		t.Errorf("%v", dmgs)
	}
	// 防御側の持ち物補正が有効であること
	a.DefenderItem = &item.StatsCorrectData{Defense: 2.0, SpDefense: 2.0}
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 48 {
		t.Errorf("%v", dmgs)
	}
}

func Test_タイプ相性(t *testing.T) {
	d := NewDamageCalculator()
	a := defaultArgs()
	a.Attacker.Types = []types.Type{types.Fire}
	a.Defender.Types = []types.Type{types.Bug}
	a.Move.Type = types.Normal

	// ノーマル->むし
	st, _ := a.Create()
	dmgs := d.CreateDamage(st)
	if dmgs.Min() != 64 {
		t.Errorf("%v", dmgs)
	}
	// いわ->むし
	a.Move.Type = types.Rock
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 128 {
		t.Errorf("%v", dmgs)
	}
	// ほのお(タイプ一致)->むし
	a.Move.Type = types.Fire
	st, _ = a.Create()
	dmgs = d.CreateDamage(st)
	if dmgs.Min() != 192 {
		t.Errorf("%v", dmgs)
	}
}

func Test_急所(t *testing.T) {
	a := defaultArgs()

	dmgs := calcDamage(a)
	if dmgs.Min() != 128 {
		t.Errorf("%v", dmgs)
	}

	a.IsCritical = true
	dmgs = calcDamage(a)
	if dmgs.Min() != 192 {
		t.Errorf("%v", dmgs)
	}
}

// 天候補正時、1.5倍
func Test_天候(t *testing.T) {
	a := defaultArgs()

	a.Move.Type = types.Fire
	a.Defender.Types = []types.Type{types.Normal}
	a.Attacker.Types = []types.Type{types.Normal}

	dmgs := calcDamage(a)
	if dmgs.Min() != 64 {
		t.Errorf("%v", dmgs)
	}

	a.Weather = field.Sunny
	dmgs = calcDamage(a)
	if dmgs.Min() != 96 {
		t.Errorf("%v", dmgs)
	}
}

// フィールド補正時、1.3倍
func Test_フィールド(t *testing.T) {
	a := defaultArgs()

	a.Move.Type = types.Electric
	a.Defender.Types = []types.Type{types.Normal}
	a.Attacker.Types = []types.Type{types.Normal}

	dmgs := calcDamage(a)
	if dmgs.Min() != 64 {
		t.Errorf("%v", dmgs)
	}

	a.Field = field.ElectricField
	dmgs = calcDamage(a)
	if dmgs.Min() != 83 {
		t.Errorf("%v", dmgs)
	}
}

func Test_すなあらし岩タイプ補正(t *testing.T) {
	a := defaultArgs()

	a.Move.Type = types.Ghost
	a.Move.Category = category.Special
	a.Defender.Types = []types.Type{types.Rock}
	a.Attacker.Types = []types.Type{types.Normal}

	dmgs := calcDamage(a)
	if dmgs.Min() != 27 {
		t.Errorf("%v", dmgs)
	}

	// すなあらし時、とくぼう1.5倍
	a.Weather = field.SandStorm
	dmgs = calcDamage(a)
	if dmgs.Min() != 18 {
		t.Errorf("%v", dmgs)
	}
}
*/
func defaultSituation() *situation.SituationData {
	d := &situation.SituationData{
		Move: "かみなりパンチ",
		Attacker: situation.PokeData{
			Name:        "ピカチュウ",
			Level:       50,
			Individuals: situation.Individuals{31, 31, 31, 31, 31, 31},
			BasePoints:  situation.BasePoints{6, 252, 0, 0, 0, 252},
			Ranks:       situation.Ranks{0, 0, 0, 0, 0},
			Ability:     "none",
			Item:        "none",
		},
		Defender: situation.PokeData{
			Name:        "ピジョット",
			Level:       50,
			Individuals: situation.Individuals{31, 31, 31, 31, 31, 31},
			BasePoints:  situation.BasePoints{252, 0, 252, 0, 6, 0},
			Ranks:       situation.Ranks{0, 0, 0, 0, 0},
			Ability:     "none",
			Item:        "none",
		},
		Weather:       field.NoWeather,
		Field:         field.NoField,
		IsCritical:    false,
		IsBurn:        false,
		IsReflector:   false,
		IsLightScreen: false,
	}
	return d
}

func repositories() (mv move.Repository, sp species.Repository, ab ability.Repository, it item.Repository) {
	return local.Move(), local.Species(), local.Ability(), local.Item()
}

func toSituation(sd *situation.SituationData) situation.SituationChecker {
	st, _ := sd.Create(repositories())
	return st
}

func calcDamage(sd *situation.SituationData) *Damages {
	d := NewDamageCalculator()
	st, _ := sd.Create(repositories())
	return d.CreateDamage(st)
}

func Test_Damages_ソートされていること(t *testing.T) {
	d := NewDamages([]uint{5, 4, 2, 10, 8, 3, 1, 6})

	if d.Min() != 1 {
		t.Error()
	}
	if d.Max() != 10 {
		t.Error()
	}
}
