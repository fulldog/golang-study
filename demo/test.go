package main

import "fmt"

type pp interface {
	Name(na string, age uint8) interface{}
}

type people interface {
	Name(name string) string
}

type man struct {
	name string
	age  uint8
	xxx  string
}

func (m *man) Name(name string, age uint8) interface{} {
	m.age = age
	m.name = name
	return m
}

type name2 man

func (m *name2) Name(n string) string {
	m.xxx = n
	return m.xxx
}

func main() {
	var t man
	var i1 pp = &t
	var i2 people = (*name2)(&t)

	i1.Name("i1", 3)
	i2.Name("i2")

	fmt.Println(t)
}
