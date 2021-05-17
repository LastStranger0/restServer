package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restserver/internal/structs/companystruct"
	"strconv"
)

// GetAllCompanies ...
func GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	companies := database.GetAllCompanys()
	bytes, err := json.Marshal(companies)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid operation", 500)
		return
	}
	_, writeError := w.Write(bytes)
	if writeError != nil {
		log.Println(err)
		http.Error(w, "Invalid operation", 500)
		return
	}
}

// CreateCompany ...
func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var result companystruct.Company
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid operation", 500)
		return
	}
	database.AddCompany(result)
}

// DeleteCompany ...
func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	idInt, _ := strconv.Atoi(idString)

	successful := database.DeleteCompany(idInt)
	if !successful {
		log.Println("DeleteCompany err")
		http.Error(w, "Invalid operation", 500)
	}
}
