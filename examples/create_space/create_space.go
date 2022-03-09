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

	newSpace := clickup.Space{
		Name: "Something",
	}

	space, err := newSpace.Create(client, "31026359")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(space.Id)
}
