package local

import "damagecalculator/infra/index"

// ポケモンの名前リストを取得する
type namesRepository struct {
}

func (r *namesRepository) Get() []string {
	return index.JpNamesArray()
}
