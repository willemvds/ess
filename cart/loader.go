package cart

import (
	"errors"
	"fmt"
)

var ErrCartNotFound = errors.New("Cart not found")

type CartLoader struct {
	Carts map[string]CartAggregate
}

func (cl CartLoader) Get(es *EventStore, id string) (*CartAggregate, error) {
	fmt.Println("Trying to load Cart Aggregate", id)
	defer func() {
		fmt.Println("Done trying to load Cart Aggregate", id)
	}()
	ca, ok := cl.Carts[id]
	cp := &ca
	if !ok {
		return &CartAggregate{}, ErrCartNotFound
	}
	for _, ev := range es.Events {
		switch ev.(type) {
		case ItemAddedToCart:
			ApplyItemAddedToCart(cp, ev.(ItemAddedToCart))
		}
	}
	return cp, nil
}
