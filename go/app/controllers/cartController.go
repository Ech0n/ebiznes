package controllers

import (
	"net/http"
	"project/database"
	"project/models"
	// "strconv"

	"github.com/labstack/echo/v4"
)

func AddToCart(c echo.Context) error {
	cart := new(models.Cart)
	if err := c.Bind(cart); err != nil {
		return err
	}
	database.DB.Create(&cart)
	return c.JSON(http.StatusCreated, cart)
}

func GetCart(c echo.Context) error {
	var carts []models.Cart
	database.DB.Preload("Product").Find(&carts)
	return c.JSON(http.StatusOK, carts)
}

func UpdateCart(c echo.Context) error {
	id := c.Param("id")
	var cart models.Cart
	database.DB.First(&cart, id)

	if err := c.Bind(&cart); err != nil {
		return err
	}
	database.DB.Save(&cart)
	return c.JSON(http.StatusOK, cart)
}

func DeleteFromCart(c echo.Context) error {
	id := c.Param("id")
	database.DB.Delete(&models.Cart{}, id)
	return c.NoContent(http.StatusNoContent)
}

