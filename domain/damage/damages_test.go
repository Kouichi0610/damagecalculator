package damage

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/types"
	"damagecalculator/infra/local"
	"reflect"

	"testing"
)

/*
	とくせい、もちもの、わざの最終チェックテストコード

*/

func Test_DamageService(t *testing.T) {
	sp := local.Species()
	mv := local.Move()
	ab := local.Ability()
	it := local.Item()
	sv := NewDamageService(sp, mv, ab, it)

	fields := &situation.FieldCondition{
		Weather:      "なし",
		Field:        "なし",
		HasReflector: false,
		IsCritical:   false,
	}
	attacker := &situation.PokeParams{
		Name:        "ピカチュウ",
		Individuals: "Max",
		BasePoints:  []uint{6, 252, 0, 0, 0, 252},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     "せいでんき",
		Item:        "None",
		Nature:      "のんき",
		Condition:   "なし",
	}
	defender := &situation.PokeParams{
		Name:        "シェルダー",
		Individuals: "Max",
		BasePoints:  []uint{252, 0, 252, 0, 0, 0},
		Ranks:       []int{0, 0, 0, 0, 0},
		Ability:     "げきりゅう",
		Item:        "None",
		Condition:   "なし",
	}
	move := "ボルテッカー"

	_, rate, err := sv.Calculate(50, attacker, defender, move, fields)
	if err != nil {
		t.Error()
	}
	if rate == nil {
		t.Error()
		return
	}

	if rate.String() != "72.3% ～ 85.4% 確定数2" {
		t.Errorf("%s", rate.String())
	}

	// TODO:ダメージがちょっとずれてるのでダメージ計算式確認
}

// デフォルトの急所補正は通常の1.5倍となること
func Test_急所補正(t *testing.T) {
	a := newSituation()
	d, r := a.calcDamage()
	defaultMin := d.Min()
	if r.DetermineCount() != 3 {
		t.Errorf("%s", r.String())
	}

	a.condition.IsCritical = true
	d, r = a.calcDamage()
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
	a := newSituation()
	a.move = "ウェザーボール"
	typesCheck := func(w field.Weather) *types.Types {
		a.condition.Weather = w.String()
		st := a.toSituation()
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
	a.attacker.Name = "バタフリー"
	a.defender.Name = "ケンタロス"

	powerCheck := func(w field.Weather) uint {
		a.condition.Weather = w.String()
		st := a.toSituation()
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
	a := newSituation()
	a.move = "ちきゅうなげ"
	a.level = 85
	// こうかはばつぐんでもいまひとつでもダメージに変化ない事
	a.defender.Name = "ラッキー"
	d, _ := a.calcDamage()
	if !(d.Min() == 85 && d.Max() == 85) {
		t.Errorf("%d", d.Min())
	}
	a.defender.Name = "スピアー"
	d, _ = a.calcDamage()
	if !(d.Min() == 85 && d.Max() == 85) {
		t.Errorf("%d", d.Min())
	}
	// 無効にするタイプ(ゴースト)には効かないこと
	a.defender.Name = "ゲンガー"
	d, _ = a.calcDamage()
	if !(d.Min() == 0 && d.Max() == 0) {
		t.Error()
	}
}

func Test_ヘビーボンバー(t *testing.T) {
	a := newSituation()
	a.move = "ヘビーボンバー"
	a.attacker.Name = "ピカチュウ"
	a.defender.Name = "ピカチュウ"
	st := a.toSituation()
	power := st.Move().Power(st)
	if power != 40 {
		t.Errorf("%d", power)
	}
	a.attacker.Name = "ハガネール"
	st = a.toSituation()
	power = st.Move().Power(st)
	if power != 120 {
		t.Errorf("%d", power)
	}
}

func Test_ジャイロボール_くろいてっきゅう(t *testing.T) {
	a := newSituation()
	a.move = "ジャイロボール"
	a.attacker.Name = "ハガネール"
	a.attacker.BasePoints[5] = 0
	a.attacker.Nature = "ゆうかん"
	a.defender.Name = "プテラ"
	a.defender.Nature = "ようき"

	st := a.toSituation()
	power := st.Move().Power(st)
	if power != 92 {
		t.Errorf("%d", power)
	}
	// くろいてっきゅうのすばやさが下がる効果で威力が上がること
	a.attacker.Item = "くろいてっきゅう"
	st = a.toSituation()
	power = st.Move().Power(st)
	if power != 150 {
		t.Errorf("%d", power)
	}
}

// フィールド時、1.3倍
func Test_フィールド補正(t *testing.T) {
	a := newSituation()
	d, _ := a.calcDamage()
	dmin := d.Min()

	a.condition.Field = "エレキフィールド"
	d, _ = a.calcDamage()
	fmin := d.Min()

	if fmin != uint(float64(dmin)*1.3) {
		t.Error()
	}
}

// 天候補正時、1.5倍
func Test_天候補正(t *testing.T) {
	a := newSituation()
	a.move = "ほのおのパンチ"
	d, _ := a.calcDamage()
	dmin := d.Min()

	a.condition.Weather = "はれ"
	d, _ = a.calcDamage()
	fmin := d.Min()

	if fmin != uint(float64(dmin)*1.5) {
		t.Error()
	}
}
func Test_すなあらしいわタイプ補正(t *testing.T) {
	a := newSituation()
	a.move = "なみのり"
	a.defender.Name = "プテラ"
	a.condition.Weather = "なし"
	d, _ := a.calcDamage()
	nmin := d.Min()
	a.condition.Weather = "すなあらし"
	d, _ = a.calcDamage()
	smin := d.Min()

	if nmin != 52 {
		t.Error()
	}
	if smin != 34 {
		t.Error()
	}
}

func Test_タイプ一致補正(t *testing.T) {
	a := newSituation()
	a.move = "エナジーボール"
	a.attacker.Name = "カイリキー"
	d, _ := a.calcDamage()
	nmin := d.Min()
	a.attacker.Name = "フシギダネ"
	d, _ = a.calcDamage()
	cmin := d.Min()
	if nmin != 16 {
		t.Error()
	}
	if cmin != 24 {
		t.Error()
	}
}

func Test_もちもの補正(t *testing.T) {
	a := newSituation()
	a.move = "なみのり"
	d, _ := a.calcDamage()
	if d.Min() != 27 {
		t.Error()
	}

	a.attacker.Item = "こだわりメガネ"
	d, _ = a.calcDamage()
	if d.Min() != 40 {
		t.Error()
	}

	a.defender.Item = "とつげきチョッキ"
	d, _ = a.calcDamage()
	if d.Min() != 27 {
		t.Error()
	}
}

type battleSituation struct {
	attacker  *situation.PokeParams
	defender  *situation.PokeParams
	level     uint
	move      string
	condition *situation.FieldCondition
}

func newSituation() *battleSituation {
	return &battleSituation{
		level: 50,
		move:  "かみなりパンチ",
		attacker: &situation.PokeParams{
			Name:        "ピカチュウ",
			Individuals: "Max",
			BasePoints:  []uint{6, 252, 0, 0, 0, 252},
			Ranks:       []int{0, 0, 0, 0, 0},
			Ability:     "None",
			Item:        "None",
			Nature:      "てれや",
			Condition:   "なし",
		},
		defender: &situation.PokeParams{
			Name:        "ピジョット",
			Individuals: "Max",
			BasePoints:  []uint{252, 0, 252, 0, 0, 0},
			Ranks:       []int{0, 0, 0, 0, 0},
			Ability:     "None",
			Item:        "None",
			Nature:      "てれや",
			Condition:   "なし",
		},
		condition: &situation.FieldCondition{
			Weather:      "なし",
			Field:        "なし",
			HasReflector: false,
			IsCritical:   false,
		},
	}
}

func (b *battleSituation) toSituation() situation.SituationChecker {
	builder := situation.NewBuilder(local.Species(), local.Ability(), local.Move(), local.Item())
	lv := stats.NewLevel(b.level)
	res, _ := builder.ToSituation(lv, b.attacker, b.defender, b.move, b.condition)
	return res
}
func (b *battleSituation) calcDamage() (Damages, DamageRate) {
	d := NewDamageCalculator()
	st := b.toSituation()
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

	d = NewDamageRate(17, NewDamages([]uint{11}))
	if d.RateMin() != 64.7 {
		t.Errorf("%f", d.RateMin())
	}
	d = NewDamageRate(17, NewDamages([]uint{13}))
	if d.RateMin() != 76.5 {
		t.Errorf("%f", d.RateMin())
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
