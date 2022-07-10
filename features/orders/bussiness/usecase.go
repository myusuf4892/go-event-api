package bussiness

import (
	"errors"
	"log"
	"project/e-comerce/features/carts"
	"project/e-comerce/features/orders"
)

type orderUseCase struct {
	orderData orders.Data
	cartData  carts.Data
}

func NewOrderBusiness(orderData orders.Data, cartData carts.Data) orders.Bussiness {
	return &orderUseCase{
		orderData: orderData,
		cartData:  cartData,
	}
}

func (uc orderUseCase) AddOrder(core orders.Core, cartID []int, userId int) error {
	dataDetailOrder, productId, totalPrice, errGetDetailOrder := uc.orderData.SelectCart(cartID, userId)
	if errGetDetailOrder != nil {
		log.Print(errGetDetailOrder)
		return errGetDetailOrder
	}

	var remain_stock_list []int
	for i, val := range productId {
		stockProduct, errCheckStock := uc.orderData.SelectProduct(val)
		if errCheckStock != nil {
			log.Print(errCheckStock)
			return errCheckStock
		}
		remain_stock := stockProduct - dataDetailOrder[i].Qty
		remain_stock_list = append(remain_stock_list, remain_stock)
		if remain_stock < 0 {
			return errors.New("product stock is not enough")
		}
	}

	addID, errAdd := uc.orderData.InsertAddress(core.Address)
	if errAdd != nil {
		log.Print(errAdd)
		return errAdd
	}

	// payID, errPay := uc.orderData.InsertPayment(core.Payment)
	// if errPay != nil {
	// 	log.Print(errPay)
	// 	return errPay
	// }

	orderID, errOrder := uc.orderData.InsertOrder(core, addID)
	if errOrder != nil {
		log.Print(errOrder)
		return errOrder
	}

	errOrderDetail := uc.orderData.InsertOrderDetail(orderID, dataDetailOrder)
	if errOrderDetail != nil {
		log.Print(errOrderDetail)
		return errOrderDetail
	}

	errUpdateOrder := uc.orderData.UpdatePriceOrder(orderID, totalPrice)
	if errUpdateOrder != nil {
		log.Print(errUpdateOrder)
		return errUpdateOrder
	}

	for i, val := range productId {
		errUpdateProduct := uc.orderData.UpdateStockProduct(val, remain_stock_list[i])
		if errUpdateProduct != nil {
			log.Print(errUpdateProduct)
			return errUpdateProduct
		}
	}

	errDeleteCart := uc.cartData.DestroyAll(userId, cartID)
	if errDeleteCart != nil {
		log.Print(errDeleteCart)
		return errDeleteCart
	}

	return nil
}

func (uc orderUseCase) GetOrder(userId int) ([]orders.Core, error) {
	respData, err := uc.orderData.SelectOrder(userId)
	return respData, err
}

func (uc orderUseCase) GetOrderDetail(orderId int) ([]orders.OrderDetail, error) {
	respData, err := uc.orderData.SelectOrderDetailByOrderID(orderId)
	return respData, err
}

func (uc orderUseCase) ConfirmOrder(paymentCore orders.PaymentCore, orderId int, userId int) error {
	paymentCore.OrderID = orderId
	respDataPay := uc.orderData.InsertPayment(paymentCore)
	if respDataPay != nil {
		log.Print(respDataPay)
		return respDataPay
	}

	respDataUpdate := uc.orderData.UpdateStatusOrder(orderId, userId, "Paid")
	if respDataUpdate != nil {
		log.Print(respDataUpdate)
		return respDataUpdate
	}
	return nil
}

func (uc orderUseCase) CancelOrder(orderId int, userId int) error {
	respDataUpdate := uc.orderData.UpdateStatusOrder(orderId, userId, "Canceled")
	if respDataUpdate != nil {
		log.Print(respDataUpdate)
		return respDataUpdate
	}
	return nil
}
