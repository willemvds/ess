package cart

import (
	"fmt"
	"sync"
)

type EventStore interface {
	AppendFor(cartId string, appendVersion int, events []event) error
	EventsFor(cartId string) (int, []event, error)
}

type ArrayEventStore struct {
	events      []event
	appendMutex sync.Mutex
}

func (store *ArrayEventStore) AppendFor(cartId string, appendVersion int, events []event) error {
	currentVersion := 0
	for _, event := range store.events {
		if event.AId() != cartId {
			continue
		}
		currentVersion++
	}

	if currentVersion != appendVersion {
		return fmt.Errorf("Version mismatch for :%s. appendVersion = :%d vs currentVersion = :%d", cartId, appendVersion, currentVersion)
	}

	store.appendMutex.Lock()
	defer store.appendMutex.Unlock()
	store.events = append(store.events, events...)
	return nil
}

func (store ArrayEventStore) EventsFor(cartId string) (int, []event, error) {
	version := 0
	eventsFor := make([]event, 0)
	for _, event := range store.events {
		if event.AId() != cartId {
			continue
		}
		eventsFor = append(eventsFor, event)
		version++
	}
	return version, eventsFor, nil
}
