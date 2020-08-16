package types

type magnificationFunc func(defense Type) Effective

var maps map[Type]magnificationFunc = map[Type]magnificationFunc{
	Normal:   normalEffective,
	Fire:     fireEffective,
	Water:    waterEffective,
	Electric: electricEffective,
	Grass:    grassEffective,
	Ice:      iceEffective,
	Fighting: fightingEffective,
	Poison:   poisonEffective,
	Ground:   groundEffective,
	Flying:   flyingEffective,
	Psychic:  psychicEffective,
	Bug:      bugEffective,
	Rock:     rockEffective,
	Ghost:    ghostEffective,
	Dragon:   dragonEffective,
	Dark:     darkEffective,
	Steel:    steelEffective,
	Fairy:    fairyEffective,
}

var names map[Type]string = map[Type]string{
	Normal:   "ノーマル",
	Fire:     "ほのお",
	Water:    "みず",
	Electric: "でんき",
	Grass:    "くさ",
	Ice:      "こおり",
	Fighting: "かくとう",
	Poison:   "どく",
	Ground:   "じめん",
	Flying:   "ひこう",
	Psychic:  "エスパー",
	Bug:      "むし",
	Rock:     "いわ",
	Ghost:    "ゴースト",
	Dragon:   "ドラゴン",
	Dark:     "あく",
	Steel:    "はがね",
	Fairy:    "フェアリー",
}

func normalEffective(df Type) Effective {
	switch df {
	case Rock, Steel:
		return NotVery
	case Ghost:
		return NoEffective
	}
	return Flat
}
func fireEffective(df Type) Effective {
	switch df {
	case Fire, Water, Rock, Dragon:
		return NotVery
	case Grass, Ice, Bug, Steel:
		return Super
	}
	return Flat
}
func waterEffective(df Type) Effective {
	switch df {
	case Water, Grass, Dragon:
		return NotVery
	case Fire, Ground, Rock:
		return Super
	}
	return Flat
}
func electricEffective(df Type) Effective {
	switch df {
	case Electric, Grass, Dragon:
		return NotVery
	case Water, Flying:
		return Super
	case Ground:
		return NoEffective
	}
	return Flat
}
func grassEffective(df Type) Effective {
	switch df {
	case Fire, Grass, Poison, Flying, Bug, Dragon, Steel:
		return NotVery
	case Water, Ground, Rock:
		return Super
	}
	return Flat
}
func iceEffective(df Type) Effective {
	switch df {
	case Fire, Water, Ice, Steel:
		return NotVery
	case Grass, Ground, Flying, Dragon:
		return Super
	}
	return Flat
}
func fightingEffective(df Type) Effective {
	switch df {
	case Poison, Flying, Psychic, Bug, Fairy:
		return NotVery
	case Normal, Ice, Rock, Dark, Steel:
		return Super
	case Ghost:
		return NoEffective
	}
	return Flat
}
func poisonEffective(df Type) Effective {
	switch df {
	case Poison, Ground, Rock, Ghost:
		return NotVery
	case Grass, Fairy:
		return Super
	case Steel:
		return NoEffective
	}
	return Flat
}
func groundEffective(df Type) Effective {
	switch df {
	case Grass, Bug:
		return NotVery
	case Fire, Electric, Poison, Rock, Steel:
		return Super
	case Flying:
		return NoEffective
	}
	return Flat
}
func flyingEffective(df Type) Effective {
	switch df {
	case Electric, Rock, Steel:
		return NotVery
	case Grass, Fighting, Bug:
		return Super
	}
	return Flat
}
func psychicEffective(df Type) Effective {
	switch df {
	case Psychic, Steel:
		return NotVery
	case Fighting, Poison:
		return Super
	case Dark:
		return NoEffective
	}
	return Flat
}
func bugEffective(df Type) Effective {
	switch df {
	case Fire, Fighting, Poison, Flying, Ghost, Steel, Fairy:
		return NotVery
	case Grass, Psychic, Dark:
		return Super
	}
	return Flat
}
func rockEffective(df Type) Effective {
	switch df {
	case Fighting, Ground, Steel:
		return NotVery
	case Fire, Ice, Flying, Bug:
		return Super
	}
	return Flat
}
func ghostEffective(df Type) Effective {
	switch df {
	case Dark:
		return NotVery
	case Psychic, Ghost:
		return Super
	case Normal:
		return NoEffective
	}
	return Flat
}
func dragonEffective(df Type) Effective {
	switch df {
	case Steel:
		return NotVery
	case Dragon:
		return Super
	case Fairy:
		return NoEffective
	}
	return Flat
}
func darkEffective(df Type) Effective {
	switch df {
	case Fighting, Dark, Fairy:
		return NotVery
	case Psychic, Ghost:
		return Super
	}
	return Flat
}
func steelEffective(df Type) Effective {
	switch df {
	case Fire, Water, Electric, Steel:
		return NotVery
	case Ice, Rock, Fairy:
		return Super
	}
	return Flat
}
func fairyEffective(df Type) Effective {
	switch df {
	case Fire, Poison, Steel:
		return NotVery
	case Fighting, Dragon, Dark:
		return Super
	}
	return Flat
}
