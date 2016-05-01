package cart

import (
	"errors"
)

var (
	ErrNoHandler = errors.New("[cart] No handler found for that command")
)

func Handle(cmd command, es EventStore) ([]event, error) {
	switch cmd.(type) {
	case CreateCommand:
		createCommand := cmd.(CreateCommand)
		version, history, err := es.EventsFor(createCommand.CartId)
		if err != nil {
			return []event{}, err
		}
		cart, err := buildFromEvents(createCommand.CartId, history)
		if err != nil {
			return []event{}, err
		}
		events, err := handleCreateCommand(cart, createCommand)
		if err != nil {
			return []event{}, err
		}
		err = es.AppendFor(createCommand.CartId, version, events)
		if err != nil {
			return []event{}, err
		}
		return events, nil
	case AddItemCommand:
		addItemCommand := cmd.(AddItemCommand)
		version, history, err := es.EventsFor(addItemCommand.CartId)
		if err != nil {
			return []event{}, err
		}
		cart, err := buildFromEvents(addItemCommand.CartId, history)
		if err != nil {
			return []event{}, err
		}
		events, err := handleAddItemCommand(cart, addItemCommand)
		if err != nil {
			return []event{}, err
		}
		err = es.AppendFor(addItemCommand.CartId, version, events)
		if err != nil {
			return []event{}, err
		}
		return events, nil
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
