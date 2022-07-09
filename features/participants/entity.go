package participants

import "time"

type Core struct {
	ID        int
	UserID    int
	EventID   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Event     Event
}

type Event struct {
	ID          int
	Url         string
	Name        string
	HostedBy    string
	Performers  string
	Date        time.Time
	City        string
	Location    string
	EventDetail string
	UserID      int
}

type Business interface {
	AddParticipant(data Core) error
	GetAllEventbyParticipant(idUser int) (data []Core, err error)
	DeleteParticipant(param, userID int) error
}

type Data interface {
	AddData(data Core) error
	SelectDataEvent(idUser int) (data []Core, err error)
	DeleteData(param, userID int) error
}
