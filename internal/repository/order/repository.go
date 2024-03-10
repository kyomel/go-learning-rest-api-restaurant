package order

import (
	"context"
	"rest-api-restaurant/internal/model"
)

type Repository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderInfo(ctx context.Context, orderID string) (model.Order, error)
}
