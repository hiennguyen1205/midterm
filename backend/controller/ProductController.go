package controller

import (
	
	"strconv"
	// "strings"

	"github.com/TechMaster/golang/15GoMySQL/model"
	"github.com/TechMaster/golang/15GoMySQL/repo"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []model.Product

	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	repo.Db.Preload("Category").
		Find(&products)
	var results []model.Product
	if offset+limit > len(products) {
		results = products[offset:]
	} else {
		results = products[offset : offset+limit]
	}

	answer := make(map[string]interface{})
	answer["listProduct"] = results
	answer["totalProducts"] = len(products)
	if len(products)%limit == 0 {
		answer["totalPages"] = len(products) / limit
	} else {
		answer["totalPages"] = len(products)/limit + 1
	}
	return c.JSON(answer)
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	product := repo.FindProductById(id)
	// fmt.Println(product)

	return c.JSON(product)
}

func GetProductsByCategory(c *fiber.Ctx) error {
	category := c.Params("name_category")
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	repo.Db.Preload("Category")
	results := repo.FindProductByCategory(category)
	var answer []model.Product
	if offset+limit > len(results) {
		answer = results[offset:]
	} else {
		answer = results[offset : offset+limit]
	}
	response := make(map[string]interface{})
	response["listProduct"] = answer
	response["totalProducts"] = len(results)
	
	if len(results)%limit == 0 {
		response["totalPages"] = len(results) / limit
	} else {
		response["totalPages"] = len(results)/limit + 1
	}
	return c.JSON(response)
}


