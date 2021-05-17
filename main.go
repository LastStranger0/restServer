package main

import (
	"flag"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"os"
	"restserver/internal/handler"
	"restserver/internal/sqlite"
)

func main() {
	_ = flag.Bool("dockerStart", false, "if u start from docker")
	flag.Parse()
	initDB()

	initLog("logging.log")

	mux := mux.NewRouter()
	mux.HandleFunc("/", handler.GetHtml)

	mux.HandleFunc("/companies", handler.GetAllCompanies)
	mux.HandleFunc("/company", handler.CreateCompany).Methods("POST")
	mux.HandleFunc("/company/{id:[0-9]+}", handler.DeleteCompany).Methods("POST")

	mux.HandleFunc("/wasd", handler.SignIn)
	mux.HandleFunc("/login", handler.HandleLogin)
	mux.HandleFunc("/callback", handler.HandleCallBack)

	mux.HandleFunc("/vk", handler.AuthVk)
	mux.HandleFunc("/me", handler.AuthVkCallback)

	mux.HandleFunc("/Hi", handler.SayHello)

	mux.HandleFunc("/products", handler.GetAllProducts).Methods("GET")
	mux.HandleFunc("/product", handler.CreateProduct).Methods("POST")
	mux.HandleFunc("/editProduct/{id:[0-9]+}", handler.EditProduct).Methods("GET")
	mux.HandleFunc("/editProduct/{id:[0-9]+}", handler.SendEditProduct).Methods("POST")
	mux.HandleFunc("/deleteProduct/{id:[0-9]+}", handler.DeleteProduct)

	http.Handle("/", mux)

	fileServer := http.FileServer(http.Dir("./"))
	mux.Handle("/main.js", http.StripPrefix("", fileServer))

	//log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	errServe := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	log.Fatal(errServe)
}

func initDB() {
	sqlite.PostgreSQL.CreateTableProducts()
	sqlite.PostgreSQL.CreateTableCompanys()
	sqlite.PostgreSQL.CreateTableHistory()
}
func initLog(logPath string) {
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	log.SetFlags(log.Lshortfile)
}
