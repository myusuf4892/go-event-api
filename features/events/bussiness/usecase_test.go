package bussiness

import (
	"fmt"
	"project3/eventapp/features/events"
	"testing"

	"github.com/stretchr/testify/assert"
)

//mock data success case
type mockEventDataSucces struct{}

func (mock mockEventDataSucces) SelectData(limit int, offset int, name string, sity string) (data []events.Core, total int64, err error) {
	return []events.Core{
		{ID: 1, Name: "festival 1", EventDetail: "detail festival 1", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1},
		{ID: 2, Name: "festival 2", EventDetail: "detail festival 2", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1},
	}, 1, nil
}

func (mock mockEventDataSucces) SelectDataByID(id int) (data events.Core, err error) {
	return events.Core{ID: 1, Name: "festival 1", EventDetail: "detail festival 1", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1}, nil
}

func (mock mockEventDataSucces) InsertData(data events.Core) (err error) {
	return nil
}

func (mock mockEventDataSucces) DeleteDataByID(id int, userid int) (err error) {
	return nil
}

func (mock mockEventDataSucces) UpdateDataByID(dataReq map[string]interface{}, id int, userId int) (err error) {
	return nil
}

func (mock mockEventDataSucces) SelectDataByUserID(userId int, limit int, offset int) (data []events.Core, total int64, err error) {
	return []events.Core{
		{ID: 1, Name: "festival 1", EventDetail: "detail festival 1", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1},
		{ID: 2, Name: "festival 2", EventDetail: "detail festival 2", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1},
	}, 1, nil
}

func (mock mockEventDataSucces) SelectParticipantData(limit int) (data []events.Participant, err error) {
	return []events.Participant{
		{ID: 1, Name: "user 1", Url: "example.com"},
		{ID: 2, Name: "user 2", Url: "example.com"},
	}, nil
}

//mock data failed case
type mockEventDataFailed struct{}

func (mock mockEventDataFailed) SelectData(limit int, offset int, name string, sity string) (data []events.Core, total int64, err error) {
	return nil, 0, fmt.Errorf("Failed to select data")
}

func (mock mockEventDataFailed) SelectDataByID(id int) (data events.Core, err error) {
	return data, fmt.Errorf("Failed to select data")
}

func (mock mockEventDataFailed) InsertData(data events.Core) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockEventDataFailed) DeleteDataByID(id int, userid int) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockEventDataFailed) UpdateDataByID(dataReq map[string]interface{}, id int, userId int) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockEventDataFailed) SelectDataByUserID(userId int, limit int, offset int) (data []events.Core, total int64, err error) {
	return nil, 0, fmt.Errorf("failed to select data ")
}

func (mock mockEventDataFailed) SelectParticipantData(event_id int) (data []events.Participant, err error) {
	return nil, fmt.Errorf("failed get participant")
}

func TestGetAllEvent(t *testing.T) {
	t.Run("Test Get All Data Success", func(t *testing.T) {
		limit := 2
		page := 1
		city := "malang"
		name := "fest"
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		result, total, err := eventBusiness.GetAllEvent(limit, page, name, city)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), total)
		assert.Equal(t, "festival 1", result[0].Name)
	})

	t.Run("Test Get All Data Failed", func(t *testing.T) {
		limit := 20
		page := 1
		city := "malang"
		name := "fest"
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		result, total, err := eventBusiness.GetAllEvent(limit, page, name, city)
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, int64(1), total)
	})
}

func TestGetEventByID(t *testing.T) {
	t.Run("Test Get event Data By ID Success", func(t *testing.T) {
		id := 1
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		result, err := eventBusiness.GetEventByID(id)
		assert.Nil(t, err)
		assert.Equal(t, "festival 1", result.Name)
	})

	t.Run("Test Get event Data By ID Failed", func(t *testing.T) {
		id := 3
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		result, err := eventBusiness.GetEventByID(id)
		assert.NotNil(t, err)
		assert.Equal(t, "", result.Name)
	})
}

func TestInsertEvent(t *testing.T) {
	t.Run("Test Insert Data Success", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		newEvent := events.Core{
			Name: "festival 1", EventDetail: "detail festival 1", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1, Url: "example.com",
		}
		err := eventBusiness.InsertEvent(newEvent)
		assert.Nil(t, err)
	})

	t.Run("Test Insert Data Failed", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		newEvent := events.Core{
			Name: "alta",
		}
		err := eventBusiness.InsertEvent(newEvent)
		assert.NotNil(t, err)
	})
}

func TestGetEventByUserID(t *testing.T) {
	t.Run("Test Get Event Data By ID User Success", func(t *testing.T) {
		id := 1
		limit := 2
		offset := 1
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		result, total, err := eventBusiness.GetEventByUserID(id, limit, offset)
		assert.Nil(t, err)
		assert.Equal(t, "festival 1", result[0].Name)
		assert.Equal(t, int64(1), total)
	})

	t.Run("Test Get Data By ID User Failed", func(t *testing.T) {
		id := 30
		limit := 2
		offset := 0
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		result, total, err := eventBusiness.GetEventByUserID(id, limit, offset)
		assert.NotNil(t, err)
		assert.Nil(t, result)
		assert.Equal(t, int64(1), total)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Test Delete Data", func(t *testing.T) {
		id := 1
		userid := 1
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		err := eventBusiness.DeleteEventByID(id, userid)
		assert.Nil(t, err)

	})
	t.Run("Test Delete Data Failed", func(t *testing.T) {
		id := 3
		userid := 1
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		err := eventBusiness.DeleteEventByID(id, userid)
		assert.NotNil(t, err)

	})
}

func TestUpdateEvent(t *testing.T) {
	t.Run("Test Update Data Success", func(t *testing.T) {
		eventBusiness := NewEventBusiness(mockEventDataSucces{})
		id := 1
		userid := 1
		newEvent := events.Core{
			Name: "festival 1", EventDetail: "detail festival 1", Location: "1,1", HostedBy: "me", Performers: "me band", City: "malang", IDUser: 1,
		}
		err := eventBusiness.UpdateEventByID(newEvent, id, userid)
		assert.Nil(t, err)
	})

	t.Run("Test Update Data Failed", func(t *testing.T) {
		id := 1
		userid := 0
		eventBusiness := NewEventBusiness(mockEventDataFailed{})
		newEvent := events.Core{
			Name: "festival abs",
		}
		err := eventBusiness.UpdateEventByID(newEvent, id, userid)
		assert.NotNil(t, err)
	})
}
