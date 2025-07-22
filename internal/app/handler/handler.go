package handler

import (
	service "test-assignment/internal/app/services"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type SubscriptionHandler struct {
	Svc service.SubscriptionService
}

func (sub *SubscriptionHandler) RegisterRoutes(r chi.Router) {
	r.Route("/subscriptions", func(r chi.Router) {
		r.Post("/", sub.Create)
		r.Get("/", sub.List)
		r.Get("/{id}", sub.Get)
		r.Put("/{id}", sub.Update)
		r.Delete("/{id}", sub.Delete)
	})
	r.Get("/total", sub.TotalCost)
}

func (sub *SubscriptionHandler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5000/swagger/doc.json"),
	))
	sub.RegisterRoutes(r)
	return r
}
