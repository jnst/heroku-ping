package main

import (
	"fmt"

	"github.com/jnst/heroku-ping/ping"
)

func main() {
	err := ping.Start()
	if err != nil {
		fmt.Println(err)
	}
}
