package main

import (
	"fmt"
)

func main() {
	//var greeter gophersGreeter
	greeter := gophersGreeter{
		"sfdsf",
		"fdsfdsf",
		888,
	}
	//fmt.Println(&greeter)
	//fmt.Printf("%p", &greeter)
	//greeter.greet("Hello", "gophers")
	say("say:" + greeter.who)
	_cap := make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		_cap = append(_cap, i)
	}
	fmt.Println(_cap[:])
	fmt.Println(greeter)
}

type gophersGreeter struct {
	how string
	who string
	age int
}

func (greeter *gophersGreeter) greet(how string, who string) {
	//fmt.Printf("%p", greeter)
	//fmt.Println(greeter)
	greeter.who = who
	greeter.how = how
	//fmt.Printf("%s %s!", greeter.how, greeter.who)
	//fmt.Println(greeter)
}

func say(a string) {
	fmt.Println(a)
}
