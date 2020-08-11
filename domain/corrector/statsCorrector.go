package corrector

/*
	TODO:補正に関するモデリング詰める
	・威力、攻撃、防御、ダメージに補正を掛ける
	・補正はわざ、とくせい、もちものから追加される
	・急所による補正、タイプ一致補正など「状況による補正」
	・状況による補正は書き換えることがある(スナイパーは急所ダメージを1.5->2.25 てきおうりょく タイプ一致補正を1.5 -> 2.0)
*/
type category uint

const (
	// Correctorsに追加する
	Power category = iota
	Attack
	Defense
	Damage

	// Correctorsに上書きする
	TypeMatch   // タイプ一致によるダメージ補正
	Critical    // 急所によるダメージ補正
	Protect     // まもるによるダメージ補正
	MultiTarget // 複数対象によるダメージ補正
	Burn        // やけどによるダメージ補正
)

type StatsCorrector struct {
	c   []Corrector
	env map[category]Corrector
}

func NewStatsCorrector() *StatsCorrector {
	return &StatsCorrector{
		c:   make([]Corrector, 0),
		env: make(map[category]Corrector, 0),
	}
}

func (s *StatsCorrector) CorrectPower(n uint) uint {
	res := n
	for _, c := range s.c {
		if c.Caterogy() != Power {
			continue
		}
		res = c.Correct(res)
	}
	return res
}
func (s *StatsCorrector) CorrectAttack(n uint) uint {
	res := n
	for _, c := range s.c {
		if c.Caterogy() != Attack {
			continue
		}
		res = c.Correct(res)
	}
	return res
}
func (s *StatsCorrector) CorrectDefense(n uint) uint {
	res := n
	for _, c := range s.c {
		if c.Caterogy() != Defense {
			continue
		}
		res = c.Correct(res)
	}
	return res
}
func (s *StatsCorrector) CorrectDamage(n uint) uint {
	res := n
	for _, c := range s.c {
		if c.Caterogy() != Damage {
			continue
		}
		res = c.Correct(res)
	}
	for _, c := range s.env {
		res = c.Correct(res)
	}
	return res
}

func (s *StatsCorrector) Appends(args ...Corrector) {
	for _, c := range args {
		s.append(c)
	}
}

func (s *StatsCorrector) append(c Corrector) {
	// 設定しているなら上書き そうでなければなにもしない
	category := c.Caterogy()
	if category.IsEnvironment() {
		_, ok := s.env[category]
		if ok {
			s.env[category] = c
		}
		return
	}
	s.c = append(s.c, c)
}

// タイプ一致補正(1.5)
func (s *StatsCorrector) ApplyTypeMatch() {
	s.env[TypeMatch] = newCorrector(TypeMatch, drop5_pick5over, 3, 2)
}

// 急所ダメージ補正(1.5)
func (s *StatsCorrector) ApplyCritical() {
	s.env[Critical] = newCorrector(Critical, drop5_pick5over, 3, 2)
}

// 複数対象によるダメージ補正(0.75)
func (s *StatsCorrector) ApplyMultiTarget() {
	s.env[MultiTarget] = newCorrector(MultiTarget, drop5_pick5over, 3, 4)
}

// やけどによるダメージ補正(0.5)
func (s *StatsCorrector) ApplyBurn() {
	s.env[Burn] = newCorrector(Burn, drop5_pick5over, 1, 2)
}

func (c category) IsEnvironment() bool {
	switch c {
	case Power:
		return false
	case Attack:
		return false
	case Defense:
		return false
	case Damage:
		return false
	}
	return true
}
