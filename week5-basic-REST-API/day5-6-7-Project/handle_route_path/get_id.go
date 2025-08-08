package handle_route_path

import (
	f "fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "This is GET method", http.StatusMethodNotAllowed)
		Error.Println("This is GET method")
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	for _, v := range products {
		if id == v.ID {
			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(v)

			if err != nil {
				f.Println("Error in encode products")
				return
			}
			return
		}
	}

	http.Error(w, "ID is not exists", http.StatusNotFound)	
	Error.Println("ID is not exists")
}