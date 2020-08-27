package item

//
type Repository interface {
	Get(name string, isAttacker bool) Item
	List() []string
}
