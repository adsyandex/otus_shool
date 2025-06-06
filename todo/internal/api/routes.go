package api

import (
	"github.com/adsyandex/otus_shool/todo/internal/logger"
	"github.com/adsyandex/otus_shool/todo/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(service *service.TaskService, log *logger.Logger) *chi.Mux {
	router := chi.NewRouter()
	handler := NewHandler(service, log)

	// Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Routes
	router.Route("/tasks", func(r chi.Router) {
		r.Post("/", handler.CreateTask)
		r.Get("/", handler.ListTasks)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handler.GetTask)
			r.Put("/", handler.UpdateTask)
			r.Delete("/", handler.DeleteTask)
		})
	})

	return router
}
