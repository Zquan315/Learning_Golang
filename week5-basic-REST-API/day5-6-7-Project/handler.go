package main

import (
	"github.com/gorilla/mux"
	"net/http"
	f "fmt"
	h "github.com/zquan315/Learning_Golang/week5-basic-REST-API/day5-6-7-Project/handle_route_path"
)

func Handler() {
	r := mux.NewRouter()

	f.Println("✅ Server is listening on: http://localhost:3101")
	h.Info.Println("Server started on port 3101")
	r.HandleFunc("/", h.Home).Methods("GET")
	r.HandleFunc("/products", h.GetAll).Methods("GET")
	r.HandleFunc("/products/{id}", h.GetID).Methods("GET")
	r.HandleFunc("/products/add", h.Post_prod).Methods("POST")
	r.HandleFunc("/products/update/{id}", h.PutID).Methods("PUT")
	r.HandleFunc("/products/del/{id}", h.DeleteID).Methods("DELETE")

	http.ListenAndServe(":3101", r)
	
}