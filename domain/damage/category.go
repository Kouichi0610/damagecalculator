package damage

/*
	物理、特殊など
	攻撃、防御に使用する能力値を取得
*/
// 攻撃に使用する能力値と防御に使用する能力値を取得
type CategoryType uint
type CategoryFunc func(at, df *Stats) (aval, dval uint)

const (
	Physical CategoryType = iota
	Special
	PsycoShock
	BodyPress
	FoulPlay
	ShellArms
)

func NewCategory(t CategoryType) CategoryFunc {
	switch t {
	case Physical:
		return categoryPhysical
	case Special:
		return categorySpecial
	case PsycoShock:
		return categoryPsycoShock
	case BodyPress:
		return categoryBodyPress
	case FoulPlay:
		return categoryFoulPlay
	case ShellArms:
		return categoryShellArms
	}
	return categoryPhysical
}

func categoryPhysical(at, df *Stats) (aval, dval uint) {
	return at.Attack(), df.Defense()
}
func categorySpecial(at, df *Stats) (aval, dval uint) {
	return at.SpAttack(), df.SpDefense()
}
func categoryPsycoShock(at, df *Stats) (aval, dval uint) {
	return at.SpAttack(), df.Defense()
}
func categoryBodyPress(at, df *Stats) (aval, dval uint) {
	return at.Defense(), df.Defense()
}
func categoryFoulPlay(at, df *Stats) (aval, dval uint) {
	return df.Attack(), df.Defense()
}

// 物理か特殊 威力の高いほうを
func categoryShellArms(at, df *Stats) (aval, dval uint) {
	p := float32(at.Attack()) / float32(df.Defense())
	s := float32(at.SpAttack()) / float32(df.SpDefense())
	if p > s {
		return at.Attack(), df.Defense()
	}
	return at.SpAttack(), df.SpDefense()
}
