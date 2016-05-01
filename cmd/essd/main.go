package main

import (
	"fmt"

	"github.com/willemvds/ess/cart"
)

func main() {
	eventStore := &cart.ArrayEventStore{}

	cartId := "someid"

	createCmd := cart.CreateCommand{cartId}
	events, err := cart.Handle(createCmd, eventStore)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(events)

	addItemCmd := cart.AddItemCommand{cartId, "something", 13}
	events, err = cart.Handle(addItemCmd, eventStore)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(events)

	addItemCmd = cart.AddItemCommand{cartId, "something-else", -50}
	events, err = cart.Handle(addItemCmd, eventStore)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(events)

	fmt.Println(eventStore)
}
