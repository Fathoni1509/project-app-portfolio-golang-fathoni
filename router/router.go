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
		// r.Get("/", handler.HandlerMenu.PersonalView)
		r.Get("/", handler.HandlerMenu.PortfolioView)
		// r.Get("/", handler.PersonalHandler.GetDataPersonal)
	})

	// edit page
	r.Route("/edit", func(r chi.Router) {
		// r.Get("/", handler.HandlerMenu.PersonalView)
		r.Get("/", handler.HandlerMenu.EditView)
		r.Post("/personal/update", handler.PersonalHandler.UpdatePersonal)
		
		r.Post("/activity/create", handler.ActivityHandler.CreateActivity)
		r.Post("/activity/update", handler.ActivityHandler.UpdateActivity)
		r.Post("/activity/delete", handler.ActivityHandler.DeleteActivity)

		r.Post("/work/create", handler.WorkHandler.CreateWork)
		r.Post("/work/update", handler.WorkHandler.UpdateWork)
		r.Post("/work/delete", handler.WorkHandler.DeleteWork)

		r.Post("/project/create", handler.ProjectHandler.CreateProject)
		r.Post("/project/update", handler.ProjectHandler.UpdateProject)
		r.Post("/project/delete", handler.ProjectHandler.DeleteProject)
	})

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	return r
}

// func Apiv1(handler handler.Handler) *chi.Mux {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	// read update personal
// 	r.Route("/personal", func(r chi.Router) {
// 		r.Get("/", handler.PersonalHandler.GetDataPersonal)
// 		r.Put("/", handler.PersonalHandler.UpdatePersonal)
// 	})

// 	return r
// }
