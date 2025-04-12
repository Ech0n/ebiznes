package controllers

import (
	"net/http"
	"project/database"
	"project/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	println("Product fields:", product.Name, product.Price, ", categoryId:",product.CategoryID)
	database.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func GetProducts(c echo.Context) error {
	var products []models.Product
	query := database.DB.Preload("Category")

	if min := c.QueryParam("min"); min != "" {
		minVal, _ := strconv.ParseFloat(min, 64)

		query = query.Scopes(models.PriceGreaterThan(minVal))

	}
	if max := c.QueryParam("max"); max != "" {
		maxVal, _ := strconv.ParseFloat(max, 64)
		query = query.Scopes(models.PriceSmallerThan( maxVal))
	}
	if cat := c.QueryParam("category"); cat != "" {
		query = query.Scopes(models.FilterByCategory(cat))
	}

	query.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if result := database.DB.Preload("Category").First(&product, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	database.DB.First(&product, id)

	if err := c.Bind(&product); err != nil {
		return err
	}
	database.DB.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	database.DB.Delete(&models.Product{}, id)
	return c.NoContent(http.StatusNoContent)
}
