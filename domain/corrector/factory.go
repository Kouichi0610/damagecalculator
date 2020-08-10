package corrector

type Category uint

const (
	Power Category = iota
	Attack
	Defense
	Damage

	TypeMatch   // タイプ一致によるダメージ補正
	Critical    // 急所によるダメージ補正
	Protect     // まもるによるダメージ補正
	MultiTarget // 複数対象によるダメージ補正
	Burn        // やけどによるダメージ補正
)

func NewPower(numer, denom uint) Corrector {
	return newCorrector(Power, drop4_pick5, numer, denom)
}

func NewAttack(numer, denom uint) Corrector {
	return newCorrector(Attack, drop4_pick5, numer, denom)
}

func NewDefense(numer, denom uint) Corrector {
	return newCorrector(Defense, drop4_pick5, numer, denom)
}

func NewDamage(numer, denom uint) Corrector {
	return newCorrector(Damage, drop5_pick5over, numer, denom)
}
