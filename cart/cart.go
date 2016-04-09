package cart

import (
	"fmt"
)

type Item struct {
	Id string
}

type CartAggregate struct {
	Id string
	Items []Item
}

func (ca CartAggregate) String() string {
	return fmt.Sprintf("I am Cart Aggregate %s with %d items", ca.Id, len(ca.Items))
}

func (ca *CartAggregate) AddToCart(i Item) error {
	ca.Items = append(ca.Items, i)
	return nil
}
