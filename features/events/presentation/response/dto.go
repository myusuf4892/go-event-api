package response

import (
	"project3/eventapp/features/events"
	"time"
)

type Event struct {
	ID         int       `json:"id_event" form:"id_event"`
	Name       string    `json:"name" form:"name"`
	URL        string    `json:"image_url" form:"imageurl"`
	HostedBy   string    `json:"hostedby" form:"hostedby"`
	Performers string    `json:"performers" form:"performers"`
	Date       time.Time `json:"date" form:"date"`
	City       string    `json:"city" form:"city"`
	Location   string    `json:"location" form:"location"`
	Detail     string    `json:"details" form:"details"`
}

type EventByID struct {
	ID          int           `json:"id_event" form:"id_event"`
	Name        string        `json:"name" form:"name"`
	URL         string        `json:"image_url" form:"imageurl"`
	HostedBy    string        `json:"hostedby" form:"hostedby"`
	Performers  string        `json:"performers" form:"performers"`
	Date        time.Time     `json:"date" form:"date"`
	City        string        `json:"city" form:"city"`
	Location    string        `json:"location" form:"location"`
	Detail      string        `json:"details" form:"details"`
	Participant []Participant `json:"participant" form:"participant"`
}

type Participant struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Url  string `json:"url" form:"url"`
}

func FromCore(data events.Core) Event {
	return Event{
		ID:         data.ID,
		Name:       data.Name,
		Detail:     data.EventDetail,
		URL:        data.Url,
		HostedBy:   data.HostedBy,
		Performers: data.Performers,
		Date:       data.Date,
		City:       data.City,
		Location:   data.Location,
	}
}

func FromCoreList(data []events.Core) []Event {
	result := []Event{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCoreByID(data events.Core) EventByID {
	return EventByID{
		ID:          data.ID,
		Name:        data.Name,
		Detail:      data.EventDetail,
		URL:         data.Url,
		HostedBy:    data.HostedBy,
		Performers:  data.Performers,
		Date:        data.Date,
		City:        data.City,
		Location:    data.Location,
		Participant: FromParticipantCoreList(data.Participant),
	}
}

func FromParticipantCore(data events.Participant) Participant {
	return Participant{
		ID:   data.ID,
		Name: data.Name,
		Url:  data.Url,
	}
}

func FromParticipantCoreList(data []events.Participant) []Participant {
	result := []Participant{}
	for key := range data {
		result = append(result, FromParticipantCore(data[key]))
	}
	return result
}
