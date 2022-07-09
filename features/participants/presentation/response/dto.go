package response

import (
	"time"

	"project3/eventapp/features/participants"
)

type Participant struct {
	ID         int       `json:"id_participant" form:"id_participant"`
	Name       string    `json:"name" form:"name"`
	Detail     string    `json:"detail" form:"detail"`
	URL        string    `json:"url" form:"url"`
	Date       time.Time `json:"time" form:"time"`
	Performers string    `json:"performers" form:"performers"`
	HostedBy   string    `json:"hostedby" form:"hostedby"`
	City       string    `json:"city" form:"city"`
	Location   string    `json:"location" form:"location"`
}

func FromCore(core participants.Core) Participant {
	return Participant{
		ID:         core.ID,
		Name:       core.Event.Name,
		Detail:     core.Event.EventDetail,
		URL:        core.Event.Url,
		Date:       core.Event.Date,
		Performers: core.Event.Performers,
		HostedBy:   core.Event.HostedBy,
		City:       core.Event.City,
		Location:   core.Event.Location,
	}
}

func FromCoreList(data []participants.Core) []Participant {
	result := []Participant{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
