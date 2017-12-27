package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of v:", p.Type())
	fmt.Println("settability of v:", p.CanSet())

	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)

	fmt.Println(v.Interface())
	fmt.Println(x)
}
