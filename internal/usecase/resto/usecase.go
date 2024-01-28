package resto

import "rest-api-restaurant/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
