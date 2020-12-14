package main

import (
	"fmt"
)

func main() {
	greeter := gophersGreeter{"Hello", "gophers", 18}
	greeter.greet("Hello", "gophers")
	say(greeter.who)
}

type gophersGreeter struct {
	how string
	who string
	age int
}

func (greeter gophersGreeter) greet(how string, who string) {
	fmt.Printf("%s %s!", greeter.how, greeter.who)
}

func say(a string) {
	fmt.Print(a)
}
