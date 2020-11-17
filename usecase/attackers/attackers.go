package attackers

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/damage"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"damagecalculator/domain/stats"
	"damagecalculator/domain/supposition"
	"sort"
)

// 仮想敵一覧(攻撃側)

type (
	Result interface {
		Target() string
		DamageMin() uint
		DamageMax() uint
		RateMin() float64
		RateMax() float64
		DetermineCount() uint
	}

	Results []Result

	Service interface {
		Create(level stats.Level, defender *situation.PokeParams, move string, condition *situation.FieldCondition) Results
	}
)

type service struct {
	nm pokenames.Repository
	sp species.Repository
	mv move.Repository
	ab ability.Repository
	it item.Repository
}

func NewService(nm pokenames.Repository, sp species.Repository, mv move.Repository, ab ability.Repository, it item.Repository) Service {
	return &service{
		nm: nm,
		sp: sp,
		mv: mv,
		ab: ab,
		it: it,
	}
}
func (s *service) Create(level stats.Level, defender *situation.PokeParams, move string, condition *situation.FieldCondition) Results {
	return nil
}

func newResult(target string, dmgs damage.Damages, rate damage.DamageRate) Result {
	return &result{
		t:    target,
		dmin: dmgs.Min(),
		dmax: dmgs.Max(),
		rmin: rate.RateMin(),
		rmax: rate.RateMax(),
		dcnt: rate.DetermineCount(),
	}
}

type result struct {
	t    string
	dmin uint
	dmax uint
	rmin float64
	rmax float64
	dcnt uint
}

func (r *result) Target() string {
	return r.t
}
func (r *result) DamageMin() uint {
	return r.dmin
}
func (r *result) DamageMax() uint {
	return r.dmax
}
func (r *result) RateMin() float64 {
	return r.rmin
}
func (r *result) RateMax() float64 {
	return r.rmax
}
func (r *result) DetermineCount() uint {
	return r.dcnt
}
