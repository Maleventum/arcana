package usecase

import (
	"fmt"

	"github.com/Maleventum/arcana/model"
)

// EventStorage object to communicate with the chosen storage option
type EventStorage interface {
	CreateEvent(event *model.Event) error
}

// Event objec for processing events business rules
type Event struct {
	storage EventStorage
}

// NewEvent creates an object for processing events
func NewEvent(eventStorage EventStorage) *Event {
	return &Event{storage: eventStorage}
}

// Create Valitades business rules before creating an event
func (u *Event) Create(event *model.Event) error {
	fmt.Println(*event)
	err := u.storage.CreateEvent(event)
	return err
}
