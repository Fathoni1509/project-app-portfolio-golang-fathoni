package router

import (
	"net/http"
	"project-app-portfolio-golang-fathoni/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(handler handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// main portfolio page
	r.Route("/", func(r chi.Router) {
		r.Get("/", handler.HandlerMenu.PortfolioView)
	})

	// edit page
	r.Route("/edit", func(r chi.Router) {
		r.Get("/", handler.HandlerMenu.EditView)

		r.Post("/personal/update", handler.PersonalHandler.UpdatePersonal)

		r.Route("/activity", func(r chi.Router) {
			r.Post("/create", handler.ActivityHandler.CreateActivity)
			r.Post("/update", handler.ActivityHandler.UpdateActivity)
			r.Post("/delete", handler.ActivityHandler.DeleteActivity)
		})
		
		r.Route("/work", func(r chi.Router) {
			r.Post("/create", handler.WorkHandler.CreateWork)
			r.Post("/update", handler.WorkHandler.UpdateWork)
			r.Post("/delete", handler.WorkHandler.DeleteWork)
		})

		r.Route("/project", func(r chi.Router) {
			r.Post("/create", handler.ProjectHandler.CreateProject)
			r.Post("/update", handler.ProjectHandler.UpdateProject)
			r.Post("/delete", handler.ProjectHandler.DeleteProject)
		})

		r.Route("/contact", func(r chi.Router) {
			r.Post("/create", handler.ContactHandler.CreateContact)
			r.Post("/update", handler.ContactHandler.UpdateContact)
			r.Post("/delete", handler.ContactHandler.DeleteContact)
		})

	})

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	return r
}