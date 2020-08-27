package calculator

import (
	"damagecalculator/domain/ability"
	"damagecalculator/domain/field"
	"damagecalculator/domain/item"
	"damagecalculator/domain/move"
	"damagecalculator/domain/pokenames"
	"damagecalculator/domain/situation"
	"damagecalculator/domain/species"
	"fmt"
)

/*
	ダメージ計算に必要なパラメータ一式をなるべく完結に設定するためのインターフェイス

*/
func NewBuilder(pk pokenames.Repository, mv move.Repository, sp species.Repository, ab ability.Repository, it item.Repository) SituationBuilder {
	return &builder{
		attacker: nil,
		defender: nil,
		moves:    nil,
		fields:   nil,
		options:  options{false, false, false, false},
		pk:       pk,
		mv:       mv,
		sp:       sp,
		ab:       ab,
		it:       it,
	}
}

type SituationBuilder interface {
	NewPokeDataBuilder() PokeDataBuilder
	SetAttacker(PokeDataBuilder)
	SetDefender(PokeDataBuilder)

	MoveList() Names
	SetMove(name string) error

	SetFields(field.Weather, field.Field)
	SetOptions(isCritical, isBurn, isReflector, isLightScreen bool)

	Build() (situation.SituationChecker, error)
}

type builder struct {
	attacker *situation.PokeData
	defender *situation.PokeData
	moves    *moves
	fields   *fields
	options  options
	pk       pokenames.Repository
	mv       move.Repository
	sp       species.Repository
	ab       ability.Repository
	it       item.Repository
}

type moves struct {
	name string
}

type fields struct {
	weather field.Weather
	field   field.Field
}

type options struct {
	isCritical    bool
	isBurn        bool
	isReflector   bool
	isLightScreen bool
}

func (b *builder) SetAttacker(d PokeDataBuilder) {
	b.moves = nil
	at, err := d.Build()
	if err != nil {
		return
	}
	b.attacker = at
}

func (b *builder) SetDefender(d PokeDataBuilder) {
	df, err := d.Build()
	if err != nil {
		return
	}
	b.defender = df
}
func (b *builder) MoveList() Names {
	if b.attacker == nil {
		return nil
	}
	sp, err := b.sp.Get(b.attacker.Name)
	if err != nil {
		return nil
	}
	return sp.Moves
}

func (b *builder) SetMove(name string) error {
	if b.attacker == nil {
		return fmt.Errorf("no attacker.")
	}
	sp, _ := b.sp.Get(b.attacker.Name)

	availableMoves := Names(sp.Moves)
	if !availableMoves.Has(name) {
		return fmt.Errorf("%s not found.", name)
	}
	b.moves = &moves{name}
	return nil
}

func (b *builder) SetFields(w field.Weather, f field.Field) {
	b.fields = &fields{
		weather: w,
		field:   f,
	}
}

func (b *builder) SetOptions(isCritical, isBurn, isReflector, isLightScreen bool) {
	b.options.isCritical = isCritical
	b.options.isBurn = isBurn
	b.options.isReflector = isReflector
	b.options.isLightScreen = isLightScreen
}

func (b *builder) Build() (situation.SituationChecker, error) {
	if b.attacker == nil || b.defender == nil || b.moves == nil || b.fields == nil {
		return nil, fmt.Errorf("no data.")
	}

	res := &situation.SituationData{
		Attacker:      *b.attacker,
		Defender:      *b.defender,
		Move:          b.moves.name,
		Weather:       b.fields.weather,
		Field:         b.fields.field,
		IsCritical:    b.options.isCritical,
		IsBurn:        b.options.isBurn,
		IsReflector:   b.options.isReflector,
		IsLightScreen: b.options.isLightScreen,
	}

	return res.Create(b.mv, b.sp, b.ab, b.it)
}

func (b *builder) NewPokeDataBuilder() PokeDataBuilder {
	return &pokeDataBuilder{
		pk: b.pk,
		mv: b.mv,
		sp: b.sp,
		ab: b.ab,
		it: b.it,
	}
}
