package main

import (
	"fmt"

	"github.com/willemvds/ess/cart"
)

type AddItemToCart struct {
	CartId string
	ItemId string
}

func main() {
	ev1 := cart.ItemAddedToCart{"CartX2", cart.Item{"Peanuts"}}
	ev2 := cart.ItemAddedToCart{"CartX2", cart.Item{"Winamp"}}

	eventStore := &cart.EventStore{}
	eventStore.Add(ev1)
	eventStore.Add(ev2)

	cmd := cart.AddItemToCart{
		CartId: "CartX2",
		ItemId: "Jingle Bells",
	}

	events, err := cart.HandleAddItemToCart(eventStore, cmd)
	fmt.Println("HandleAddItemToCart (cmd1) err=", err, ", events=", events)

	cmd2 := cart.AddItemToCart{
		CartId: "CartX3",
		ItemId: "Noobs",
	}

	events, err = cart.HandleAddItemToCart(eventStore, cmd2)
	fmt.Println("HandleAddItemToCart (cmd2) err=", err, ", events=", events)
}
