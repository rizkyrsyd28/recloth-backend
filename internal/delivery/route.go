package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/delivery/handler"
	"github.com/rizkyrsyd28/recloth-backend/internal/middleware"
	"github.com/rizkyrsyd28/recloth-backend/internal/repository"
	"github.com/rizkyrsyd28/recloth-backend/internal/usecase"
)

func MainRoutes(f *fiber.App) {
	r := repository.NewRepo()
	u := usecase.NewUsecase(r)

	f.Route("/auth", AuthRoute(u))
	f.Route("/api", APIRoute(u))
}

func AuthRoute(u usecase.Usecase) func(router fiber.Router) {
	return func(api fiber.Router) {
		api.Get("/user-info", middleware.JWTMiddleware(), handler.UserInfo(u))
		api.Post("/register", handler.Register(u))
		api.Post("/login", handler.Login(u))
		api.Post("/logout", handler.Logout(u))
	}
}

func APIRoute(u usecase.Usecase) func(router fiber.Router) {
	return func(api fiber.Router) {

		api.Get("/products/:page", handler.GetItems(u))
		api.Get("/product/:id", handler.GetItem(u))

		productPro := api.Group("/product", middleware.JWTMiddleware())
		productPro.Put("/:id", handler.UpdateItem(u))
		productPro.Delete("/:id", handler.DeleteItem(u))
		productPro.Post("/", handler.PostItem(u))

		transaction := api.Group("/transaction", middleware.JWTMiddleware())
		transaction.Get("/", handler.GetTransaction(u))
		transaction.Post("/checkout", handler.Checkout(u))

		cart := api.Group("/cart", middleware.JWTMiddleware())
		cart.Get("/", handler.GetCart(u))
		cart.Put("/", handler.UpdateCart(u))

	}
}
