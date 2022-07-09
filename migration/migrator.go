package migration

import (
	_mComment "project3/eventapp/features/comments/data"
	_mEvent "project3/eventapp/features/events/data"
	_mParticipant "project3/eventapp/features/participants/data"
	_mUser "project3/eventapp/features/users/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(_mUser.User{})
	db.AutoMigrate(_mEvent.Event{})
	db.AutoMigrate(_mParticipant.Participant{})
	db.AutoMigrate(_mComment.Comment{})
}
