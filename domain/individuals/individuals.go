package individuals

/*
	個体値
	HP、こうげき、ぼうぎょ、とくこう、とくぼう、すばやさそれぞれに設定
	個別の範囲は0～31

*/
type (
	Type        uint
	Individual  uint
	Individuals interface {
		HP() Individual
		Attack() Individual
		Defense() Individual
		SpAttack() Individual
		SpDefense() Individual
		Speed() Individual
	}
)

const (
	Max Type = iota
	Slowest
	Weakest
)

func NewIndividual(v uint) Individual {
	return clamp(v)
}

func ToIndividualType(name string) Type {
	t, ok := typesMap[name]
	if !ok {
		return Max
	}
	return t
}

func ToIndividuals(name string) Individuals {
	return ToIndividualType(name).ToIndividuals()
}

func (t Type) ToIndividuals() Individuals {
	if t == Slowest {
		return New(31, 31, 31, 31, 31, 0)
	}
	if t == Weakest {
		return New(31, 0, 31, 31, 31, 31)
	}
	return New(31, 31, 31, 31, 31, 31)
}

func New(hp, at, df, sa, sd, sp uint) Individuals {
	return &individuals{
		hp: clamp(hp),
		at: clamp(at),
		df: clamp(df),
		sa: clamp(sa),
		sd: clamp(sd),
		sp: clamp(sp),
	}
}

func clamp(v uint) Individual {
	const max = 31
	if v > max {
		return max
	}
	return Individual(v)
}

type individuals struct {
	hp, at, df, sa, sd, sp Individual
}

func (i *individuals) HP() Individual {
	return i.hp
}
func (i *individuals) Attack() Individual {
	return i.at
}
func (i *individuals) Defense() Individual {
	return i.df
}
func (i *individuals) SpAttack() Individual {
	return i.sa
}
func (i *individuals) SpDefense() Individual {
	return i.sd
}
func (i *individuals) Speed() Individual {
	return i.sp
}

var typesMap map[string]Type

// TODO:SlowestとWeakestを兼用できない
func init() {
	typesMap = make(map[string]Type, 0)
	typesMap["Max"] = Max
	typesMap["Slowest"] = Slowest
	typesMap["Weakest"] = Weakest
}
