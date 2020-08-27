package pokenames

type Repository interface {
	Get() []string
}
