package handle_route_path

import (
	"os"
	"log"
)

type Product struct {
	ID string `json:"ID"`
	Name string `json:"Name"`
	Price int32 `json:"Price"`
}

var Info *log.Logger
var Error *log.Logger

var products []Product

func Init() {

	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Cannot open log file:", err)
	}

	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	
	p := Product{
		ID : "1",
		Name: "Candy",
		Price: 15000,
	}

	products = append(products, p)


}