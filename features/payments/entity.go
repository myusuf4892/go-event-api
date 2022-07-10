package payments

import (
	_user "project3/eventapp/features/users/data"
	"time"
)

type Core struct {
	ID        int
	Name      string
	Details   string
	CreatedAt time.Time
	User      _user.User
	Address   Address
	Payment   Payment_method_detail
}

type Address struct {
	ID         int
	Street     string
	PostalCode int
	Province   string
	City       string
	User       _user.User
}

type Payment_method struct {
	ID   int
	Name string
}

type Payment_method_detail struct {
	ID      string
	Details string
	Payment Payment_method
}

type Business interface {
}

type Data interface {
}
