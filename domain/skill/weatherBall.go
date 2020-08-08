package skill

/*
	ウェザーボール
	天候 晴れ、雨、あられ、砂嵐の時に威力2倍

	晴れの時ほのおタイプ
	雨の時みずタイプ
	あられの時こおりタイプ
	砂嵐のときにいわタイプ
*/

type weatherBall struct {
	skill
}

/*
	Correctors(SituationChecker) []corrector.Corrector
	Types(SituationChecker) *types.Types
	Power(SituationChecker) uint
	PickStats(SituationChecker) (at, df *status.RankedValue)
	Calculate(level, power, attack, defense uint) []uint
*/
