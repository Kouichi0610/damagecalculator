package damage

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/types"
	"damagecalculator/infra/local"
	"reflect"

	"testing"
)

/*
	とくせい、もちもの、わざの最終チェックテストコード

*/

// デフォルトの急所補正は通常の1.5倍となること
func Test_急所補正(t *testing.T) {
	a := defaultSituation()
	d, r := calcDamage(a)
	defaultMin := d.Min()
	if r.DetermineCount() != 3 {
		t.Errorf("%s", r.String())
	}

	a.IsCritical = true
	d, r = calcDamage(a)
	criticalMin := d.Min()

	if criticalMin != uint(float64(defaultMin)*1.5) {
		t.Error()
	}
	if r.DetermineCount() != 2 {
		t.Errorf("%s", r.String())
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
	d, _ := calcDamage(a)
	if !(d.Min() == 85 && d.Max() == 85) {
		t.Errorf("%d", d.Min())
	}
	a.Defender.Name = "スピアー"
	d, _ = calcDamage(a)
	if !(d.Min() == 85 && d.Max() == 85) {
		t.Errorf("%d", d.Min())
	}
	// 無効にするタイプ(ゴースト)には効かないこと
	a.Defender.Name = "ゲンガー"
	d, _ = calcDamage(a)
	if !(d.Min() == 0 && d.Max() == 0) {
		t.Error()
	}
}

func Test_ヘビーボンバー(t *testing.T) {
	a := defaultSituation()
	a.Move = "ヘビーボンバー"
	a.Attacker.Name = "ピカチュウ"
	a.Defender.Name = "ピカチュウ"
	st := toSituation(a)
	power := st.Move().Power(st)
	if power != 40 {
		t.Errorf("%d", power)
	}
	a.Attacker.Name = "ハガネール"
	st = toSituation(a)
	power = st.Move().Power(st)
	if power != 120 {
		t.Errorf("%d", power)
	}
}

func Test_ジャイロボール_くろいてっきゅう(t *testing.T) {
	a := defaultSituation()
	a.Move = "ジャイロボール"
	a.Attacker.Name = "ハガネール"
	a.Attacker.BasePoints.Speed = 0
	a.Attacker.Nature = stats.Brave
	a.Defender.Name = "プテラ"
	a.Defender.Nature = stats.Jolly

	st := toSituation(a)
	power := st.Move().Power(st)
	if power != 92 {
		t.Errorf("%d", power)
	}
	// くろいてっきゅうのすばやさが下がる効果で威力が上がること
	a.Attacker.Item = "くろいてっきゅう"
	st = toSituation(a)
	power = st.Move().Power(st)
	if power != 150 {
		t.Errorf("%d", power)
	}
}

// フィールド時、1.3倍
func Test_フィールド補正(t *testing.T) {
	a := defaultSituation()
	d, _ := calcDamage(a)
	dmin := d.Min()

	a.Field = field.ElectricField
	d, _ = calcDamage(a)
	fmin := d.Min()

	if fmin != uint(float64(dmin)*1.3) {
		t.Error()
	}
}

// 天候補正時、1.5倍
func Test_天候補正(t *testing.T) {
	a := defaultSituation()
	a.Move = "ほのおのパンチ"
	d, _ := calcDamage(a)
	dmin := d.Min()

	a.Weather = field.Sunny
	d, _ = calcDamage(a)
	fmin := d.Min()

	if fmin != uint(float64(dmin)*1.5) {
		t.Error()
	}
}
func Test_すなあらしいわタイプ補正(t *testing.T) {
	a := defaultSituation()
	a.Move = "なみのり"
	a.Defender.Name = "プテラ"
	a.Weather = field.NoWeather
	d, _ := calcDamage(a)
	nmin := d.Min()
	a.Weather = field.SandStorm
	d, _ = calcDamage(a)
	smin := d.Min()

	if nmin != 50 {
		t.Error()
	}
	if smin != 34 {
		t.Error()
	}
}

func Test_タイプ一致補正(t *testing.T) {
	// カイリキー,フシギダネ(spAttack65)
	a := defaultSituation()
	a.Move = "エナジーボール"
	a.Attacker.Name = "カイリキー"
	d, _ := calcDamage(a)
	nmin := d.Min()
	a.Attacker.Name = "フシギダネ"
	d, _ = calcDamage(a)
	cmin := d.Min()
	if nmin != 16 {
		t.Error()
	}
	if cmin != 24 {
		t.Error()
	}
}

func Test_もちもの補正(t *testing.T) {
	a := defaultSituation()
	a.Move = "なみのり"
	d, _ := calcDamage(a)
	if d.Min() != 27 {
		t.Error()
	}

	a.Attacker.Item = "こだわりメガネ"
	d, _ = calcDamage(a)
	if d.Min() != 39 {
		t.Error()
	}

	a.Defender.Item = "とつげきチョッキ"
	d, _ = calcDamage(a)
	if d.Min() != 27 {
		t.Error()
	}
}

func Test_Default(t *testing.T) {
	a := defaultSituation()
	a.Defender.Name = "ゼニガメ"
	dmg, rate := calcDamage(a)
	t.Errorf("Damages:%s", dmg.String())
	t.Errorf("Rate:%s", rate.String())
}

func defaultSituation() *situation.SituationData {
	d := &situation.SituationData{
		Move: "かみなりパンチ",
		Attacker: situation.PokeData{
			Name:        "ピカチュウ",
			Level:       50,
			Individuals: stats.IndividualTypeMax,
			BasePoints:  situation.BasePoints{6, 252, 0, 0, 0, 252},
			Ranks:       situation.Ranks{0, 0, 0, 0, 0},
			Ability:     "none",
			Item:        "none",
		},
		Defender: situation.PokeData{
			Name:        "ピジョット",
			Level:       50,
			Individuals: stats.IndividualTypeMax,
			BasePoints:  situation.BasePoints{252, 0, 252, 0, 6, 0},
			Ranks:       situation.Ranks{0, 0, 0, 0, 0},
			Ability:     "none",
			Item:        "none",
		},
		Weather:       field.NoWeather,
		Field:         field.NoField,
		IsCritical:    false,
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

func calcDamage(sd *situation.SituationData) (Damages, DamageRate) {
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

func Test_DamageRate(t *testing.T) {
	d := NewDamageRate(100, NewDamages([]uint{49, 50, 51}))
	if d.String() != "49.0% ～ 51.0% 確定数3" {
		t.Errorf("%s", d.String())
	}

	d = NewDamageRate(100, NewDamages([]uint{250}))
	if d.String() != "250.0% 確定数1" {
		t.Errorf("%s", d.String())
	}
}

func Test_DamagesArray(t *testing.T) {
	d := NewDamages([]uint{1, 2, 3, 4, 5})
	e := NewDamages([]uint{1, 2, 3, 4, 5})
	a := d.Array()
	a[0] = 0

	if !reflect.DeepEqual(d, e) {
		t.Error()
	}
}
