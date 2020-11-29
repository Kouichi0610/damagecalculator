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
	"fmt"
	"sort"
)

// 仮想敵一覧(攻撃側)

type (
	Result interface {
		Target() string
		Move() string
		DamageMin() uint
		DamageMax() uint
		RateMin() float64
		RateMax() float64
		DetermineCount() uint
		String() string
	}

	Results []Result

	Service interface {
		Create(level stats.Level, defender *situation.PokeParams, condition *situation.FieldCondition) Results
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
func (s *service) Create(level stats.Level, defender *situation.PokeParams, condition *situation.FieldCondition) Results {
	res := make(Results, 0)
	maker := supposition.NewAttackersMaker(s.nm, s.sp, s.mv, s.ab, s.it)
	attackers := maker.Attackers(supposition.Defender(defender.Name), supposition.DefenderAbility(defender.Ability))
	rateService := damage.NewDamageService(s.sp, s.mv, s.ab, s.it)

	for _, attacker := range attackers {

		damages, rates, err := rateService.Calculate(level, attacker.Param(), defender, attacker.Move(), condition)
		if err != nil {
			continue
		}
		r := newResult(attacker.Param().AttackersInfo(), attacker.Move(), damages, rates)
		res = append(res, r)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].RateMin() > res[j].RateMin()
	})

	return res
}

func newResult(target, move string, dmgs damage.Damages, rate damage.DamageRate) Result {
	return &result{
		t:    target,
		move: move,
		dmin: dmgs.Min(),
		dmax: dmgs.Max(),
		rmin: rate.RateMin(),
		rmax: rate.RateMax(),
		dcnt: rate.DetermineCount(),
	}
}

type result struct {
	t    string
	move string
	dmin uint
	dmax uint
	rmin float64
	rmax float64
	dcnt uint
}

func (r *result) Target() string {
	return r.t
}
func (r *result) Move() string {
	return r.move
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

func (r *result) String() string {
	return fmt.Sprintf("%s %s Damage:%d-%d Rate:%0.1f-%0.1f, 確定数:%d", r.Target(), r.Move(), r.DamageMin(), r.DamageMax(), r.RateMin(), r.RateMax(), r.DetermineCount())
}
