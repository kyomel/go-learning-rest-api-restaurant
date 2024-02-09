package constant

import "rest-api-restaurant/internal/model"

const (
	OrderStatusProcessed model.OrderStatus = "Processed"
	OrderStatusFinished  model.OrderStatus = "Finished"
	OrderStatusFailed    model.OrderStatus = "Failed"
)

const (
	ProductOrderStatusPreparing model.ProductOrderStatus = "Preparing"
	ProductOrderStatusFinished  model.ProductOrderStatus = "Finished"
)
