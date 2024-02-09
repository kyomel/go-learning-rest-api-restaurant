package database

import (
	"rest-api-restaurant/internal/model"
	"rest-api-restaurant/internal/model/constant"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{})
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmi",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     2500,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es_jeruk",
			Price:     2500,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}

}
