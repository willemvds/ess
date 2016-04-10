package cart

import (
	"fmt"
)

func HandleAddItemToCart(es *EventStore, cmd AddItemToCart) ([]interface{}, error) {
	fmt.Println("Handling AddItemToCart command", cmd)

	cl := CartLoader{map[string]CartAggregate{"CartX2": CartAggregate{Id: "CartX2"}}}

	ca, err := cl.Get(es, cmd.CartId)
	if err != nil {
		return nil, err
	}

	item := Item{Id: "Cheats"}
	ca.AddToCart(item)

	return nil, nil
}
