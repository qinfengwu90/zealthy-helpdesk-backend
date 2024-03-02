package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/admins", func(r chi.Router) {
		r.Get("/create", createAdminHandler)
		r.Post("/login", loginAdminHandler)
		r.Post("/password", changeAdminPasswordHandler)
		r.Get("/all-tickets", getAllTicketsHandler)
	})

	// some function needs to be behind middleware
	// update ticket status
	// edit ticket
	// change password
	// logout

	router.Route("/users", func(r chi.Router) {
		r.Get("/all-tickets", getAllTicketsFromUserHandler)
		r.Post("/edit-ticket", editUserTicketHandler)
		r.Post("/create", createTicketHandler)
	})

	return router
}
