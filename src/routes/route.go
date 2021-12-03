package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-ambassador/src/controllers"
	"go-ambassador/src/middlewares"
)

func Setup(app *fiber.App)  {
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login",controllers.Login)

""/admin/register""

	adminMiddleware := admin.Use(middlewares.IsAuthenticated)
	adminMiddleware.Get("user",controllers.User)
	adminMiddleware.Post("logout",controllers.Logout)
	adminMiddleware.Put("user/info",controllers.UpdateInfo)
	adminMiddleware.Put("user/password",controllers.UpdatePassword)
	adminMiddleware.Get("ambassador",controllers.Ambassadors)
	adminMiddleware.Post("products",controllers.CreateProducts)
	adminMiddleware.Get("products",controllers.Products)
	adminMiddleware.Put("products/:id",controllers.UpdateProduct)
	adminMiddleware.Get("products/:id",controllers.GetProduct)
	adminMiddleware.Delete("products/:id",controllers.DeleteProduct)
	adminMiddleware.Get("users/:id/links",controllers.Link)
	adminMiddleware.Get("orders",controllers.Orders)


	ambassador := api.Group("ambassador")

	ambassador.Post("register", controllers.Register)
	ambassador.Post("login",controllers.Login)
	ambassador.Get("produts/frontend",controllers.ProductsFrontEnd)
	ambassador.Get("produts/backend",controllers.ProductBackend)

	ambassadorAuth := ambassador.Use(middlewares.IsAuthenticated)
	ambassadorAuth.Get("user",controllers.User)
	ambassadorAuth.Post("logout",controllers.Logout)
	ambassadorAuth.Put("users/info",controllers.UpdateInfo)
	ambassadorAuth.Put("users/password",controllers.UpdatePassword)
	ambassadorAuth.Post("links",controllers.CreateLink)
	ambassadorAuth.Get("stats",controllers.Stats)
	ambassadorAuth.Get("rankings",controllers.Rankings)

}