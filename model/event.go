package model

// Event description object
//
// swagger:model EventObj
type Event struct {
	// Identifier generated by the storage TODO: remove
	ID int64 `db:"id" json:"id"`

	// Number of the owner of the event
	//
	// required: true
	Owner int64 `db:"owner" json:"owner"`

	// Desired event name
	//
	// required: true
	Name string `db:"name" json:"name"`

	// More detailed description of the event
	Description string `db:"description" json:"description"`
}

// EventResponse description object response
//
// swagger:response EventResponse
type EventResponse struct {
	// in:body
	Event Event
}

// EventParams -
//
// swagger:parameters EventCreate
type EventParams struct {
	// a EventParams describes an object to be created

	// in: body
	// required: true
	Event Event

	// in: path
	// required: true
	UserID int64
}
