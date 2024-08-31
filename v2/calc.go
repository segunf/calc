// A package that performs addition of two numbers
package calc

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add is a function that takes two numbers values and returns their sum
// It follows the rules specified by https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
