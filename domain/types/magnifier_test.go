package types

import (
	"testing"
)

// むしタイプ効果チェック
func Test_Bug(t *testing.T) {
	attacker := Bug
	expects := expects{
		Fire:     NotVery,
		Fighting: NotVery,
		Poison:   NotVery,
		Flying:   NotVery,
		Ghost:    NotVery,
		Steel:    NotVery,
		Fairy:    NotVery,
		Grass:    Super,
		Psychic:  Super,
		Dark:     Super,
		Bug:      Flat,
		Dragon:   Flat,
		Electric: Flat,
		Ground:   Flat,
		Ice:      Flat,
		Normal:   Flat,
		Rock:     Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// あくタイプ効果チェック
func Test_Dark(t *testing.T) {
	attacker := Dark
	expects := expects{
		Fighting: NotVery,
		Dark:     NotVery,
		Fairy:    NotVery,
		Psychic:  Super,
		Ghost:    Super,
		Steel:    Flat,
		Dragon:   Flat,
		Bug:      Flat,
		Electric: Flat,
		Fire:     Flat,
		Flying:   Flat,
		Grass:    Flat,
		Ground:   Flat,
		Ice:      Flat,
		Normal:   Flat,
		Poison:   Flat,
		Rock:     Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// ドラゴンタイプ効果チェック
func Test_Dragon(t *testing.T) {
	attacker := Dragon
	expects := expects{
		Steel:    NotVery,
		Dragon:   Super,
		Fairy:    NoEffective,
		Bug:      Flat,
		Dark:     Flat,
		Electric: Flat,
		Fighting: Flat,
		Fire:     Flat,
		Flying:   Flat,
		Ghost:    Flat,
		Grass:    Flat,
		Ground:   Flat,
		Ice:      Flat,
		Normal:   Flat,
		Poison:   Flat,
		Psychic:  Flat,
		Rock:     Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// でんきタイプ効果チェック
func Test_Electric(t *testing.T) {
	attacker := Electric
	expects := expects{
		Electric: NotVery,
		Grass:    NotVery,
		Dragon:   NotVery,
		Water:    Super,
		Flying:   Super,
		Ground:   NoEffective,
		Bug:      Flat,
		Dark:     Flat,
		Fairy:    Flat,
		Fighting: Flat,
		Fire:     Flat,
		Ghost:    Flat,
		Ice:      Flat,
		Normal:   Flat,
		Poison:   Flat,
		Psychic:  Flat,
		Rock:     Flat,
		Steel:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// フェアリータイプ効果チェック
func Test_Fairy(t *testing.T) {
	attacker := Fairy
	expects := expects{
		Fire:     NotVery,
		Poison:   NotVery,
		Steel:    NotVery,
		Fighting: Super,
		Dragon:   Super,
		Dark:     Super,
		Bug:      Flat,
		Electric: Flat,
		Fairy:    Flat,
		Flying:   Flat,
		Ghost:    Flat,
		Grass:    Flat,
		Ground:   Flat,
		Ice:      Flat,
		Normal:   Flat,
		Psychic:  Flat,
		Rock:     Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// かくとうタイプ効果チェック
func Test_Fighting(t *testing.T) {
	attacker := Fighting
	expects := expects{
		Poison:   NotVery,
		Flying:   NotVery,
		Psychic:  NotVery,
		Bug:      NotVery,
		Fairy:    NotVery,
		Normal:   Super,
		Ice:      Super,
		Rock:     Super,
		Dark:     Super,
		Steel:    Super,
		Ghost:    NoEffective,
		Dragon:   Flat,
		Electric: Flat,
		Fighting: Flat,
		Fire:     Flat,
		Grass:    Flat,
		Ground:   Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// ほのおタイプ効果チェック
func Test_Fire(t *testing.T) {
	attacker := Fire
	expects := expects{
		Fire:     NotVery,
		Water:    NotVery,
		Rock:     NotVery,
		Dragon:   NotVery,
		Grass:    Super,
		Ice:      Super,
		Bug:      Super,
		Steel:    Super,
		Dark:     Flat,
		Electric: Flat,
		Fairy:    Flat,
		Fighting: Flat,
		Flying:   Flat,
		Ghost:    Flat,
		Ground:   Flat,
		Normal:   Flat,
		Poison:   Flat,
		Psychic:  Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// ひこうタイプ効果チェック
func Test_Flying(t *testing.T) {
	attacker := Flying
	expects := expects{
		Electric: NotVery,
		Rock:     NotVery,
		Steel:    NotVery,
		Grass:    Super,
		Fighting: Super,
		Bug:      Super,
		Dark:     Flat,
		Dragon:   Flat,
		Fairy:    Flat,
		Fire:     Flat,
		Flying:   Flat,
		Ghost:    Flat,
		Ground:   Flat,
		Ice:      Flat,
		Normal:   Flat,
		Poison:   Flat,
		Psychic:  Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// ゴーストタイプ効果チェック
func Test_Ghost(t *testing.T) {
	attacker := Ghost
	expects := expects{
		Dark:     NotVery,
		Psychic:  Super,
		Ghost:    Super,
		Normal:   NoEffective,
		Bug:      Flat,
		Dragon:   Flat,
		Electric: Flat,
		Fairy:    Flat,
		Fighting: Flat,
		Fire:     Flat,
		Flying:   Flat,
		Grass:    Flat,
		Ground:   Flat,
		Ice:      Flat,
		Poison:   Flat,
		Rock:     Flat,
		Steel:    Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// くさタイプ効果チェック
func Test_Grass(t *testing.T) {
	attacker := Grass
	expects := expects{
		Fire:     NotVery,
		Grass:    NotVery,
		Poison:   NotVery,
		Flying:   NotVery,
		Bug:      NotVery,
		Dragon:   NotVery,
		Steel:    NotVery,
		Water:    Super,
		Ground:   Super,
		Rock:     Super,
		Dark:     Flat,
		Electric: Flat,
		Fairy:    Flat,
		Fighting: Flat,
		Ghost:    Flat,
		Ice:      Flat,
		Normal:   Flat,
		Psychic:  Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// じめんタイプ効果チェック
func Test_Ground(t *testing.T) {
	attacker := Ground
	expects := expects{
		Grass:    NotVery,
		Bug:      NotVery,
		Fire:     Super,
		Electric: Super,
		Poison:   Super,
		Rock:     Super,
		Steel:    Super,
		Flying:   NoEffective,
		Dark:     Flat,
		Dragon:   Flat,
		Fairy:    Flat,
		Fighting: Flat,
		Ghost:    Flat,
		Ground:   Flat,
		Ice:      Flat,
		Normal:   Flat,
		Psychic:  Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// こおりタイプ効果チェック
func Test_Ice(t *testing.T) {
	attacker := Ice
	expects := expects{
		Fire:     NotVery,
		Water:    NotVery,
		Ice:      NotVery,
		Steel:    NotVery,
		Grass:    Super,
		Ground:   Super,
		Flying:   Super,
		Dragon:   Super,
		Bug:      Flat,
		Dark:     Flat,
		Electric: Flat,
		Fairy:    Flat,
		Fighting: Flat,
		Ghost:    Flat,
		Normal:   Flat,
		Poison:   Flat,
		Psychic:  Flat,
		Rock:     Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// ノーマルタイプ効果チェック
func Test_Normal(t *testing.T) {
	attacker := Normal
	expects := expects{
		Rock:     NotVery,
		Steel:    NotVery,
		Ghost:    NoEffective,
		Bug:      Flat,
		Dark:     Flat,
		Dragon:   Flat,
		Electric: Flat,
		Fairy:    Flat,
		Fighting: Flat,
		Fire:     Flat,
		Flying:   Flat,
		Grass:    Flat,
		Ground:   Flat,
		Ice:      Flat,
		Normal:   Flat,
		Poison:   Flat,
		Psychic:  Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// どくタイプ効果チェック
func Test_Poison(t *testing.T) {
	attacker := Poison
	expects := expects{
		Poison:   NotVery,
		Ground:   NotVery,
		Rock:     NotVery,
		Ghost:    NotVery,
		Grass:    Super,
		Fairy:    Super,
		Steel:    NoEffective,
		Bug:      Flat,
		Dark:     Flat,
		Dragon:   Flat,
		Electric: Flat,
		Fighting: Flat,
		Fire:     Flat,
		Flying:   Flat,
		Ice:      Flat,
		Normal:   Flat,
		Psychic:  Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// エスパータイプ効果チェック
func Test_Psychic(t *testing.T) {
	attacker := Psychic
	expects := expects{
		Psychic:  NotVery,
		Steel:    NotVery,
		Fighting: Super,
		Poison:   Super,
		Dark:     NoEffective,
		Bug:      Flat,
		Dragon:   Flat,
		Electric: Flat,
		Fairy:    Flat,
		Fire:     Flat,
		Flying:   Flat,
		Ghost:    Flat,
		Grass:    Flat,
		Ground:   Flat,
		Ice:      Flat,
		Normal:   Flat,
		Rock:     Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// いわタイプ効果チェック
func Test_Rock(t *testing.T) {
	attacker := Rock
	expects := expects{
		Fighting: NotVery,
		Ground:   NotVery,
		Steel:    NotVery,
		Fire:     Super,
		Ice:      Super,
		Flying:   Super,
		Bug:      Super,
		Dark:     Flat,
		Dragon:   Flat,
		Electric: Flat,
		Fairy:    Flat,
		Ghost:    Flat,
		Grass:    Flat,
		Normal:   Flat,
		Poison:   Flat,
		Psychic:  Flat,
		Rock:     Flat,
		Water:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// はがねタイプ効果チェック
func Test_Steel(t *testing.T) {
	attacker := Steel
	expects := expects{
		Fire:     NotVery,
		Water:    NotVery,
		Electric: NotVery,
		Steel:    NotVery,
		Ice:      Super,
		Rock:     Super,
		Fairy:    Super,
		Bug:      Flat,
		Dark:     Flat,
		Dragon:   Flat,
		Fighting: Flat,
		Flying:   Flat,
		Ghost:    Flat,
		Grass:    Flat,
		Ground:   Flat,
		Normal:   Flat,
		Poison:   Flat,
		Psychic:  Flat,
	}
	magnificationTester(attacker, &expects, t)
}

// みずタイプ効果チェック
func Test_Water(t *testing.T) {
	attacker := Water
	expects := expects{
		Water:    NotVery,
		Grass:    NotVery,
		Dragon:   NotVery,
		Fire:     Super,
		Ground:   Super,
		Rock:     Super,
		Bug:      Flat,
		Dark:     Flat,
		Electric: Flat,
		Fairy:    Flat,
		Fighting: Flat,
		Flying:   Flat,
		Ghost:    Flat,
		Ice:      Flat,
		Normal:   Flat,
		Poison:   Flat,
		Psychic:  Flat,
		Steel:    Flat,
	}
	magnificationTester(attacker, &expects, t)
}

type expects struct {
	Bug,
	Dark,
	Dragon,
	Electric,
	Fairy,
	Fighting,
	Fire,
	Flying,
	Ghost,
	Grass,
	Ground,
	Ice,
	Normal,
	Poison,
	Psychic,
	Rock,
	Steel,
	Water Effective
}

func (e expects) ToMap() map[Type]Effective {
	res := map[Type]Effective{
		Bug:      e.Bug,
		Dark:     e.Dark,
		Dragon:   e.Dragon,
		Electric: e.Electric,
		Fairy:    e.Fairy,
		Fighting: e.Fighting,
		Fire:     e.Fire,
		Flying:   e.Flying,
		Ghost:    e.Ghost,
		Grass:    e.Grass,
		Ground:   e.Ground,
		Ice:      e.Ice,
		Normal:   e.Normal,
		Poison:   e.Poison,
		Psychic:  e.Psychic,
		Rock:     e.Rock,
		Steel:    e.Steel,
		Water:    e.Water,
	}
	return res
}

func magnificationTester(attacker Type, e *expects, t *testing.T) {
	expects := e.ToMap()
	for defender, expect := range expects {
		actual := attacker.Magnification(defender)
		if actual != expect {
			t.Errorf("failed %T, %f, %f\n", defender, expect, actual)
		}
	}
}
