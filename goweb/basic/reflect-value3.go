package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is of uint8: ", v.Kind() == reflect.Uint8)
	x = uint8(v.Uint())
	fmt.Println("X: ", x)
}
