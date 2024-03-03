package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

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

func EditRecord(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal("Edited the item")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal("Deleted the item")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(jsonData)
}
