package types

var strings map[Type]string

func init() {
	strings = map[Type]string{
		Normal:   "ノーマル",
		Fire:     "ほのお",
		Water:    "みず",
		Electric: "でんき",
		Grass:    "くさ",
		Ice:      "こおり",
		Fighting: "かくとう",
		Poison:   "どく",
		Ground:   "じめん",
		Flying:   "ひこう",
		Psychic:  "エスパー",
		Bug:      "むし",
		Rock:     "いわ",
		Ghost:    "ゴースト",
		Dragon:   "ドラゴン",
		Dark:     "あく",
		Steel:    "はがね",
		Fairy:    "フェアリー",
	}

}

func typeToString(t Type) string {
	res, ok := strings[t]
	if !ok {
		return ""
	}
	return res
}
