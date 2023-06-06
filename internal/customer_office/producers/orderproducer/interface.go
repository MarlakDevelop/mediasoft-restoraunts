package orderproducer

import customer "restaurant/pkg/contracts/customer_office"

type OrderProducer interface {
	SendOrder(order *customer.OrderJSON)
}
