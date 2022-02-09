package math

func Fib(index int) int {
	if index < 2 {
		return 1
	}
	return Fib(index-1) + Fib(index-2)
}
