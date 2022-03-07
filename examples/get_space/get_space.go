package main

import (
	"fmt"
	"log"

	"github.com/End313234/clickup-wrapper/clickup"
)

func main() {
	client, err := clickup.New(clickup.Config{
		Token: "pk_43081155_BJ12D0ADF3SOHK2BC41DRDD11KBUD3BU",
	})
	if err != nil {
		log.Fatal(err)
	}

	space, err := client.GetSpace("49191295")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(space)
}
