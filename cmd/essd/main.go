package main

import (
	"fmt"

	"github.com/willemvds/ess/cart"
)

func main() {
	eventStore := cart.EventStore{}

	createCmd := cart.CreateCommand{"someid"}
	events, err := cart.Handle(createCmd, &eventStore)
	if err != nil {
		fmt.Println(err)
	}
	eventStore.Add(events)

	addItemCmd := cart.AddItemCommand{"someid", "something", 13}
	events, err = cart.Handle(addItemCmd, &eventStore)
	if err != nil {
		fmt.Println(err)
	}
	eventStore.Add(events)

	addItemCmd = cart.AddItemCommand{"someid", "something-else", -50}
	events, err = cart.Handle(addItemCmd, &eventStore)
	if err != nil {
		fmt.Println(err)
	}
	eventStore.Add(events)

	fmt.Println(eventStore)
}
