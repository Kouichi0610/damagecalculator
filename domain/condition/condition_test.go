package condition

import (
	"testing"
)

func Test_FromString(t *testing.T) {
	expects := map[string]ConditionType{
		"なし":  None,
		"どく":  Poison,
		"まひ":  Paralysis,
		"やけど": Burn,
		"こおり": Ice,
		"ねむり": Sleep,
		"":    None,
		"その他": None,
	}

	for name, expect := range expects {
		actual := FromString(name)
		if actual != expect {
			t.Errorf("Actual:%d Expect:%d", actual, expect)
		}
	}
}
