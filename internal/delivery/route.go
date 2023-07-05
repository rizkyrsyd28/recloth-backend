package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizkyrsyd28/recloth-backend/internal/delivery/handler"
	"github.com/rizkyrsyd28/recloth-backend/internal/repository"
	"github.com/rizkyrsyd28/recloth-backend/internal/usecase"
)

func MainRoutes(f *fiber.App) {
	r := repository.NewRepo()
	u := usecase.NewUsecase(r)

	f.Route("/auth", AuthRoute(u))
	f.Route("/api", ProtectedRoute(u))
}

func AuthRoute(u usecase.Usecase) func(router fiber.Router) {
	return func(api fiber.Router) {
		api.Post("/register", handler.Register(u))
		api.Post("/login", handler.Login(u))
		api.Post("/logout", handler.Logout(u))
	}
}

func ProtectedRoute(u usecase.Usecase) func(router fiber.Router) {
	return func(api fiber.Router) {
		api.Post("/register", handler.Register(u))
		api.Post("/login", handler.Login(u))
		api.Post("/logout", handler.Logout(u))
	}
}
