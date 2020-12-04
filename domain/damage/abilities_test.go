package damage

import (
	"damagecalculator/domain/types"
	"testing"
)

/*
	とくせいがダメージ計算に与える影響をテスト

*/
func Test_とくせい_へんげんじざい(t *testing.T) {
	a := newSituation()
	a.attacker.Name = "ミュウツー"
	a.attacker.Ability = "へんげんじざい"

	a.move = "れいとうビーム"
	st := a.toSituation()
	if !st.Attacker().Types().Equal(types.NewTypes(types.Ice)) {
		t.Error()
	}

	a.attacker.Ability = "リベロ"
	a.move = "じしん"
	st = a.toSituation()
	if !st.Attacker().Types().Equal(types.NewTypes(types.Ground)) {
		t.Error()
	}
}

func Test_とくせい_ふしぎなまもり(t *testing.T) {
	a := newSituation()
	a.defender.Name = "ヌケニン"
	a.defender.Ability = "ふしぎなまもり"

	a.move = "れいとうビーム"
	d, r := a.calcDamage()
	if d.Min() != 0 {
		t.Error()
	}
	if r.DetermineCount() != 0 {
		t.Error()
	}

	a.move = "かえんほうしゃ"
	d, r = a.calcDamage()
	if d.Min() == 0 {
		t.Error()
	}
	if r.DetermineCount() != 2 {
		t.Errorf("%s", r.String())
	}
}

func Test_とくせい_かたやぶり(t *testing.T) {
	a := newSituation()
	a.defender.Name = "ヌケニン"
	a.defender.Ability = "ふしぎなまもり"
	a.attacker.Ability = "かたやぶり"

	a.move = "れいとうビーム"
	d, _ := a.calcDamage()
	if d.Min() == 0 {
		t.Error()
	}
}

func Test_とくせい_ちからもち(t *testing.T) {
	a := newSituation()
	a.move = "かわらわり"

	d, _ := a.calcDamage()
	min := d.Min()

	a.attacker.Ability = "ちからもち"
	d, _ = a.calcDamage()

	amin := d.Min()

	if amin != min*2 {
		t.Error()
	}
}
func Test_とくせい_かがくへんかガス(t *testing.T) {
	a := newSituation()
	a.move = "かわらわり"
	a.defender.Ability = "かがくへんかガス"

	d, _ := a.calcDamage()
	min := d.Min()

	a.attacker.Ability = "ちからもち"
	d, _ = a.calcDamage()

	amin := d.Min()

	if amin == min*2 {
		t.Error()
	}
}
