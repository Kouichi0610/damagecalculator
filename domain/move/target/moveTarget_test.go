package target

import "testing"

func Test_MoveTargets(t *testing.T) {
	expects := map[MoveTarget]bool{
		Select:       false,
		User:         false,
		AllOpponents: true,
		AllOther:     true,
	}

	for tg, expect := range expects {
		if tg.EnableMultiTarget() != expect {
			t.Error()
		}
	}
}
