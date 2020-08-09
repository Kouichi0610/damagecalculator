package situation

import (
	"damagecalculator/domain/field"
	"damagecalculator/domain/skill"
	"damagecalculator/domain/status"
)

type SituationData struct {
	Skill              *skill.SkillData
	Attacker, Defender *status.StatusData
	Weather            field.Weather
	Field              field.Field
}

func (s *SituationData) Create() (SituationChecker, error) {
	sk, err := s.Skill.Create()
	if err != nil {
		return nil, err
	}
	return &situation{
		at: s.Attacker.Create(),
		df: s.Defender.Create(),
		sk: sk,
		fl: field.NewFields(s.Field, s.Weather),
	}, nil
}
