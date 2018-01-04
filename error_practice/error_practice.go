package main

import (
	"fmt"
	"math"
)
type ErrNegtiveSqrt float64

func (e ErrNegtiveSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return  math.Sqrt(x), nil
	}
	return  x, ErrNegtiveSqrt(x)
}

func main() {
	fmt.Println(Sqrt(-2))
}