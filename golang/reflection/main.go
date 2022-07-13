package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

func main() {
	author := "draven"
	fmt.Println("TypeOf author : ", reflect.TypeOf(author))
	fmt.Println("ValueOf author ： ", reflect.ValueOf(author))

	v := reflect.ValueOf(1)
	fmt.Println(v.Interface().(int))


	s := new(Student)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(s).Elem())
	fmt.Println(reflect.TypeOf(*s))

	r := reflect.ValueOf(s).Elem()
	r.Field(0).SetString("66")
	fmt.Printf("%#v\n", v.Interface())
}
