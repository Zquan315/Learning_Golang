package handle_route_path

import (
	f "fmt"
	"net/http"
	"encoding/json"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "This is GET method", http.StatusMethodNotAllowed)
		Error.Println("This is GET method")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		f.Println("Error in encode products")
		return
	}
}