package main

import (
	"dev-solution/controller"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {

	router := chi.NewRouter()
	corsOption := cors.Options{
		AllowedOrigins:   []string{"https://localhost:3000", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}

	router.Use(cors.Handler(corsOption))
	router.Use(middleware.CleanPath)
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentType("application/json", "text/xml"))

	//We will do filtering on this request
	router.Get("/record/list", controller.GetRecordList)
	router.Patch("/record/{id}", controller.EditRecord)
	router.Delete("/record/{id}", controller.DeleteRecord)

	//It will used to generate pdf file and send it the client side
	router.Get("/generate-pdf", controller.GeneratePdf)

	http.ListenAndServe(":8080", router)
}
