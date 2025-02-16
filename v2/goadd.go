package goadd

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes 2 floats or integers and returns the sum.
// [Learn more about addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](num1, num2 T) T {
	return num1 + num2
}
