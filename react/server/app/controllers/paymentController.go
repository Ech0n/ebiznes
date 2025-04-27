package controllers

import (
	"net/http"
	"project/database"
	"project/models"

	"github.com/labstack/echo/v4"
)

type CreateOrderRequest struct {
	Items []struct {
		ProductID uint `json:"productId"`
		Quantity  uint `json:"quantity"`
	} `json:"items"`
}

func ConfirmOrder(c echo.Context) error {
	var req CreateOrderRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	var totalPrice float64
	var orderItems []models.OrderItem

	for _, item := range req.Items {
		var product models.Product
		if err := database.DB.First(&product, item.ProductID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Product not found"})
		}

		orderItem := models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
		}
		orderItems = append(orderItems, orderItem)

		totalPrice += product.Price * float64(item.Quantity)
	}

	order := models.Order{
		TotalPrice: totalPrice,
		OrderItems: orderItems,
	}

	database.DB.Create(&order)

	return c.JSON(http.StatusCreated, order)
}