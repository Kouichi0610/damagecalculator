package supposition

import (
	_ "damagecalculator/domain/move"
	_ "damagecalculator/domain/species"
	_ "damagecalculator/domain/types"
)

/*
	仮想敵の判定

*/
type AttackerSupposition interface {
	// name nature basePoints info move
}
