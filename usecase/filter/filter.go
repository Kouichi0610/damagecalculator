/*
	絞り込むためのAPI

	タイプ
	種族値合計(ツボツボは？) (仮想的絞り込みのため)
	しきい値?(全部69以下を除外)threshold


	TODO:Speciesなどinterfaceにしておく
*/
package filter

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/species"
	"damagecalculator/domain/types"
)

type filterParams struct {
	t types.Type
}

// typefree

func Sample(n pokenames.Repository, s species.Repository, m move.Repository, a ability.Repository, i item.Repository) {
}
