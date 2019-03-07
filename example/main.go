package main

import (
	"fmt"

	"github.com/jnst/heroku-ping"
)

func main() {
	err := ping.Start()
	if err != nil {
		fmt.Println(err)
	}
}
