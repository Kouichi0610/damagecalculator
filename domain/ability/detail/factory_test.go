package detail

import (
	"reflect"
	"testing"
)

// AbilityBuilderが対応するとくせいを生成すること
func Test_Factory(t *testing.T) {
	expects := map[AbilityBuilder]reflect.Type{
		&NoEffectBuilder{}:               reflect.TypeOf(new(abilityImpl)),
		&MoldBreakderBuilder{}:           reflect.TypeOf(new(moldBreaker)),
		&NewtraizingGasBuilder{}:         reflect.TypeOf(new(newtralizingGas)),
		&ForecastBuilder{}:               reflect.TypeOf(new(forecast)),
		&MimicryBuilder{}:                reflect.TypeOf(new(mimicry)),
		&ProteanBuilder{}:                reflect.TypeOf(new(protean)),
		&SkillLinkBuilder{}:              reflect.TypeOf(new(skillLink)),
		&StatusCorrectorBuilder{}:        reflect.TypeOf(new(statusCorrector)),
		&WeatherStatusCorrectorBuilder{}: reflect.TypeOf(new(weatherStatusCorrector)),
		&FieldStatusCorrectorBuilder{}:   reflect.TypeOf(new(fieldStatusCorrector)),
		&WonderGuardBuilder{}:            reflect.TypeOf(new(wonderGuard)),
		&PowerCorrectorBuilder{}:         reflect.TypeOf(new(powerCorrector)),
	}

	for builder, expect := range expects {
		ability := builder.Create()
		actual := reflect.TypeOf(ability)
		if actual != expect {
			t.Errorf("actual[%v] expect[%v]", actual, expect)
		}
	}

}
