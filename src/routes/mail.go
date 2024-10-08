package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/eirene/src/endpoint"
	"github.com/rulanugrh/eirene/src/internal/middleware"
)

func MailRoutes(f *fiber.App, endpoint endpoint.MailEndpoint) {
	mail := f.Group("/api/v1/mail")
	mail.Post("/sent", endpoint.Sent)
	mail.Use(middleware.JWTVerify())
	mail.Get("/inbox", endpoint.Inbox)
	mail.Get("/starred", endpoint.Starred)
	mail.Get("/archive", endpoint.Archive)
	mail.Delete("/delete/:id", endpoint.Delete)
	mail.Put("/update/:id", endpoint.Update)
}
