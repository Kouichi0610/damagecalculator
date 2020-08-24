package local

import (
	"damagecalculator/domain/item"
)

type itemRepository struct {
}

func (r *itemRepository) Get(name string, isAttacker bool) item.Item {
	res, ok := itemMap[name]
	if !ok {
		res = &item.NoItem{}
	}
	return res.Create(isAttacker)
}

var itemMap map[string]item.ItemCreator = map[string]item.ItemCreator{
	"こだわりハチマキ": &item.StatsCorrectData{Attack: 1.5},
	"こだわりメガネ":  &item.StatsCorrectData{SpAttack: 1.5},
	"とつげきチョッキ": &item.StatsCorrectData{SpDefense: 1.5},
	"くろいてっきゅう": &item.StatsCorrectData{Speed: 0.5},
}

/*
たつじんのおび	抜群のダメージ1.2
ちからのハチマキ		物理威力1.1
ものしりメガネ	特殊威力1.1
とつげきチョッキ	とくぼう1.5
ねらいのまと
半減きのみ	(効果抜群時)0.5
しんかのきせき	ぼうぎょ、とくぼう1.5(進化前)

*/
