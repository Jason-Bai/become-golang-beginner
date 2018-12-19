package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func noNeedInspectType() {
	o := order{
		ordId:      1234,
		customerId: 567,
	}
	fmt.Println(createQuery(o))
}

type order1 struct {
	ordId      int
	customerId int
}

type employee1 struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery1(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func reflectPackage() {
	o := order1{
		ordId:      456,
		customerId: 56,
	}
	createQuery1(o)
	e := employee1{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery1(e)
	i := 90
	createQuery1(i)
}

func main() {
	fmt.Println("1. inspect type")
	noNeedInspectType()
	fmt.Println("2. reflect package")
	reflectPackage()
}
