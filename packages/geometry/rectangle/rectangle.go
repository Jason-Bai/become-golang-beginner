package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("rectangle package initialized")
}

// Area return area
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal return diagonal
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
