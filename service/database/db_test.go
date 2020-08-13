package database

import (
	"fmt"
	"log"
	"testing"

	"github.com/Maleventum/arcana/model"
	"github.com/stretchr/testify/assert"
)

var testEvent = model.Event{Owner: 1, Name: "test event CRUD", Description: "making sure sql works"}

func TestEvent(t *testing.T) {
	db := New()
	defer db.DB.Close()

	err := db.CreateEvent(&testEvent)
	assert.Nil(t, err)

	events := db.GetEvent(0)
	log.Println(events)
	assert.NotEmpty(t, events)

	for _, evt := range events {
		fmt.Printf("deleting: %v\n", evt)
		err = db.DeleteEvent(evt.ID)
		assert.Nil(t, err)
	}

	straps := model.Event{Owner: 3, Name: "straps", Description: "aerial straps skills"}
	db.CreateEvent(&model.Event{Owner: 2, Name: "silks", Description: "aerial silks dance"})
	db.CreateEvent(&straps)
	db.CreateEvent(&model.Event{Owner: 4, Name: "pole", Description: "pole fitness"})

	byOwner := db.GetEvent(straps.Owner)
	straps.ID = byOwner[0].ID
	assert.Equal(t, straps, byOwner[0])
}

func TestEventSelect(t *testing.T) {
	db := New()
	defer db.DB.Close()

	events := db.GetEvent(0)
	log.Println(events)
}
