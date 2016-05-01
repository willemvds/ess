package cart

type EventStore struct {
	events []event
}

func (es *EventStore) Add(events []event) {
	es.events = append(es.events, events...)
}

func (es EventStore) EventsFor(cartId string) []event {
	eventsFor := make([]event, 0)
	for _, event := range es.events {
		if event.AId() != cartId {
			continue
		}
		eventsFor = append(eventsFor, event)
	}
	return eventsFor
}
