package ability

/*
	発動時に素早さが上昇/下降するとくせい
	・条件を満たすとランクが上がる
	・条件を満たしていると実数が上がる
	・条件を満たすと相手の素早さランクを下げる
	・条件の間素早さ実数を下げる(スロースタート)
*/

type ownerSpeedCorrectorGenerator struct {
}
type otherSpeedCorrectorGenerator struct {
}

/*
	TODO:ランクも能力補正もない場合影響なしとみなす
	　　　　をフロント側にゆだねている状況
*/
func (s *ownerSpeedCorrectorGenerator) Create(name string) OwnerSpeedCorrector {
	owner, ok := owners[name]
	if !ok {
		owner = OwnerSpeedCorrector{
			Comment:            "",
			Rank:               0,
			StatsMagnification: 1.0,
		}
	}
	return owner
}
func (s *otherSpeedCorrectorGenerator) Create(name string) OtherSpeedCorrector {
	other, ok := others[name]
	if !ok {
		other = OtherSpeedCorrector{
			Comment:            "",
			Rank:               0,
			StatsMagnification: 1.0,
		}
	}
	return other
}

var owners map[string]OwnerSpeedCorrector
var others map[string]OtherSpeedCorrector

func init() {
	owners = map[string]OwnerSpeedCorrector{
		"かそく":     {"ターン経過ごとにすばやさランク上昇", 1, 1.0},
		"でんきエンジン": {"電気タイプのわざをうけるとすばやさランク上昇", 1, 1.0},
		"すいすい":    {"あめの時すばやさ2倍", 0, 2.0},
		"すなかき":    {"すなあらしの時すばやさ2倍", 0, 2.0},
		"ゆきかき":    {"あられの時すばやさ2倍", 0, 2.0},
		"ようりょくそ":  {"はれの時すばやさ2倍", 0, 2.0},
		"スロースタート": {"5ターンこうげき、すばやさ半分", 0, 0.5},
	}
	others = map[string]OtherSpeedCorrector{
		"ぬめぬめ": {"直接攻撃を受けると相手のすばやさランクを下げる", -1, 1.0},
		"わたげ":  {"攻撃を受けると、自分以外のすばやさランクを下げる", -1, 1.0},
	}
}
