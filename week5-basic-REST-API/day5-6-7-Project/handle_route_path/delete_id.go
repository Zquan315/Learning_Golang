package handle_route_path

import (
	f "fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func DeleteID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		Error.Println("This is DELETE method")
		http.Error(w, "This is DELETE method", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	for i, v := range products {
		if id == v.ID {
			w.Header().Set("Content-Type", "application/json")

			products = append(products[:i], products[i+1:]...)
			f.Fprintf(w, "✅ Delete a product successfully: ID: %s", id)
			Info.Println("✅ Delete a product successfully with ID:", id)
			return
		}
	}

	http.Error(w, "ID is not exists", http.StatusNotFound)
	Error.Println("ID is not exists")	
}