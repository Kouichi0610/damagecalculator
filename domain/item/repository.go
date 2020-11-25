package item

//
type (
	ItemInfo interface {
		Name() string
		Description() string
	}
	Repository interface {
		Get(name string, isAttacker bool) Item
		List() []ItemInfo
	}
)
