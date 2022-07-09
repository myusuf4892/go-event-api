package events

import (
	"time"
)

type Core struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	EventName   string
	EventDetail string
	Url         string
	Date        time.Time
	Performers  string
	HostedBy    string
	City        string
	Location    string
	IDUser      int
	Participant []Participant
}

type Participant struct {
	ID   int
	Name string
	Url  string
}

type Business interface {
	GetAllEvent(limit int, offset int, city string, name string) (data []Core, total int64, err error)
	GetEventByID(param int) (data Core, err error)
	InsertEvent(dataReq Core) (err error)
	DeleteEventByID(id int, userId int) (err error)
	UpdateEventByID(dataReq Core, id int, userId int) (err error)
	GetEventByUserID(id_user, limit, offset int) (data []Core, total int64, err error)
}

type Data interface {
	SelectData(limit int, offset int, city string, name string) (data []Core, total int64, err error)
	SelectDataByID(param int) (data Core, err error)
	InsertData(dataReq Core) (err error)
	DeleteDataByID(id int, userId int) (err error)
	UpdateDataByID(dataReq map[string]interface{}, id int, userId int) (err error)
	SelectDataByUserID(id_user, limit, offset int) (data []Core, total int64, err error)
	SelectParticipantData(id_event int) (data []Participant, err error)
}
