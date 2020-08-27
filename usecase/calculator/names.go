package calculator

type Names []string

func (n Names) Has(name string) bool {
	for _, m := range n {
		if m == name {
			return true
		}
	}
	return false
}
