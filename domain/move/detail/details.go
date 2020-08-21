package detail

// 特殊なわざなどの定義
type Detail uint

const (
	Default     Detail = iota
	SeismicToss        // ちきゅうなげ
	WeatherBall        // ウェザーボール
	GyroBall           // ジャイロボール
	HeavySlam          // ヘビーボンバー
)
