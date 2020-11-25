package usableItem

import (
	"damagecalculator/domain/item"
	"damagecalculator/infra/pokeapi"
)

// 使用可能な持ち物リスト
type (
	UsableItemService interface {
		List(name string) []item.ItemInfo
	}
)

func NewService(it item.Repository) UsableItemService {
	return &service{it: it}
}

type service struct {
	it item.Repository
}

func (s *service) List(name string) []item.ItemInfo {
	ev := pokeapi.NewEvolution()
	list := s.it.List()
	res := make([]item.ItemInfo, 0)
	for _, item := range list {
		if item.Name() == "しんかのきせき" {
			if !ev.HasEvolution(name) {
				continue
			}
		}
		res = append(res, item)
	}
	return res
}
