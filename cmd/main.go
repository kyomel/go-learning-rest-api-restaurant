package main

import (
	"crypto/rand"
	"crypto/rsa"
	"rest-api-restaurant/internal/database"
	"rest-api-restaurant/internal/delivery/rest"
	"rest-api-restaurant/internal/logger"
	mRepo "rest-api-restaurant/internal/repository/menu"
	oRepo "rest-api-restaurant/internal/repository/order"
	uRepo "rest-api-restaurant/internal/repository/user"
	rUseCase "rest-api-restaurant/internal/usecase/resto"
	"time"

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
	logger.Init()
	e := echo.New()
	// localhost:14045/menu/food

	db := database.GetDB(dbAddress)
	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60*time.Second)
	if err != nil {
		panic(err)
	}

	restoUseCase := rUseCase.GetUseCase(menuRepo, orderRepo, userRepo)
	h := rest.NewHandler(restoUseCase)

	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}
