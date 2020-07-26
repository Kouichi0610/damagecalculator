package types

import (
	"testing"
)

// むしタイプ効果チェック
func Test_Bug(t *testing.T) {
	attacker := new(bug)
	expects := expects{
		Fire:     NotVeryEffective(),
		Fighting: NotVeryEffective(),
		Poison:   NotVeryEffective(),
		Flying:   NotVeryEffective(),
		Ghost:    NotVeryEffective(),
		Steel:    NotVeryEffective(),
		Fairy:    NotVeryEffective(),
		Grass:    SuperEffective(),
		Psychic:  SuperEffective(),
		Dark:     SuperEffective(),
		Bug:      FlatEffective(),
		Dragon:   FlatEffective(),
		Electric: FlatEffective(),
		Ground:   FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Rock:     FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// あくタイプ効果チェック
func Test_Dark(t *testing.T) {
	attacker := new(dark)
	expects := expects{
		Fighting: NotVeryEffective(),
		Dark:     NotVeryEffective(),
		Fairy:    NotVeryEffective(),
		Psychic:  SuperEffective(),
		Ghost:    SuperEffective(),
		Steel:    FlatEffective(),
		Dragon:   FlatEffective(),
		Bug:      FlatEffective(),
		Electric: FlatEffective(),
		Fire:     FlatEffective(),
		Flying:   FlatEffective(),
		Grass:    FlatEffective(),
		Ground:   FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Poison:   FlatEffective(),
		Rock:     FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// ドラゴンタイプ効果チェック
func Test_Dragon(t *testing.T) {
	attacker := new(dragon)
	expects := expects{
		Steel:    NotVeryEffective(),
		Dragon:   SuperEffective(),
		Fairy:    NoEffective(),
		Bug:      FlatEffective(),
		Dark:     FlatEffective(),
		Electric: FlatEffective(),
		Fighting: FlatEffective(),
		Fire:     FlatEffective(),
		Flying:   FlatEffective(),
		Ghost:    FlatEffective(),
		Grass:    FlatEffective(),
		Ground:   FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Poison:   FlatEffective(),
		Psychic:  FlatEffective(),
		Rock:     FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// でんきタイプ効果チェック
func Test_Electric(t *testing.T) {
	attacker := new(electric)
	expects := expects{
		Electric: NotVeryEffective(),
		Grass:    NotVeryEffective(),
		Dragon:   NotVeryEffective(),
		Water:    SuperEffective(),
		Flying:   SuperEffective(),
		Ground:   NoEffective(),
		Bug:      FlatEffective(),
		Dark:     FlatEffective(),
		Fairy:    FlatEffective(),
		Fighting: FlatEffective(),
		Fire:     FlatEffective(),
		Ghost:    FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Poison:   FlatEffective(),
		Psychic:  FlatEffective(),
		Rock:     FlatEffective(),
		Steel:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// フェアリータイプ効果チェック
func Test_Fairy(t *testing.T) {
	attacker := new(fairy)
	expects := expects{
		Fire:     NotVeryEffective(),
		Poison:   NotVeryEffective(),
		Steel:    NotVeryEffective(),
		Fighting: SuperEffective(),
		Dragon:   SuperEffective(),
		Dark:     SuperEffective(),
		Bug:      FlatEffective(),
		Electric: FlatEffective(),
		Fairy:    FlatEffective(),
		Flying:   FlatEffective(),
		Ghost:    FlatEffective(),
		Grass:    FlatEffective(),
		Ground:   FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Psychic:  FlatEffective(),
		Rock:     FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// かくとうタイプ効果チェック
func Test_Fighting(t *testing.T) {
	attacker := new(fighting)
	expects := expects{
		Poison:   NotVeryEffective(),
		Flying:   NotVeryEffective(),
		Psychic:  NotVeryEffective(),
		Bug:      NotVeryEffective(),
		Fairy:    NotVeryEffective(),
		Normal:   SuperEffective(),
		Ice:      SuperEffective(),
		Rock:     SuperEffective(),
		Dark:     SuperEffective(),
		Steel:    SuperEffective(),
		Ghost:    NoEffective(),
		Dragon:   FlatEffective(),
		Electric: FlatEffective(),
		Fighting: FlatEffective(),
		Fire:     FlatEffective(),
		Grass:    FlatEffective(),
		Ground:   FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// ほのおタイプ効果チェック
func Test_Fire(t *testing.T) {
	attacker := new(fire)
	expects := expects{
		Fire:     NotVeryEffective(),
		Water:    NotVeryEffective(),
		Rock:     NotVeryEffective(),
		Dragon:   NotVeryEffective(),
		Grass:    SuperEffective(),
		Ice:      SuperEffective(),
		Bug:      SuperEffective(),
		Steel:    SuperEffective(),
		Dark:     FlatEffective(),
		Electric: FlatEffective(),
		Fairy:    FlatEffective(),
		Fighting: FlatEffective(),
		Flying:   FlatEffective(),
		Ghost:    FlatEffective(),
		Ground:   FlatEffective(),
		Normal:   FlatEffective(),
		Poison:   FlatEffective(),
		Psychic:  FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// ひこうタイプ効果チェック
func Test_Flying(t *testing.T) {
	attacker := new(flying)
	expects := expects{
		Electric: NotVeryEffective(),
		Rock:     NotVeryEffective(),
		Steel:    NotVeryEffective(),
		Grass:    SuperEffective(),
		Fighting: SuperEffective(),
		Bug:      SuperEffective(),
		Dark:     FlatEffective(),
		Dragon:   FlatEffective(),
		Fairy:    FlatEffective(),
		Fire:     FlatEffective(),
		Flying:   FlatEffective(),
		Ghost:    FlatEffective(),
		Ground:   FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Poison:   FlatEffective(),
		Psychic:  FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// ゴーストタイプ効果チェック
func Test_Ghost(t *testing.T) {
	attacker := new(ghost)
	expects := expects{
		Dark:     NotVeryEffective(),
		Psychic:  SuperEffective(),
		Ghost:    SuperEffective(),
		Normal:   NoEffective(),
		Bug:      FlatEffective(),
		Dragon:   FlatEffective(),
		Electric: FlatEffective(),
		Fairy:    FlatEffective(),
		Fighting: FlatEffective(),
		Fire:     FlatEffective(),
		Flying:   FlatEffective(),
		Grass:    FlatEffective(),
		Ground:   FlatEffective(),
		Ice:      FlatEffective(),
		Poison:   FlatEffective(),
		Rock:     FlatEffective(),
		Steel:    FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// くさタイプ効果チェック
func Test_Grass(t *testing.T) {
	attacker := new(grass)
	expects := expects{
		Fire:     NotVeryEffective(),
		Grass:    NotVeryEffective(),
		Poison:   NotVeryEffective(),
		Flying:   NotVeryEffective(),
		Bug:      NotVeryEffective(),
		Dragon:   NotVeryEffective(),
		Steel:    NotVeryEffective(),
		Water:    SuperEffective(),
		Ground:   SuperEffective(),
		Rock:     SuperEffective(),
		Dark:     FlatEffective(),
		Electric: FlatEffective(),
		Fairy:    FlatEffective(),
		Fighting: FlatEffective(),
		Ghost:    FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Psychic:  FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// じめんタイプ効果チェック
func Test_Ground(t *testing.T) {
	attacker := new(ground)
	expects := expects{
		Grass:    NotVeryEffective(),
		Bug:      NotVeryEffective(),
		Fire:     SuperEffective(),
		Electric: SuperEffective(),
		Poison:   SuperEffective(),
		Rock:     SuperEffective(),
		Steel:    SuperEffective(),
		Flying:   NoEffective(),
		Dark:     FlatEffective(),
		Dragon:   FlatEffective(),
		Fairy:    FlatEffective(),
		Fighting: FlatEffective(),
		Ghost:    FlatEffective(),
		Ground:   FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Psychic:  FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// こおりタイプ効果チェック
func Test_Ice(t *testing.T) {
	attacker := new(ice)
	expects := expects{
		Fire:     NotVeryEffective(),
		Water:    NotVeryEffective(),
		Ice:      NotVeryEffective(),
		Steel:    NotVeryEffective(),
		Grass:    SuperEffective(),
		Ground:   SuperEffective(),
		Flying:   SuperEffective(),
		Dragon:   SuperEffective(),
		Bug:      FlatEffective(),
		Dark:     FlatEffective(),
		Electric: FlatEffective(),
		Fairy:    FlatEffective(),
		Fighting: FlatEffective(),
		Ghost:    FlatEffective(),
		Normal:   FlatEffective(),
		Poison:   FlatEffective(),
		Psychic:  FlatEffective(),
		Rock:     FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// ノーマルタイプ効果チェック
func Test_Normal(t *testing.T) {
	attacker := new(normal)
	expects := expects{
		Rock:     NotVeryEffective(),
		Steel:    NotVeryEffective(),
		Ghost:    NoEffective(),
		Bug:      FlatEffective(),
		Dark:     FlatEffective(),
		Dragon:   FlatEffective(),
		Electric: FlatEffective(),
		Fairy:    FlatEffective(),
		Fighting: FlatEffective(),
		Fire:     FlatEffective(),
		Flying:   FlatEffective(),
		Grass:    FlatEffective(),
		Ground:   FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Poison:   FlatEffective(),
		Psychic:  FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// どくタイプ効果チェック
func Test_Poison(t *testing.T) {
	attacker := new(poison)
	expects := expects{
		Poison:   NotVeryEffective(),
		Ground:   NotVeryEffective(),
		Rock:     NotVeryEffective(),
		Ghost:    NotVeryEffective(),
		Grass:    SuperEffective(),
		Fairy:    SuperEffective(),
		Steel:    NoEffective(),
		Bug:      FlatEffective(),
		Dark:     FlatEffective(),
		Dragon:   FlatEffective(),
		Electric: FlatEffective(),
		Fighting: FlatEffective(),
		Fire:     FlatEffective(),
		Flying:   FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Psychic:  FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// エスパータイプ効果チェック
func Test_Psychic(t *testing.T) {
	attacker := new(psychic)
	expects := expects{
		Psychic:  NotVeryEffective(),
		Steel:    NotVeryEffective(),
		Fighting: SuperEffective(),
		Poison:   SuperEffective(),
		Dark:     NoEffective(),
		Bug:      FlatEffective(),
		Dragon:   FlatEffective(),
		Electric: FlatEffective(),
		Fairy:    FlatEffective(),
		Fire:     FlatEffective(),
		Flying:   FlatEffective(),
		Ghost:    FlatEffective(),
		Grass:    FlatEffective(),
		Ground:   FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Rock:     FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// いわタイプ効果チェック
func Test_Rock(t *testing.T) {
	attacker := new(rock)
	expects := expects{
		Fighting: NotVeryEffective(),
		Ground:   NotVeryEffective(),
		Steel:    NotVeryEffective(),
		Fire:     SuperEffective(),
		Ice:      SuperEffective(),
		Flying:   SuperEffective(),
		Bug:      SuperEffective(),
		Dark:     FlatEffective(),
		Dragon:   FlatEffective(),
		Electric: FlatEffective(),
		Fairy:    FlatEffective(),
		Ghost:    FlatEffective(),
		Grass:    FlatEffective(),
		Normal:   FlatEffective(),
		Poison:   FlatEffective(),
		Psychic:  FlatEffective(),
		Rock:     FlatEffective(),
		Water:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// はがねタイプ効果チェック
func Test_Steel(t *testing.T) {
	attacker := new(steel)
	expects := expects{
		Fire:     NotVeryEffective(),
		Water:    NotVeryEffective(),
		Electric: NotVeryEffective(),
		Steel:    NotVeryEffective(),
		Ice:      SuperEffective(),
		Rock:     SuperEffective(),
		Fairy:    SuperEffective(),
		Bug:      FlatEffective(),
		Dark:     FlatEffective(),
		Dragon:   FlatEffective(),
		Fighting: FlatEffective(),
		Flying:   FlatEffective(),
		Ghost:    FlatEffective(),
		Grass:    FlatEffective(),
		Ground:   FlatEffective(),
		Normal:   FlatEffective(),
		Poison:   FlatEffective(),
		Psychic:  FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// みずタイプ効果チェック
func Test_Water(t *testing.T) {
	attacker := new(water)
	expects := expects{
		Water:    NotVeryEffective(),
		Grass:    NotVeryEffective(),
		Dragon:   NotVeryEffective(),
		Fire:     SuperEffective(),
		Ground:   SuperEffective(),
		Rock:     SuperEffective(),
		Bug:      FlatEffective(),
		Dark:     FlatEffective(),
		Electric: FlatEffective(),
		Fairy:    FlatEffective(),
		Fighting: FlatEffective(),
		Flying:   FlatEffective(),
		Ghost:    FlatEffective(),
		Ice:      FlatEffective(),
		Normal:   FlatEffective(),
		Poison:   FlatEffective(),
		Psychic:  FlatEffective(),
		Steel:    FlatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

type expects struct {
	Bug      Effective
	Dark     Effective
	Dragon   Effective
	Electric Effective
	Fairy    Effective
	Fighting Effective
	Fire     Effective
	Flying   Effective
	Ghost    Effective
	Grass    Effective
	Ground   Effective
	Ice      Effective
	Normal   Effective
	Poison   Effective
	Psychic  Effective
	Rock     Effective
	Steel    Effective
	Water    Effective
}

func (e expects) ToMap() map[magnifier]Effective {
	res := map[magnifier]Effective{
		new(bug):      e.Bug,
		new(dark):     e.Dark,
		new(dragon):   e.Dragon,
		new(electric): e.Electric,
		new(fairy):    e.Fairy,
		new(fighting): e.Fighting,
		new(fire):     e.Fire,
		new(flying):   e.Flying,
		new(ghost):    e.Ghost,
		new(grass):    e.Grass,
		new(ground):   e.Ground,
		new(ice):      e.Ice,
		new(normal):   e.Normal,
		new(poison):   e.Poison,
		new(psychic):  e.Psychic,
		new(rock):     e.Rock,
		new(steel):    e.Steel,
		new(water):    e.Water,
	}
	return res
}

func magnificationTester(attacker magnifier, e *expects, t *testing.T) {
	expects := e.ToMap()
	for defender, expect := range expects {
		actual := attacker.Magnification((defender))
		if actual != expect {
			t.Errorf("failed %T, %f, %f\n", defender, expect, actual)
		}
	}
}
