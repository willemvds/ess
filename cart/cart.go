package cart

import (
	"fmt"
)

type Item struct {
	Id string
}

type CartAggregate struct {
	Id    int
	Items []Item
}

func (ca CartAggregate) String() string {
	return fmt.Sprintf("I am Cart Aggregate %s with %d items", ca.Id, len(ca.Items))
}

func (ca *CartAggregate) AddItem(i Item) error {
	ca.ApplyItemAdded(ItemAddedToCart{ca.Id, i})
	return nil
}

func (ca *CartAggregate) ApplyItemAdded(ev ItemAddedToCart) {
	fmt.Println("Applying ItemAddedToCart event", ev, len(ca.Items))
	ca.Items = append(ca.Items, ev.Item)
}
