package handle_route_path

import (
	f "fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func PutID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "This is PUT method", http.StatusMethodNotAllowed)
		Error.Println("This is PUT method")
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	for i, v := range products {
		if id == v.ID {
			w.Header().Set("Content-Type", "application/json")

			err := json.NewDecoder(r.Body).Decode(&products[i])

			if err != nil {
				f.Println("Error in Decode products")
				return
			}
			f.Fprintf(w, "✅ Update a product successfully: ID: %s", id)
			Info.Println("✅ Update a product successfully")
			return
		}
	}

	http.Error(w, "ID is not exists", http.StatusNotFound)	
	Error.Println("ID is not exists")
}