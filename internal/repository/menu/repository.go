package menu

import "rest-api-restaurant/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
