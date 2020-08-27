package pokeapi

import "damagecalculator/infra/index"

// ポケモンの名前リストを取得する
type namesRepository struct {
}

// pokeapiからすべて取れないのでindexから借りる
func (r *namesRepository) Get() []string {
	return index.JpNamesArray()
}
