package data

import (
	_mProduct "project/e-comerce/features/carts/data"
	"project/e-comerce/features/orders"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    int
	Price     int
	Status    string
	AddressID int
	Address   Address
}

type OrderDetail struct {
	ID          int
	OrderID     int
	ProductName string
	Price       int
	Url         string
	Qty         int
}

type Cart struct {
	gorm.Model
	ProductID int
	UserID    int
	Qty       int
	Product   _mProduct.Product
}

type Address struct {
	ID       int
	Receiver string
	Phone    string
	Address  string
}

type Payment struct {
	ID          int
	PaymentName string
	NumberCard  string
	PaymentCode string
	orderID     int
}

func (data *Order) toCore() orders.Core {
	return orders.Core{
		ID:        int(data.ID),
		Price:     data.Price,
		Status:    data.Status,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
		AddressID: data.AddressID,
		Address: orders.AddressCore{
			ID:       data.Address.ID,
			Receiver: data.Address.Receiver,
			Phone:    data.Address.Phone,
			Address:  data.Address.Address,
		},
	}
}

func ToCoreList(data []Order) []orders.Core {
	result := []orders.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *OrderDetail) toOrderDetailCore() orders.OrderDetail {
	return orders.OrderDetail{
		ID:          int(data.ID),
		OrderID:     data.OrderID,
		Price:       data.Price,
		Qty:         data.Qty,
		ProductName: data.ProductName,
	}
}

func ToOrderDetailCoreList(data []OrderDetail) []orders.OrderDetail {
	result := []orders.OrderDetail{}
	for key := range data {
		result = append(result, data[key].toOrderDetailCore())
	}
	return result
}

func (data *Cart) toOrderDetailFromCart() orders.OrderDetail {
	return orders.OrderDetail{
		ID:          int(data.ID),
		Qty:         data.Qty,
		ProductName: data.Product.Name,
		Price:       data.Product.Price,
	}
}

func ToOrderDetailCoreListFromCart(data []Cart) []orders.OrderDetail {
	result := []orders.OrderDetail{}
	for key := range data {
		result = append(result, data[key].toOrderDetailFromCart())
	}
	return result
}

func fromCore(core orders.Core) Order {
	return Order{
		AddressID: core.AddressID,
		Price:     core.Price,
		UserID:    core.UserID,
	}
}

func fromOrderDetailCore(orderDetailCore orders.OrderDetail) OrderDetail {
	return OrderDetail{
		ProductName: orderDetailCore.ProductName,
		Price:       orderDetailCore.Price,
		OrderID:     orderDetailCore.OrderID,
		Qty:         orderDetailCore.Qty,
	}
}

func fromOrderDetailCoreList(data []orders.OrderDetail) []OrderDetail {
	result := []OrderDetail{}
	for key := range data {
		result = append(result, fromOrderDetailCore(data[key]))
	}
	return result
}

func fromAddressCore(addressCore orders.AddressCore) Address {
	return Address{
		Receiver: addressCore.Receiver,
		Phone:    addressCore.Phone,
		Address:  addressCore.Address,
	}
}

func fromPaymentCore(paymentCore orders.PaymentCore) Payment {
	return Payment{
		NumberCard:  paymentCore.NumberCard,
		PaymentCode: paymentCore.PaymentCode,
		PaymentName: paymentCore.PaymentName,
		orderID:     paymentCore.OrderID,
	}
}
