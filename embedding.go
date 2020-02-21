package main

import (
	"fmt"
)

type parent struct {
	Name string
}

type childern struct {
	parent
	roll int32
}

func (p *parent) show() string {
	return "parent class"
}

func (c *childern) setName(name string) {
	c.Name = name
}

func (c *childern) getName() string {
	return c.Name
}

func main() {
	s := childern{}
	s.setName("shohag")
	fmt.Println(s.getName())
	fmt.Println(s.show())
}
