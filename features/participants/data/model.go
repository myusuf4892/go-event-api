package data

import (
	"project3/eventapp/features/events/data"
	"project3/eventapp/features/participants"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	UserID  int `json:"user_id" form:"user_id"`
	EventID int `json:"event_id" form:"event_id"`
	Event   data.Event
}

func fromCore(core participants.Core) Participant {
	return Participant{
		UserID:  core.UserID,
		EventID: core.EventID,
	}
}

func (data *Participant) toCore() participants.Core {
	return participants.Core{
		ID: int(data.EventID),
		Event: participants.Event{
			Url:         data.Event.URL,
			Name:        data.Event.Name,
			HostedBy:    data.Event.HostedBy,
			Performers:  data.Event.Performers,
			Date:        data.Event.Date,
			City:        data.Event.City,
			Location:    data.Event.Location,
			EventDetail: data.Event.Detail,
		},
	}
}

func ToCoreList(data []Participant) []participants.Core {
	result := []participants.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}
