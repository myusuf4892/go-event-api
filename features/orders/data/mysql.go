package data

import (
	"errors"
	"project/e-comerce/features/carts/data"
	"project/e-comerce/features/orders"

	"gorm.io/gorm"
)

type mysqlOrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(conn *gorm.DB) orders.Data {
	return &mysqlOrderRepository{
		db: conn,
	}
}

func (repo *mysqlOrderRepository) InsertAddress(address orders.AddressCore) (int, error) {
	addressModel := fromAddressCore(address)
	result := repo.db.Create(&addressModel)

	if result.Error != nil || result.RowsAffected == 0 {
		return 0, errors.New("failed add address")
	}
	return addressModel.ID, nil
}

func (repo *mysqlOrderRepository) InsertOrder(order orders.Core, addID int) (int, error) {
	orderModel := fromCore(order)
	// orderModel.PaymentID = payID
	orderModel.AddressID = addID
	result := repo.db.Create(&orderModel)

	if result.Error != nil || result.RowsAffected == 0 {
		return 0, errors.New("failed add order")
	}
	return int(orderModel.ID), nil
}

func (repo *mysqlOrderRepository) InsertOrderDetail(orderID int, orderDetail []orders.OrderDetail) error {
	orderDetailModel := fromOrderDetailCoreList(orderDetail)
	for i := 0; i < len(orderDetailModel); i++ {
		orderDetailModel[i].OrderID = orderID
	}
	result := repo.db.Create(&orderDetailModel)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("failed add order")
	}
	return nil
}

func (repo *mysqlOrderRepository) SelectOrder(userId int) ([]orders.Core, error) {
	orderData := []Order{}
	result := repo.db.Preload("Address").Find(&orderData, "user_id = ?", userId)
	if result.Error != nil {
		return []orders.Core{}, errors.New("failed get orders")
	}
	orderDataList := ToCoreList(orderData)

	return orderDataList, nil
}

func (repo *mysqlOrderRepository) SelectOrderDetailByOrderID(orderId int) ([]orders.OrderDetail, error) {
	orderDetailData := []OrderDetail{}
	result := repo.db.Find(&orderDetailData, "order_id = ?", orderId)
	if result.Error != nil {
		return []orders.OrderDetail{}, errors.New("failed get orders")
	}

	orderDataList := ToOrderDetailCoreList(orderDetailData)

	return orderDataList, nil
}

func (repo *mysqlOrderRepository) SelectCart(cartId []int, userId int) ([]orders.OrderDetail, []int, int, error) {
	cartData := []Cart{}
	totalPrice := 0
	var productid []int
	result := repo.db.Where("user_id = ?", userId).Preload("Product").Find(&cartData, cartId)
	if result.Error != nil {
		return []orders.OrderDetail{}, productid, 0, errors.New("failed get orders")
	}
	if len(cartData) == 0 {
		return []orders.OrderDetail{}, productid, 0, errors.New("cart not found for this user")
	}

	for i := 0; i < len(cartData); i++ {
		cartData[i].Product.Price = cartData[i].Product.Price * cartData[i].Qty
		totalPrice += cartData[i].Product.Price
		productid = append(productid, cartData[i].ProductID)
	}

	orderDetailDataList := ToOrderDetailCoreListFromCart(cartData)
	return orderDetailDataList, productid, totalPrice, nil
}

func (repo *mysqlOrderRepository) UpdatePriceOrder(orderId int, price int) error {
	order := Order{}
	order.ID = uint(orderId)
	result := repo.db.Model(order).Updates(map[string]interface{}{"price": price, "status": "Unpaid"})
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("failed update price orders")
	}
	return nil
}

func (repo *mysqlOrderRepository) SelectProduct(productId int) (qty int, err error) {
	product := data.Product{}
	product.ID = uint(productId)
	result := repo.db.First(&product)
	if result.Error != nil || result.RowsAffected == 0 {
		return 0, errors.New("failed update price orders")
	}
	return product.Stock, nil
}

func (repo *mysqlOrderRepository) UpdateStockProduct(productid int, qty int) error {
	product := data.Product{}

	result := repo.db.Model(product).Where("id = ?", productid).Update("stock", qty)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("failed update stock product")
	}
	return nil

}

func (repo *mysqlOrderRepository) UpdateStatusOrder(orderId int, userId int, status string) error {
	order := Order{}
	order.ID = uint(orderId)
	result := repo.db.Model(order).Update("status", status).Where("user_id = ?", userId)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("failed update status orders")
	}
	return nil
}

func (repo *mysqlOrderRepository) InsertPayment(payment orders.PaymentCore) error {
	paymentModel := fromPaymentCore(payment)
	result := repo.db.Create(&paymentModel)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("failed add payment")
	}
	return nil
}
