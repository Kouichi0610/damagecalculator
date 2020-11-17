package moves

import (
	"damagecalculator/domain/move"
	"damagecalculator/domain/species"
)

type (
	Loader interface {
		Moves(name string) Moves
	}
	Move struct {
		Name     string `json:"name"`
		Type     string `json:"type"`
		Power    int    `json:"power"`
		Accuracy int    `json:"accuracy"`
		Category string `json:"category"`
		Mention  string `json:"mention"`
	}
	Moves []*Move
)

func NewLoader(sp species.Repository, mv move.Repository) Loader {
	return &loaderImpl{
		sp: sp,
		mv: mv,
	}
}

type loaderImpl struct {
	sp species.Repository
	mv move.Repository
}

func (s *loaderImpl) Moves(name string) Moves {
	res := make(Moves, 0)
	sp, err := s.sp.Get(name)
	if err != nil {
		return res
	}
	for _, n := range sp.Moves {
		m, err := s.mv.Get(n)
		if err != nil {
			continue
		}
		d := &Move{
			Name:     m.Name,
			Type:     m.Types(),
			Power:    int(m.Power),
			Accuracy: int(m.Accuracy),
			Category: m.Category.String(),
			Mention:  m.Mention(),
		}
		res = append(res, d)
	}
	return res
}
