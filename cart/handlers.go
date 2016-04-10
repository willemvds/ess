package cart

import (
	"fmt"
)

func HandleAddItemToCart(es *EventStore, cmd AddItemToCart) ([]interface{}, error) {
	fmt.Println("Handling AddItemToCart command", cmd)

	cl := CartLoader{map[int]CartAggregate{42: CartAggregate{Id: 42}}}

	ca, err := cl.Get(es, cmd.CartId)
	if err != nil {
		return nil, err
	}

	item := Item{Id: "Cheats"}
	ca.AddItem(item)

	return nil, nil
}
