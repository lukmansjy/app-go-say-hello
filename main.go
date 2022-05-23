package main

import (
	"fmt"

	go_say_hello "github.com/lukmansjy/go-say-hello/v2"
)

func main() {
	name := "Lukman"
	fmt.Println(go_say_hello.SayHello(name))
}
