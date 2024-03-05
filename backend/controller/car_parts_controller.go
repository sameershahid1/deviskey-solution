package controller

import (
	"dev-solution/database"
	"dev-solution/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

type vehicleRequestBody struct {
	Name        string
	Description string
	Price       float64
}

type response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type paginate struct {
	PageNo     int                  `json:"pageNo"`
	PerPage    int                  `json:"perPage"`
	TotalCount *int64               `json:"totalCount"`
	RecordList *[]model.VehiclePart `json:"recordList"`
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
	var pagination paginate
	// var dataList []model.VehiclePart
	var count int64
	var offset int

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&pagination)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	offset = (pagination.PageNo - 1) * pagination.PerPage
	database.GormDB.Model(&model.VehiclePart{}).Count(&count)
	database.GormDB.Limit(pagination.PerPage).Offset(offset).Find(&pagination.RecordList)
	pagination.TotalCount = &count

	jsonData, err := json.Marshal(pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func PostVehiclePart(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var data vehicleRequestBody
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	vehiclePart := model.VehiclePart{Name: data.Name, Description: data.Description, Price: data.Price}
	result := database.GormDB.Create(&vehiclePart)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(response{Status: true, Message: "Successfully added Vehicle record"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func EditRecord(w http.ResponseWriter, r *http.Request) {
	var idTemp uint64
	var parse string = chi.URLParam(r, "vehicleId")
	idTemp, err := strconv.ParseUint(parse, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	id := uint(idTemp)

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var data vehicleRequestBody
	err = decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	result := database.GormDB.Model(&model.VehiclePart{}).Where("id", id).Updates(data)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(response{Status: true, Message: "Successfully Edited Vehicle Record"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	var idTemp uint64
	var parse string = chi.URLParam(r, "vehicleId")
	idTemp, err := strconv.ParseUint(parse, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	id := uint(idTemp)

	database.GormDB.Delete(&model.VehiclePart{}, id)
	jsonData, err := json.Marshal(response{Status: true, Message: "Successfully Deleted Vehicle record"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	jsonData, err := json.Marshal(response{Status: false, Message: "Route does not exist"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func HandleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	jsonData, err := json.Marshal(response{Status: false, Message: "Invalid Method"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
