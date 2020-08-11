package status

type StatsCorrectors struct {
	at, df, sa, sd, sp statsCorrector
}

func NewStatsCorrectors() *StatsCorrectors {
	return &StatsCorrectors{
		at: defaultStatsCorrector(),
		df: defaultStatsCorrector(),
		sa: defaultStatsCorrector(),
		sd: defaultStatsCorrector(),
		sp: defaultStatsCorrector(),
	}
}

func (s *StatsCorrectors) Attack(numer, denom uint) *StatsCorrectors {
	s.at = newStatsCorrector(numer, denom)
	return s
}
func (s *StatsCorrectors) Defense(numer, denom uint) *StatsCorrectors {
	s.df = newStatsCorrector(numer, denom)
	return s
}
func (s *StatsCorrectors) SpAttack(numer, denom uint) *StatsCorrectors {
	s.sa = newStatsCorrector(numer, denom)
	return s
}
func (s *StatsCorrectors) SpDefense(numer, denom uint) *StatsCorrectors {
	s.sd = newStatsCorrector(numer, denom)
	return s
}
func (s *StatsCorrectors) Speed(numer, denom uint) *StatsCorrectors {
	s.sp = newStatsCorrector(numer, denom)
	return s
}

type statsCorrector interface {
	Correct(n uint) uint
}

type statsCorrectorImpl struct {
	m uint
}

func defaultStatsCorrector() statsCorrector {
	return &statsCorrectorImpl{
		m: 4096,
	}
}

func newStatsCorrector(numer, denom uint) statsCorrector {
	return &statsCorrectorImpl{
		m: numer * one / denom,
	}
}

// 四捨五入
func (s *statsCorrectorImpl) Correct(n uint) uint {
	res := n * s.m
	if res%one >= 2048 {
		return res/one + 1
	}
	return res / one
}

const one = 4096
const half = 2048
