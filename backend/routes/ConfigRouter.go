package routes

import (
	"github.com/TechMaster/golang/15GoMySQL/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigRouter(app *fiber.App) {

	productAPI := app.Group("/api/product")
	routeProduct(&productAPI)
}


func routeProduct(router *fiber.Router) {
	(*router).Get("/", controller.GetProducts) //Liệt kê

	(*router).Get("/:id", controller.GetProductById)// tìm theo id

	(*router).Get("/category/:name_category", controller.GetProductsByCategory) // tìm theo category
}
