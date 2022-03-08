package main

import (
	"fmt"
	"log"

	"github.com/End313234/clickup-wrapper/clickup"
)

func main() {
	client, err := clickup.New(clickup.Config{
		Token: "your-token-goes-here",
	})
	if err != nil {
		log.Fatal(err)
	}

	space, err := client.CreateSpace("31026359", clickup.Space{
		Name: "Something",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(space.Id)
}
