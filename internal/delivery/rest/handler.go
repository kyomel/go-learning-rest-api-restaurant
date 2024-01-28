package rest

import "rest-api-restaurant/internal/usecase/resto"

type handler struct {
	restoUsecase resto.Usecase
}

func NewHandler(restoUseCase resto.Usecase) *handler {
	return &handler{
		restoUsecase: restoUseCase,
	}
}
