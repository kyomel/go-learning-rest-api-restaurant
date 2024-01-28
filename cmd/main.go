package main

import (
	"rest-api-restaurant/internal/database"
	"rest-api-restaurant/internal/delivery/rest"
	mRepo "rest-api-restaurant/internal/repository/menu"
	rUseCase "rest-api-restaurant/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost user=postgres password=password dbname=go_resto_app port=5432 sslmode=disable"
)

// Not Use Because Don't Need But Example
// func getFoodMenu(c echo.Context) error {
// 	// return c.JSON(http.StatusOK, foodMenu)
// 	// Optional using map string

// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic(err)
// 	}

// 	var menuData []MenuItem

// 	db.Where(MenuItem{Type: MenuTypeFood}).Find(&menuData)

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": menuData,
// 	})
// }

// func getDrinkMenu(c echo.Context) error {
// 	// return c.JSON(http.StatusOK, foodMenu)
// 	// Optional using map string

// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic(err)
// 	}

// 	var menuData []MenuItem

// 	db.Where(MenuItem{Type: MenuTypeDrink}).Find(&menuData)

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": menuData,
// 	})
// }

func main() {
	e := echo.New()
	// localhost:14045/menu/food

	db := database.GetDB(dbAddress)

	menuRepo := mRepo.GetRepository(db)
	restoUseCase := rUseCase.GetUseCase(menuRepo)
	h := rest.NewHandler(restoUseCase)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}
