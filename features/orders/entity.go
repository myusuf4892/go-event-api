package orders

import (
	"time"
)

type Core struct {
	ID        int
	UserID    int
	AddressID int
	Price     int
	Status    string
	Address   AddressCore
	CreatedAt time.Time
}

type OrderDetail struct {
	ID          int
	OrderID     int
	ProductName string
	Price       int
	Qty         int
}

type AddressCore struct {
	ID       int
	Receiver string
	Phone    string
	Address  string
}

type PaymentCore struct {
	ID          int
	PaymentName string
	NumberCard  string
	PaymentCode string
	OrderID     int
}

type Bussiness interface {
	AddOrder(core Core, orderdetailid []int, userId int) error
	GetOrder(orderId int) (orderData []Core, err error)
	GetOrderDetail(orderId int) ([]OrderDetail, error)
	ConfirmOrder(core PaymentCore, orderId int, userId int) error
	CancelOrder(orderId int, userId int) error
}

type Data interface {
	InsertOrder(core Core, addID int) (orderID int, err error)
	InsertAddress(AddressCore) (addID int, err error)
	InsertPayment(PaymentCore) (err error)
	InsertOrderDetail(int, []OrderDetail) error
	SelectOrder(int) ([]Core, error)
	SelectCart(cartId []int, userId int) (orderDetail []OrderDetail, productId []int, Price int, err error)
	SelectOrderDetailByOrderID(orderID int) ([]OrderDetail, error)
	UpdatePriceOrder(orderid int, price int) error
	SelectProduct(productid int) (stock int, err error)
	UpdateStockProduct(productid int, qty int) error
	UpdateStatusOrder(orderId int, userId int, status string) error
}
