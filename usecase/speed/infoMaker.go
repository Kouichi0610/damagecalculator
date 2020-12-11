package speed

import (
	"damagecalculator/domain/stats"
	"fmt"
)

/*
	すばやさ種族値に応じて実数値のリストを計算
	49以下 -> 最遅、無振り
	50以上 -> 最速、無振り、補正なし+4、補正なし+252

*/
func createInfos(names []string, level, species uint) SpeedInfos {
	res := make(SpeedInfos, 0)
	makers := createInfoMakers(species)

	for _, maker := range makers {
		res = append(res, maker.create(names, level, species))
	}
	return res
}
func createInfoMakers(species uint) []infoMaker {
	if species <= 49 {
		return []infoMaker{
			new(raw),
			new(lowest),
		}
	}
	return []infoMaker{
		new(highest),
		new(raw252),
		new(raw),
	}
}

type infoMaker interface {
	create(names []string, level, species uint) SpeedInfo
}

type highest struct{}
type lowest struct{}
type raw struct{}
type raw4 struct{}
type raw252 struct{}

// 名前リスト必要
func (s *highest) create(names []string, level, species uint) SpeedInfo {
	return newInfo(makeInfo(species, "最速", names), stats.Highest(stats.NewLevel(level), stats.Species(species)))
}
func (s *lowest) create(names []string, level, species uint) SpeedInfo {
	return newInfo(makeInfo(species, "最遅", names), stats.Lowest(stats.NewLevel(level), stats.Species(species)))
}
func (s *raw) create(names []string, level, species uint) SpeedInfo {
	return newInfo(makeInfo(species, "補正なし+0", names), stats.Raw(stats.NewLevel(level), stats.Species(species)))
}
func (s *raw4) create(names []string, level, species uint) SpeedInfo {
	return newInfo(makeInfo(species, "補正なし+4", names), stats.Raw4(stats.NewLevel(level), stats.Species(species)))
}
func (s *raw252) create(names []string, level, species uint) SpeedInfo {
	return newInfo(makeInfo(species, "補正なし+252", names), stats.Raw252(stats.NewLevel(level), stats.Species(species)))
}

func makeInfo(species uint, title string, names []string) string {
	res := fmt.Sprintf("種族値%d %s", species, title)
	cnt := 0
	for _, name := range names {
		res += " " + name
		cnt++
		if cnt == 3 {
			break
		}
	}
	return res
}

type infoImpl struct {
	info  string
	speed uint
}

func newInfo(info string, speed uint) SpeedInfo {
	return &infoImpl{
		info:  info,
		speed: speed,
	}
}

func (i *infoImpl) Info() string {
	return i.info
}
func (i *infoImpl) Speed() uint {
	return i.speed
}
