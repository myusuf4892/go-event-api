package response

import (
	_order "project/e-comerce/features/orders"
	"time"
)

type Order struct {
	ID     int       `json:"id" form:"id"`
	Time   time.Time `json:"time" form:"time"`
	Status string    `json:"status"`
	Price  int       `json:"price" form:"price"`
	UserID int       `json:"userid" form:"userid"`
}

type OrderDetail struct {
	ID          int    `json:"id" form:"id"`
	ProductName string `json:"productname" form:"Productname"`
	Url         string `json:"url" form:"url"`
	Qty         int    `json:"qty" form:"qty"`
	Price       int    `json:"price" form:"price"`
}

func OrderFromCore(data _order.Core) Order {
	return Order{
		ID:     data.ID,
		Time:   data.CreatedAt,
		Price:  data.Price,
		UserID: data.UserID,
		Status: data.Status,
	}
}

func FromCoreList(data []_order.Core) []Order {
	result := []Order{}
	for key := range data {
		result = append(result, OrderFromCore(data[key]))
	}
	return result
}

func OrderDetailFromCore(data _order.OrderDetail) OrderDetail {
	return OrderDetail{
		ID:          data.ID,
		ProductName: data.ProductName,
		Price:       data.Price,
	}
}

func FromOrderDetailCoreList(data []_order.OrderDetail) []OrderDetail {
	result := []OrderDetail{}
	for key := range data {
		result = append(result, OrderDetailFromCore(data[key]))
	}
	return result
}
