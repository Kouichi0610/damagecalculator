package local

import (
	"damagecalculator/domain/item"
)

type itemRepository struct {
}

func (r *itemRepository) List() []item.ItemInfo {
	res := make([]item.ItemInfo, 0)
	for _, item := range itemList {
		res = append(res, item)
	}
	return res
}

func (r *itemRepository) Get(name string, isAttacker bool) item.Item {
	res, ok := itemMap[name]
	if !ok {
		res = &item.NoItem{}
	}
	return res.Create(isAttacker)
}

type info struct {
	name, desc string
}

func (i *info) Name() string {
	return i.name
}
func (i *info) Description() string {
	return i.desc
}

var itemList []*info = []*info{
	{"なし", ""},
	{"こだわりハチマキ", "こうげき1.5倍、同じ技しか出せなくなる"},
	{"こだわりメガネ", "とくこう1.5倍、同じ技しか出せなくなる"},
	{"とつげきチョッキ", "とくぼう1.5倍、変化技を出せなくなる"},
	{"くろいてっきゅう", "すばやさ半分、ひこうタイプ、ふゆうにじめんタイプが当たる"},
	{"たつじんのおび", "こうかばつぐんの威力1.2倍"},
	{"しんかのきせき", "ぼうぎょ＆とくぼう1.5倍"},
	{"かるいし", "重さ半分"},
}

var itemMap map[string]item.ItemCreator = map[string]item.ItemCreator{
	"こだわりハチマキ": &item.StatsCorrectData{Attack: 1.5},
	"こだわりメガネ":  &item.StatsCorrectData{SpAttack: 1.5},
	"とつげきチョッキ": &item.StatsCorrectData{SpDefense: 1.5},
	"くろいてっきゅう": &item.StatsCorrectData{Speed: 0.5},
	"たつじんのおび":  &item.SuperEffectiveCorrectData{Scale: 1.2},
	"しんかのきせき":  &item.StatsCorrectData{Defense: 1.5, SpDefense: 1.5},
	"かるいし":     &item.WeightCorrectData{Scale: 0.5},
}

/*
ちからのハチマキ		物理威力1.1
ものしりメガネ	特殊威力1.1
ねらいのまと
半減きのみ	(効果抜群時)威力0.5

*/
