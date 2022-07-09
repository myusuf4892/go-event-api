package data

import (
	"project3/eventapp/features/events"
	"project3/eventapp/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name       string    `json:"name" form:"name"`
	Detail     string    `json:"detail" form:"detail"`
	URL        string    `json:"url" form:"url"`
	Date       time.Time `json:"date" form:"date"`
	Performers string    `json:"performers" form:"performers"`
	HostedBy   string    `json:"hostedby" form:"hostedby"`
	City       string    `json:"city" form:"city"`
	Location   string    `json:"location" form:"location"`
	UserID     int       `json:"user_id" form:"user_id"`
	User       data.User
}

type Participant struct {
	ID       int    `json:"participant_id" form:"participant_id"`
	User_ID  int    `json:"user_id" form:"user_id"`
	Event_ID int    `json:"event_id" form:"event_id"`
	Name     string `json:"name" form:"name"`
	Url      string `json:"url" form:"url"`
	User     data.User
}

//DTO

func (data *Event) toCore() events.Core {
	return events.Core{
		ID:          int(data.ID),
		Name:        data.Name,
		EventDetail: data.Detail,
		Url:         data.URL,
		City:        data.City,
		HostedBy:    data.HostedBy,
		Performers:  data.Performers,
		Location:    data.Location,
		Date:        data.Date,
		IDUser:      data.UserID,
	}
}

func ToCoreList(data []Event) []events.Core {
	result := []events.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core events.Core) Event {
	return Event{
		Name:       core.Name,
		Detail:     core.EventDetail,
		URL:        core.Url,
		HostedBy:   core.HostedBy,
		Performers: core.Performers,
		Location:   core.Location,
		Date:       core.Date,
		City:       core.City,
		UserID:     core.IDUser,
	}
}

func toCore(data Event) events.Core {
	return data.toCore()
}

func (data *Participant) toParticipantCore() events.Participant {
	return events.Participant{
		ID:   data.ID,
		Name: data.User.Name,
		Url:  data.User.URL,
	}
}

func ToParticipantCoreList(data []Participant) []events.Participant {
	result := []events.Participant{}
	for key := range data {
		result = append(result, data[key].toParticipantCore())
	}
	return result
}
