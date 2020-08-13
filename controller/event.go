package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Maleventum/arcana/model"
)

// EventUsecase business rules to process events
type EventUsecase interface {
	Create(event *model.Event) error
}

// EventController for saving the interface for events
type EventController struct {
	eventUsecase EventUsecase
}

// NewEvent Controller for event related CRUD operations
func NewEvent(eventUsecase EventUsecase) *EventController {
	return &EventController{eventUsecase: eventUsecase}
}

// swagger:route POST /v1/user/{UserID}/event event EventCreate
// Creates a new event.
//
// responses:
//   500: Internal Server Error
//   400: Bad Request
//   201: EventResponse

// Create Event create entry point
// curl -v -X POST 127.0.0.1:8081/v1/user/PHONE/event -H "Content-Type: application/json" --data '{}'
func (c *EventController) Create(resp http.ResponseWriter, req *http.Request) {
	event := model.Event{}
	err := unmarshalBody(&event, resp, req)
	if err != nil {
		return
	}

	err = c.eventUsecase.Create(&event)

	if err != nil {
		resp.WriteHeader(http.StatusForbidden)
		resp.Write([]byte(err.Error()))
	}
	resp.WriteHeader(http.StatusCreated)
}

func unmarshalBody(v interface{}, resp http.ResponseWriter, req *http.Request) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(err.Error()))
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(err.Error()))
		return err
	}

	return nil
}
