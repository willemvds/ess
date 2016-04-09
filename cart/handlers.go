package cart

import (
	"fmt"
)

func ApplyItemAddedToCart(ca *CartAggregate, ev ItemAddedToCart) {
	fmt.Println("Applying ItemAddedToCart event", ev)
}

func HandleAddItemToCart(es *EventStore, cmd AddItemToCart) ([]interface{}, error) {
	fmt.Println("Handling AddItemToCart command", cmd)

	cl := CartLoader{map[string]CartAggregate{"CartX2":CartAggregate{Id:"CartX2"}}}

	ca, err := cl.Get(es, cmd.CartId)
	if err != nil {
		return nil, err
	}

	item := Item{Id:"Cheats"}
	ca.AddToCart(item)

	ApplyItemAddedToCart(ca, ItemAddedToCart{cmd.CartId, cmd.ItemId})
	return []interface{}{ItemAddedToCart{cmd.CartId, cmd.ItemId}}, nil
}
