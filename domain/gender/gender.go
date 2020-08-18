package gender

type (
	Gender uint
)

const (
	MaleFemale Gender = iota
	MaleOnly
	FemaleOnly
	Unknown
)
