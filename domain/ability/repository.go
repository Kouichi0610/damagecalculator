package ability

type Repository interface {
	Get(at, df string) AbilityField
}
