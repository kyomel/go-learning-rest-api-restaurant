package resto

import (
	"errors"
	"rest-api-restaurant/internal/model"
	"rest-api-restaurant/internal/model/constant"
	"rest-api-restaurant/internal/repository/menu"
	"rest-api-restaurant/internal/repository/order"
	"rest-api-restaurant/internal/repository/user"
	"rest-api-restaurant/internal/tracing"

	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
	userRepo  user.Repository
}

func GetUseCase(menuRepo menu.Repository, orderRepo order.Repository, userRepo user.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (r *restoUsecase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()

	return r.menuRepo.GetMenuList(menuType)
}

func (r *restoUsecase) Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "Order")
	defer span.End()

	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}

		productOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  menuData.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: int64(menuData.Price) * (int64(orderProduct.Quantity)),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}

	orderData := model.Order{
		ID:            uuid.New().String(),
		UserID:        request.UserID,
		Status:        constant.OrderStatusProcessed,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}

	createdOrderData, err := r.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, err
	}

	return createdOrderData, err
}
func (r *restoUsecase) GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetOrderInfo")
	defer span.End()

	orderData, err := r.orderRepo.GetOrderInfo(ctx, request.OrderID)
	if err != nil {
		return orderData, err
	}

	if orderData.UserID != request.UserID {
		return model.Order{}, errors.New("unauthorized")
	}

	return orderData, nil
}

func (r *restoUsecase) RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	userRegistered, err := r.userRepo.CheckRegistered(request.Username)
	if err != nil {
		return model.User{}, err
	}

	if userRegistered {
		return model.User{}, errors.New("user already registered")
	}

	userHash, err := r.userRepo.GenerateUserHash(request.Password)
	if err != nil {
		return model.User{}, err
	}

	userData, err := r.userRepo.RegisterUser(model.User{
		ID:       uuid.New().String(),
		Username: request.Username,
		Hash:     userHash,
	})
	if err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (r *restoUsecase) Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "Login")
	defer span.End()

	userData, err := r.userRepo.GetUserData(request.Username)
	if err != nil {
		return model.UserSession{}, err
	}

	verified, err := r.userRepo.VerifyLogin(request.Username, request.Password, userData)
	if err != nil {
		return model.UserSession{}, err
	}

	if !verified {
		return model.UserSession{}, errors.New("can't verify user login")
	}

	userSession, err := r.userRepo.CreateUserSession(userData.ID)
	if err != nil {
		return model.UserSession{}, err
	}

	return userSession, nil
}

func (r *restoUsecase) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "CheckSession")
	defer span.End()

	userID, err = r.userRepo.CheckSession(data)
	if err != nil {
		return "", err
	}

	return userID, nil
}
