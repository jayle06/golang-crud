package routes

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/train/controller"
)

func ConfigBookRouter(router *fiber.Router) {
	(*router).Get("/books", controller.GetAllBooks)
	(*router).Post("/books", controller.CreateNewBook)
	(*router).Get("/books/:id", controller.GetBookById)
	(*router).Delete("/books/:id", controller.DeleteBookById)
	(*router).Patch("/books", controller.UpdateBookById)
	(*router).Put("/books/:id", controller.UpsertBook)
}

func ConfigReviewRouter(router *fiber.Router) {
	(*router).Get("/:id/reviews", controller.GetReviewByBookId)
	(*router).Get("/", controller.GetAllReview)
	(*router).Patch("/:id", controller.UpdateReviewById)
	(*router).Post("/", controller.CreateNewReView)
}
