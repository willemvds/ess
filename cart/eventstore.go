package cart

type EventStore struct {
	Events []interface{}
}

func (es *EventStore) Add(ev interface{}) {
	es.Events = append(es.Events, ev)
}
