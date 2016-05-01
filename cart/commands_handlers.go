package cart

import (
	"errors"
)

var (
	ZeroEvents   []event
	ErrNoHandler = errors.New("[cart] No handler found for that command")
)

func init() {
	ZeroEvents = make([]event, 0)
}

func Handle(cmd command, es EventStore) ([]event, error) {
	switch cmd.(type) {
	case CreateCommand:
		createCommand := cmd.(CreateCommand)
		return handleCreateCommand(createCommand, es)
	case AddItemCommand:
		addItemCommand := cmd.(AddItemCommand)
		return handleAddItemCommand(addItemCommand, es)
	default:
		return nil, ErrNoHandler
	}
}

func handleCreateCommand(cmd CreateCommand, es EventStore) ([]event, error) {
	var version int
	var history []event
	var err error
	var crt *cart
	var events []event

	if version, history, err = es.EventsFor(cmd.CartId); err != nil {
		return ZeroEvents, err
	}

	if crt, err = buildFromEvents(cmd.CartId, history); err != nil {
		return ZeroEvents, err
	}

	if events, err = crt.create(cmd.CartId); err != nil {
		return ZeroEvents, err
	}

	if err = es.AppendFor(cmd.CartId, version, events); err != nil {
		return ZeroEvents, err
	}

	if err = crt.Apply(events); err != nil {
		return ZeroEvents, err
	}

	return events, nil
}

func handleAddItemCommand(cmd AddItemCommand, es EventStore) ([]event, error) {
	var version int
	var history []event
	var err error
	var crt *cart
	var events []event

	if version, history, err = es.EventsFor(cmd.CartId); err != nil {
		return ZeroEvents, err
	}

	if crt, err = buildFromEvents(cmd.CartId, history); err != nil {
		return ZeroEvents, err
	}

	if events, err = crt.addItem(cmd.ItemId, cmd.Qty); err != nil {
		return ZeroEvents, err
	}

	if err = es.AppendFor(cmd.CartId, version, events); err != nil {
		return ZeroEvents, err
	}

	if err = crt.Apply(events); err != nil {
		return ZeroEvents, err
	}

	return events, nil
}
