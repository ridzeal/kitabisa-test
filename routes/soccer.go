package routes

import (
	"github.com/go-chi/chi"

	"kitabisa-test/handler"
)

// SoccerRouter add routing to soccer domain
func SoccerRouter(r chi.Router) {
	r.Route("/team", func(r chi.Router) {
		r.Get("/", handler.GetTeams)
		r.Post("/", handler.AddTeam)

		r.Route("/{teamID}", func(r chi.Router) {
      r.Get("/", handler.GetTeamDetail)
      r.Delete("/", handler.RemoveTeam)
    })
	})

	r.Route("/player", func(r chi.Router) {
		r.Post("/", handler.AddPlayer)

		r.Route("/{playerID}", func(r chi.Router) {
			r.Delete("/", handler.RemovePlayer)
		})
	})
}