package gender

/*
	ほぼ とうそうしん判定用
	(性別同じとき攻撃が上がる、違うとき下がる)
	とうそうしんの選択肢を用意する
*/
type (
	Gender uint
)

const (
	MaleFemale Gender = iota
	MaleOnly
	FemaleOnly
	Unknown
)

type Relation interface {
	//　無関係
	Unrelated() bool
	// 異なる場合もある
	CouldBeDifferent() bool
	// 必ず同じ
	AlwaysSame() bool
	// 必ず異なる
	AlwaysDifferent() bool
}

func NewRelation(a, b Gender) Relation {
	unrelated := a == Unknown || b == Unknown
	couldbe := !unrelated && (a == MaleFemale || b == MaleFemale)
	alwaysSame := (a == MaleOnly && b == MaleOnly) || (a == FemaleOnly && b == FemaleOnly)
	alwaysDifferent := (a == MaleOnly && b == FemaleOnly) || (a == FemaleOnly && b == MaleOnly)

	return &relation{
		unrelated:        unrelated,
		couldBeDifferent: couldbe,
		alwaysSame:       alwaysSame,
		alwaysDifferent:  alwaysDifferent,
	}
}

type relation struct {
	unrelated, couldBeDifferent, alwaysSame, alwaysDifferent bool
}

func (r *relation) Unrelated() bool {
	return r.unrelated
}
func (r *relation) CouldBeDifferent() bool {
	return r.couldBeDifferent
}
func (r *relation) AlwaysSame() bool {
	return r.alwaysSame
}
func (r *relation) AlwaysDifferent() bool {
	return r.alwaysDifferent
}
