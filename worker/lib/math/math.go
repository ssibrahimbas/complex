package math

type Math struct {
}

func New() *Math {
	return &Math{}
}

func (m *Math) Fib(index int) int {
	if index < 2 {
		return 1
	}
	return m.Fib(index-1) + m.Fib(index-2)
}
