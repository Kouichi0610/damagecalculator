package damage

import (
	"damagecalculator/domain/types"
	"testing"
)

/*
	とくせいがダメージ計算に与える影響をテスト

*/
func Test_とくせい_へんげんじざい(t *testing.T) {
	a := defaultSituation()
	a.Attacker.Name = "ミュウツー"
	a.Attacker.Ability = "へんげんじざい"

	a.Move = "れいとうビーム"
	st := toSituation(a)
	if !st.Attacker().Types().Equal(types.NewTypes(types.Ice)) {
		t.Error()
	}

	a.Attacker.Ability = "リベロ"
	a.Move = "じしん"
	st = toSituation(a)
	if !st.Attacker().Types().Equal(types.NewTypes(types.Ground)) {
		t.Error()
	}
}

func Test_とくせい_ふしぎなまもり(t *testing.T) {
	a := defaultSituation()
	a.Defender.Name = "ヌケニン"
	a.Defender.Ability = "ふしぎなまもり"

	a.Move = "れいとうビーム"
	d := calcDamage(a)
	if d.Min() != 0 {
		t.Error()
	}

	a.Move = "かえんほうしゃ"
	d = calcDamage(a)
	if d.Min() == 0 {
		t.Error()
	}
}

func Test_とくせい_かたやぶり(t *testing.T) {
	a := defaultSituation()
	a.Defender.Name = "ヌケニン"
	a.Defender.Ability = "ふしぎなまもり"
	a.Attacker.Ability = "かたやぶり"

	a.Move = "れいとうビーム"
	d := calcDamage(a)
	if d.Min() == 0 {
		t.Error()
	}
}

func Test_とくせい_ちからもち(t *testing.T) {
	a := defaultSituation()
	a.Move = "かわらわり"

	d := calcDamage(a)
	min := d.Min()

	a.Attacker.Ability = "ちからもち"
	d = calcDamage(a)

	amin := d.Min()

	if amin != min*2 {
		t.Error()
	}
}
func Test_とくせい_かがくへんかガス(t *testing.T) {
	a := defaultSituation()
	a.Move = "かわらわり"
	a.Defender.Ability = "かがくへんかガス"

	d := calcDamage(a)
	min := d.Min()

	a.Attacker.Ability = "ちからもち"
	d = calcDamage(a)

	amin := d.Min()

	if amin == min*2 {
		t.Error()
	}
}
