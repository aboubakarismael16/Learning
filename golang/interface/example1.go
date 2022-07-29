package main

import "fmt"

type Gopher struct {
	name string
}

type coder interface {
	code()
	debug()
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.name)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debugging %s language\n", p.name)
}

func main() {
	s := 100
	var any interface{} = s

	fmt.Println(any)

	g := Gopher{"Go"}
	g.code()
	g.debug()


}
