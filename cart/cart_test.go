package cart

import (
	"testing"
	"time"
)

type bogusEvent struct{}

func (be bogusEvent) AId() string {
	return "BOGUS"
}

func TestShouldFailIfNoApplierFound(t *testing.T) {
	bs := bogusEvent{}

	_, err := buildFromEvents("BOGUS", []event{bs})
	if err == nil {
		t.Error("Expected missing applier error was not returned")
	}
}

func TestBuildFromEvents(t *testing.T) {
	cartId := "Test-Cart-Id"

	createdEv := createdEvent{
		cartId,
		time.Now(),
	}
	itemAddedEv := itemAddedEvent{
		cartId,
		"Test-Item-Id",
		42,
	}

	events := []event{createdEv, itemAddedEv}

	c, err := buildFromEvents(cartId, events)
	if err != nil {
		t.Errorf("Building from events failed. This is bad.")
	}
	length := len(c.items)
	expectedLength := 1
	if length != expectedLength {
		t.Errorf("Expected item count :%d but got :%d", expectedLength, length)
		return
	}
	qty := c.items[0].qty
	expectedQty := 42
	if qty != expectedQty {
		t.Errorf("Expected item 0 quantity :%d but got :%d", expectedQty, qty)
	}
}
