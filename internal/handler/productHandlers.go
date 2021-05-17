package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"restserver/internal/dbInterface"
	"restserver/internal/sqlite"
	"restserver/internal/structs/productstruct"
	"restserver/internal/tokens"
	"strconv"
)

var database dbInterface.Database

func init() {
	database = sqlite.PostgreSQL
}

// GetAllProducts ...
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := database.GetAllProducts()
	bytes, err := json.Marshal(products)
	if err != nil {
		log.Println(err)
		return
	}
	_, writeError := w.Write(bytes)
	if writeError != nil {
		log.Println(err)
		return
	}
}

// CreateProduct ...
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var result productstruct.Product
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid operation", 500)
		return
	}

	database.AddProduct(result)
}

// EditProduct ...
func EditProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	idInt, _ := strconv.Atoi(idString)
	product, successful := database.GetProduct(idInt)
	if !successful {
		log.Println("editProduct: Not found")
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	bytes, err := json.Marshal(product)
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

// SendEditProduct ...
func SendEditProduct(w http.ResponseWriter, r *http.Request) {
	var result productstruct.Product
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid operation", 500)
		return
	}
	database.UpdateProduct(result)
}

// DeleteProduct ...
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	cookie := r.Cookies()

	if len(cookie) < 1 {
		log.Println("err")
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userID, err := tokens.ParseTokens(cookie[0].Value)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	idString := vars["id"]
	idInt, _ := strconv.Atoi(idString)

	successful := database.DeleteProduct(idInt)
	if !successful {
		log.Println("error")
		http.Error(w, "Invalid operation", 500)
		return
	}

	err = sqlite.PostgreSQL.SaveChangeHistory(userID, idString)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetHtml(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
