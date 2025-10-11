package main

import (
	"fmt"

	"github.com/vinodhalaharvi/purekernels/pkg/monoid"
)

func main() {
	m := monoid.NewSumMonoid[int]()
	result := monoid.Reduce(m, []int{1, 2, 3, 4, 5})
	fmt.Println("Sum:", result) // Should print: Sum: 15
}
