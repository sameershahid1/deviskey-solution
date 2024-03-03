package main

import (
	"dev-solution/controller"
	"dev-solution/database"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {

	database.ConnectToDatabase()
	database.MigrateData()
	defer database.SqlDB.Close()

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
	router.Use(middleware.RequestID)
	router.Use(middleware.CleanPath)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.AllowContentType("application/json"))

	router.Get("/vehicle-part/list", controller.GetRecordList)
	router.Post("/vehicle-part", controller.PostVehiclePart)
	router.Patch("/vehicle-part/{vehicleId}", controller.EditRecord)
	router.Get("/vehicle-part/{vehicleId}", controller.DeleteRecord)
	router.Get("/generate-pdf", controller.GeneratePdf)

	router.NotFound(controller.HandleNotFound)
	router.MethodNotAllowed(controller.HandleMethodNotAllowed)

	http.ListenAndServe(":8080", router)
}
