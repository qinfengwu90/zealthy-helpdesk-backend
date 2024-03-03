package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
	"zealthy-helpdesk-backend/viper"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	tokenAuth := jwtauth.New(jwt.SigningMethodHS256.Name, []byte(viper.ViperReadEnvVar("JWT_SECRET")), nil)

	router.Route("/admins", func(r chi.Router) {
		r.Get("/login", loginAdminHandler)

		r.Group(func(r chi.Router) {
			//r.Use(adminAuthMiddleware)
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))

			r.Get("/all-tickets", getAllTicketsHandler)
			r.Post("/register-admin", registerAdminHandler)
			r.Post("/change-password", changeAdminPasswordHandler)
			r.Post("/update-ticket-status", updateTicketStatusHandler)
		})
	})

	router.Route("/users", func(r chi.Router) {
		r.Get("/all-tickets", getAllTicketsForUserHandler)
		r.Post("/create-ticket", createTicketHandler)
		r.Post("/edit-ticket", editUserTicketHandler)
	})

	return router
}
