package cart

import (
	"errors"
)

var (
	ErrNoHandler = errors.New("[cart] No handler found for that command")
)

func Handle(cmd command, es *EventStore) ([]event, error) {
	switch cmd.(type) {
	case CreateCommand:
		createCommand := cmd.(CreateCommand)
		cart, err := buildFromEvents(createCommand.CartId, es.EventsFor(createCommand.CartId))
		if err != nil {
			return nil, err
		}
		return handleCreateCommand(cart, createCommand)
	case AddItemCommand:
		addItemCommand := cmd.(AddItemCommand)
		cart, err := buildFromEvents(addItemCommand.CartId, es.EventsFor(addItemCommand.CartId))
		if err != nil {
			return nil, err
		}
		return handleAddItemCommand(cart, addItemCommand)
	default:
		return nil, ErrNoHandler
	}
}

func handleCreateCommand(cart *cart, cmd CreateCommand) ([]event, error) {
	return cart.create(cmd.CartId)
}

func handleAddItemCommand(cart *cart, cmd AddItemCommand) ([]event, error) {
	return cart.addItem(cmd.ItemId, cmd.Qty)
}
