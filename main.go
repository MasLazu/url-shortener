package main

import (
	"log"
	"net/http"
	"url-shortener/config"
	"url-shortener/handler"
	"url-shortener/repository"
	"url-shortener/views"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func main() {
	//load .env
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	//bootstrap
	views := views.NewView()
	database := config.NewDatabase()
	urlRepository := repository.NewUrlRepository(database)
	urlHandler := handler.NewUrlHandler(urlRepository, views)

	//close database conn
	defer database.Close()

	//router
	r := mux.NewRouter()
	r.HandleFunc("/url", urlHandler.Add).Methods("POST")
	r.HandleFunc("/{url}", urlHandler.Visit).Methods("GET")

	log.Println("server runing on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
