package main

import "fmt"

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

func (c *Cat) Quack()  {
	fmt.Println(c.Name + " meow")
}

type TestStruct struct {}

func NilOrNot(v interface{}) bool  {
	return v == nil
}

func main() {
	var c interface{} = &Cat{Name: "Draven"}
	switch c.(type) {
	case *Cat:
		cat := c.(*Cat)
		cat.Quack()
	}

	var s *TestStruct
	fmt.Println(s == nil)
	fmt.Println(NilOrNot(s))
}

