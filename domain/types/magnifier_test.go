package types

import (
	"testing"
)

// むしタイプ効果チェック
func Test_Bug(t *testing.T) {
	attacker := new(bug)
	expects := expects{
		Fire:     notVeryEffective(),
		Fighting: notVeryEffective(),
		Poison:   notVeryEffective(),
		Flying:   notVeryEffective(),
		Ghost:    notVeryEffective(),
		Steel:    notVeryEffective(),
		Fairy:    notVeryEffective(),
		Grass:    superEffective(),
		Psychic:  superEffective(),
		Dark:     superEffective(),
		Bug:      flatEffective(),
		Dragon:   flatEffective(),
		Electric: flatEffective(),
		Ground:   flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Rock:     flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// あくタイプ効果チェック
func Test_Dark(t *testing.T) {
	attacker := new(dark)
	expects := expects{
		Fighting: notVeryEffective(),
		Dark:     notVeryEffective(),
		Fairy:    notVeryEffective(),
		Psychic:  superEffective(),
		Ghost:    superEffective(),
		Steel:    flatEffective(),
		Dragon:   flatEffective(),
		Bug:      flatEffective(),
		Electric: flatEffective(),
		Fire:     flatEffective(),
		Flying:   flatEffective(),
		Grass:    flatEffective(),
		Ground:   flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Poison:   flatEffective(),
		Rock:     flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// ドラゴンタイプ効果チェック
func Test_Dragon(t *testing.T) {
	attacker := new(dragon)
	expects := expects{
		Steel:    notVeryEffective(),
		Dragon:   superEffective(),
		Fairy:    noEffective(),
		Bug:      flatEffective(),
		Dark:     flatEffective(),
		Electric: flatEffective(),
		Fighting: flatEffective(),
		Fire:     flatEffective(),
		Flying:   flatEffective(),
		Ghost:    flatEffective(),
		Grass:    flatEffective(),
		Ground:   flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Poison:   flatEffective(),
		Psychic:  flatEffective(),
		Rock:     flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// でんきタイプ効果チェック
func Test_Electric(t *testing.T) {
	attacker := new(electric)
	expects := expects{
		Electric: notVeryEffective(),
		Grass:    notVeryEffective(),
		Dragon:   notVeryEffective(),
		Water:    superEffective(),
		Flying:   superEffective(),
		Ground:   noEffective(),
		Bug:      flatEffective(),
		Dark:     flatEffective(),
		Fairy:    flatEffective(),
		Fighting: flatEffective(),
		Fire:     flatEffective(),
		Ghost:    flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Poison:   flatEffective(),
		Psychic:  flatEffective(),
		Rock:     flatEffective(),
		Steel:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// フェアリータイプ効果チェック
func Test_Fairy(t *testing.T) {
	attacker := new(fairy)
	expects := expects{
		Fire:     notVeryEffective(),
		Poison:   notVeryEffective(),
		Steel:    notVeryEffective(),
		Fighting: superEffective(),
		Dragon:   superEffective(),
		Dark:     superEffective(),
		Bug:      flatEffective(),
		Electric: flatEffective(),
		Fairy:    flatEffective(),
		Flying:   flatEffective(),
		Ghost:    flatEffective(),
		Grass:    flatEffective(),
		Ground:   flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Psychic:  flatEffective(),
		Rock:     flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// かくとうタイプ効果チェック
func Test_Fighting(t *testing.T) {
	attacker := new(fighting)
	expects := expects{
		Poison:   notVeryEffective(),
		Flying:   notVeryEffective(),
		Psychic:  notVeryEffective(),
		Bug:      notVeryEffective(),
		Fairy:    notVeryEffective(),
		Normal:   superEffective(),
		Ice:      superEffective(),
		Rock:     superEffective(),
		Dark:     superEffective(),
		Steel:    superEffective(),
		Ghost:    noEffective(),
		Dragon:   flatEffective(),
		Electric: flatEffective(),
		Fighting: flatEffective(),
		Fire:     flatEffective(),
		Grass:    flatEffective(),
		Ground:   flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// ほのおタイプ効果チェック
func Test_Fire(t *testing.T) {
	attacker := new(fire)
	expects := expects{
		Fire:     notVeryEffective(),
		Water:    notVeryEffective(),
		Rock:     notVeryEffective(),
		Dragon:   notVeryEffective(),
		Grass:    superEffective(),
		Ice:      superEffective(),
		Bug:      superEffective(),
		Steel:    superEffective(),
		Dark:     flatEffective(),
		Electric: flatEffective(),
		Fairy:    flatEffective(),
		Fighting: flatEffective(),
		Flying:   flatEffective(),
		Ghost:    flatEffective(),
		Ground:   flatEffective(),
		Normal:   flatEffective(),
		Poison:   flatEffective(),
		Psychic:  flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// ひこうタイプ効果チェック
func Test_Flying(t *testing.T) {
	attacker := new(flying)
	expects := expects{
		Electric: notVeryEffective(),
		Rock:     notVeryEffective(),
		Steel:    notVeryEffective(),
		Grass:    superEffective(),
		Fighting: superEffective(),
		Bug:      superEffective(),
		Dark:     flatEffective(),
		Dragon:   flatEffective(),
		Fairy:    flatEffective(),
		Fire:     flatEffective(),
		Flying:   flatEffective(),
		Ghost:    flatEffective(),
		Ground:   flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Poison:   flatEffective(),
		Psychic:  flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// ゴーストタイプ効果チェック
func Test_Ghost(t *testing.T) {
	attacker := new(ghost)
	expects := expects{
		Dark:     notVeryEffective(),
		Psychic:  superEffective(),
		Ghost:    superEffective(),
		Normal:   noEffective(),
		Bug:      flatEffective(),
		Dragon:   flatEffective(),
		Electric: flatEffective(),
		Fairy:    flatEffective(),
		Fighting: flatEffective(),
		Fire:     flatEffective(),
		Flying:   flatEffective(),
		Grass:    flatEffective(),
		Ground:   flatEffective(),
		Ice:      flatEffective(),
		Poison:   flatEffective(),
		Rock:     flatEffective(),
		Steel:    flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// くさタイプ効果チェック
func Test_Grass(t *testing.T) {
	attacker := new(grass)
	expects := expects{
		Fire:     notVeryEffective(),
		Grass:    notVeryEffective(),
		Poison:   notVeryEffective(),
		Flying:   notVeryEffective(),
		Bug:      notVeryEffective(),
		Dragon:   notVeryEffective(),
		Steel:    notVeryEffective(),
		Water:    superEffective(),
		Ground:   superEffective(),
		Rock:     superEffective(),
		Dark:     flatEffective(),
		Electric: flatEffective(),
		Fairy:    flatEffective(),
		Fighting: flatEffective(),
		Ghost:    flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Psychic:  flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// じめんタイプ効果チェック
func Test_Ground(t *testing.T) {
	attacker := new(ground)
	expects := expects{
		Grass:    notVeryEffective(),
		Bug:      notVeryEffective(),
		Fire:     superEffective(),
		Electric: superEffective(),
		Poison:   superEffective(),
		Rock:     superEffective(),
		Steel:    superEffective(),
		Flying:   noEffective(),
		Dark:     flatEffective(),
		Dragon:   flatEffective(),
		Fairy:    flatEffective(),
		Fighting: flatEffective(),
		Ghost:    flatEffective(),
		Ground:   flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Psychic:  flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// こおりタイプ効果チェック
func Test_Ice(t *testing.T) {
	attacker := new(ice)
	expects := expects{
		Fire:     notVeryEffective(),
		Water:    notVeryEffective(),
		Ice:      notVeryEffective(),
		Steel:    notVeryEffective(),
		Grass:    superEffective(),
		Ground:   superEffective(),
		Flying:   superEffective(),
		Dragon:   superEffective(),
		Bug:      flatEffective(),
		Dark:     flatEffective(),
		Electric: flatEffective(),
		Fairy:    flatEffective(),
		Fighting: flatEffective(),
		Ghost:    flatEffective(),
		Normal:   flatEffective(),
		Poison:   flatEffective(),
		Psychic:  flatEffective(),
		Rock:     flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// ノーマルタイプ効果チェック
func Test_Normal(t *testing.T) {
	attacker := new(normal)
	expects := expects{
		Rock:     notVeryEffective(),
		Steel:    notVeryEffective(),
		Ghost:    noEffective(),
		Bug:      flatEffective(),
		Dark:     flatEffective(),
		Dragon:   flatEffective(),
		Electric: flatEffective(),
		Fairy:    flatEffective(),
		Fighting: flatEffective(),
		Fire:     flatEffective(),
		Flying:   flatEffective(),
		Grass:    flatEffective(),
		Ground:   flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Poison:   flatEffective(),
		Psychic:  flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// どくタイプ効果チェック
func Test_Poison(t *testing.T) {
	attacker := new(poison)
	expects := expects{
		Poison:   notVeryEffective(),
		Ground:   notVeryEffective(),
		Rock:     notVeryEffective(),
		Ghost:    notVeryEffective(),
		Grass:    superEffective(),
		Fairy:    superEffective(),
		Steel:    noEffective(),
		Bug:      flatEffective(),
		Dark:     flatEffective(),
		Dragon:   flatEffective(),
		Electric: flatEffective(),
		Fighting: flatEffective(),
		Fire:     flatEffective(),
		Flying:   flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Psychic:  flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// エスパータイプ効果チェック
func Test_Psychic(t *testing.T) {
	attacker := new(psychic)
	expects := expects{
		Psychic:  notVeryEffective(),
		Steel:    notVeryEffective(),
		Fighting: superEffective(),
		Poison:   superEffective(),
		Dark:     noEffective(),
		Bug:      flatEffective(),
		Dragon:   flatEffective(),
		Electric: flatEffective(),
		Fairy:    flatEffective(),
		Fire:     flatEffective(),
		Flying:   flatEffective(),
		Ghost:    flatEffective(),
		Grass:    flatEffective(),
		Ground:   flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Rock:     flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// いわタイプ効果チェック
func Test_Rock(t *testing.T) {
	attacker := new(rock)
	expects := expects{
		Fighting: notVeryEffective(),
		Ground:   notVeryEffective(),
		Steel:    notVeryEffective(),
		Fire:     superEffective(),
		Ice:      superEffective(),
		Flying:   superEffective(),
		Bug:      superEffective(),
		Dark:     flatEffective(),
		Dragon:   flatEffective(),
		Electric: flatEffective(),
		Fairy:    flatEffective(),
		Ghost:    flatEffective(),
		Grass:    flatEffective(),
		Normal:   flatEffective(),
		Poison:   flatEffective(),
		Psychic:  flatEffective(),
		Rock:     flatEffective(),
		Water:    flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// はがねタイプ効果チェック
func Test_Steel(t *testing.T) {
	attacker := new(steel)
	expects := expects{
		Fire:     notVeryEffective(),
		Water:    notVeryEffective(),
		Electric: notVeryEffective(),
		Steel:    notVeryEffective(),
		Ice:      superEffective(),
		Rock:     superEffective(),
		Fairy:    superEffective(),
		Bug:      flatEffective(),
		Dark:     flatEffective(),
		Dragon:   flatEffective(),
		Fighting: flatEffective(),
		Flying:   flatEffective(),
		Ghost:    flatEffective(),
		Grass:    flatEffective(),
		Ground:   flatEffective(),
		Normal:   flatEffective(),
		Poison:   flatEffective(),
		Psychic:  flatEffective(),
	}
	magnificationTester(attacker, &expects, t)
}

// みずタイプ効果チェック
func Test_Water(t *testing.T) {
	attacker := new(water)
	expects := expects{
		Water:    notVeryEffective(),
		Grass:    notVeryEffective(),
		Dragon:   notVeryEffective(),
		Fire:     superEffective(),
		Ground:   superEffective(),
		Rock:     superEffective(),
		Bug:      flatEffective(),
		Dark:     flatEffective(),
		Electric: flatEffective(),
		Fairy:    flatEffective(),
		Fighting: flatEffective(),
		Flying:   flatEffective(),
		Ghost:    flatEffective(),
		Ice:      flatEffective(),
		Normal:   flatEffective(),
		Poison:   flatEffective(),
		Psychic:  flatEffective(),
		Steel:    flatEffective(),
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
