package model

// swagger:parameters EventParam

// Event description object
type Event struct {
	// Body of an event object.
	// in:body
	// out:body
	ID          int64  `db:"id" json:"id"`
	Owner       int64  `db:"owner" json:"owner"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}
