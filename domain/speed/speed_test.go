package speed

import (
	"damagecalculator/infra/local"
	"testing"
)

func Test_SpeedService(t *testing.T) {
	args := &ServiceArgs{
		Name:        "プテラ",
		Level:       50,
		Individuals: "Max",
		BasePoint:   0,
		Ability:     "いしあたま",
		Nature:      "いじっぱり",
		Item:        "なし",
		Condition:   "なし",
		Weather:     "なし",
		Field:       "なし",
	}

	sv := NewService(local.Species(), local.Ability(), local.Move(), local.Item())
	speed := sv.Speed(args)
	if speed != 150 {
		t.Errorf("%d", speed)
	}

	args.BasePoint = 252
	speed = sv.Speed(args)
	if speed != 182 {
		t.Errorf("%d", speed)
	}

	args.Individuals = "Slowest"
	speed = sv.Speed(args)
	if speed != 166 {
		t.Errorf("%d", speed)
	}
	args.Individuals = "Max"
	args.Nature = "ようき"
	speed = sv.Speed(args)
	if speed != 200 {
		t.Errorf("%d", speed)
	}

	args.Weather = "はれ"
	args.Ability = "ようりょくそ"
	speed = sv.Speed(args)
	if speed != 400 {
		t.Errorf("%d", speed)
	}

	args.Item = "こだわりスカーフ"
	speed = sv.Speed(args)
	if speed != 600 {
		t.Errorf("%d", speed)
	}
}
