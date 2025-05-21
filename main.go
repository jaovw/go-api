package main

import (
	"fmt"

	"github.com/jaovw/go-api/config"
	"github.com/jaovw/go-api/router"
)

func main() {
	err := config.Init()

	if err != nil {
		fmt.Println(err)
		return
	}
	router.Initialize()
}
