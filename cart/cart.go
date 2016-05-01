package cart

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrEventApplierNotFound = errors.New("[cart] No applier found for that event")
	ErrInvalidQuantity      = errors.New("Invalid quantity")
)

type item struct {
	id  string
	qty int
}

type cart struct {
	id      string
	created time.Time
	items   []item
}

func (cart cart) String() string {
	return fmt.Sprintf("I am Cart :%s with :%d items", cart.id, len(cart.items))
}

func buildFromEvents(cartId string, events []event) (*cart, error) {
	c := &cart{id: cartId}
	for _, event := range events {
		err := c.Apply(event)
		if err != nil {
			return &cart{}, err
		}
	}
	return c, nil
}

func (cart *cart) Apply(ev event) error {
	switch ev.(type) {
	case createdEvent:
		cart.applyCreated(ev.(createdEvent))
		return nil
	case itemAddedEvent:
		cart.applyItemAdded(ev.(itemAddedEvent))
		return nil
	}
	return ErrEventApplierNotFound
}

func (cart *cart) create(id string) ([]event, error) {
	created := time.Now()
	ev := createdEvent{id, created}
	// cart.applyCreated(ev)
	return []event{ev}, nil
}

func (cart *cart) applyCreated(ev createdEvent) {
	cart.created = ev.Created
}

func (cart *cart) addItem(itemId string, qty int) ([]event, error) {
	if qty <= 0 {
		return []event{}, ErrInvalidQuantity
	}

	// is this where we read from another aggregate like product to see if item exists?

	ev := itemAddedEvent{cart.id, itemId, qty}
	// cart.applyItemAdded(ev)
	return []event{ev}, nil
}

func (cart *cart) applyItemAdded(ev itemAddedEvent) {
	cart.items = append(cart.items, item{
		id:  ev.ItemId,
		qty: ev.Qty,
	})
}
