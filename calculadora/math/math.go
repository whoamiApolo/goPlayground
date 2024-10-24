package math

func Calc(a, b int) (int, int, int, int) {
	outSum := sum(a, b)
	outSub := sub(a, b)
	outMulti := mult(a, b)
	outDiv := div(a, b)

	return outSum, outSub, outMulti, outDiv
}
func sum(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mult(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
