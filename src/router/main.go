package router

import (
	"gotest/src/controllers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

// Router main rules of routes
func Router() http.Handler {
	r := chi.NewRouter()

	//CORS setup "http://localhost:4200",
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "DeviceID"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	r.Get("/test", controllers.Init())
	return r
}
