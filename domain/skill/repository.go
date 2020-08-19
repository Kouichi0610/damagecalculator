package skill

import (
	"damagecalculator/domain/types"
)

/*
	できれば
	TODO:factoryをリポジトリ経由で
	TODO:skill -> move
	TODO:他のパッケージでskill.Knuckleなど定数を使っているところを修正(主にability)
	TODO:ファイル多すぎるのでrepository以外のフォルダ階層を一段下げる
*/
type (
	DamageClass int
	MoveTarget  int
	// TODO:Skillを統一したい
	Move struct {
		Name     string
		Damage   DamageClass
		Power    int
		Type     types.Type
		Accuracy int
		MinHits  int
		MaxHits  int
		Target   MoveTarget
	}

	Repository interface {
		Get(name string) (*Move, error)
	}
)

const (
	Physical DamageClass = iota
	Special
	Status
)
const (
	Select       MoveTarget = iota // 対象1体
	User                           // 自身
	AllOpponents                   // 全ての相手
	AllOther                       // 自分以外全て
)
