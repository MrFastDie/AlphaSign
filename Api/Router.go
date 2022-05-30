package Api

import (
	"github.com/MrFastDie/AlphaSign/Api/Routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var router chi.Router

func GetRouter() chi.Router {
	if nil != router {
		return router
	}

	router = createRouter()
	return router
}

func createRouter() chi.Router {
	router = chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", Routes.Root)

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/reset-ram", Routes.ResetRam)
		r.Post("/write-text", Routes.WriteText)
		r.Post("/special-sign", Routes.SpecialSign)
		r.Post("/set-date", Routes.SetDate)
	})

	return router
}
