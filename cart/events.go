package cart

import (
	"time"
)

type event interface {
	AId() string
}

type createdEvent struct {
	CartId  string
	Created time.Time
}

func (ev createdEvent) AId() string {
	return ev.CartId
}

type itemAddedEvent struct {
	CartId string
	ItemId string
	Qty    int
}

func (ev itemAddedEvent) AId() string {
	return ev.CartId
}

type itemRemovedEvent struct {
	CartId string
	ItemId string
	Qty    int
}

func (ev itemRemovedEvent) AId() string {
	return ev.CartId
}
