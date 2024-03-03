package controller

import (
	"dev-solution/database"
	"dev-solution/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

type postRequestBody struct {
	Name        string
	Description string
	Price       float64
}

type postRequestResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func GeneratePdf(w http.ResponseWriter, r *http.Request) {
	// number := chi.URLParam(r, "number")

	pdfGenerator := pdf.NewMaroto(consts.Portrait, consts.A4)
	pdfGenerator.SetPageMargins(20, 10, 20)
	buildHeading(pdfGenerator)
	buildFooter(pdfGenerator)
	buildFruitList(pdfGenerator)

	err := pdfGenerator.OutputFileAndClose("./assets/example.pdf")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	pdf, err := os.Open("./assets/example.pdf")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer pdf.Close()

	pdfInfo, err := pdf.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=record.pdf")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", pdfInfo.Size()))

	io.Copy(w, pdf)
}

func GetRecordList(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal("Get items")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func PostVehiclePart(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var data postRequestBody
	err := decoder.Decode(&data)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	vehiclePart := model.VehiclePart{Name: data.Name, Description: data.Description, Price: data.Price}
	result := database.GormDB.Create(&vehiclePart)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	}

	jsonData, err := json.Marshal(postRequestResponse{Status: true, Message: "Successfully added Vehicle record"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func EditRecord(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal("Edited the item")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "vehicleId")
	fmt.Println(id)
	// "Deleted the item"
	jsonData, err := json.Marshal(fmt.Sprintln("ID: ", id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("route does not exist"))
}

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	w.Write([]byte("method is not valid"))
}
