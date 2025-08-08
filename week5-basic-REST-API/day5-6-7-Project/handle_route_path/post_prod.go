package handle_route_path

import (
	f "fmt"
	"net/http"
	"encoding/json"
)

func Post_prod(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This is POST method", http.StatusMethodNotAllowed)
		Error.Println("This is POST method")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var p Product

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		f.Println("Error in decode to product")
		return
	}

	if CheckID(p.ID) {
		http.Error(w, "This ID is exists", http.StatusConflict)
		Error.Println("This ID is exists")
		return
	}

	products = append(products, p)
	f.Fprintf(w, "✅ Add a product successfully:\nID: %s\nName: %s\nPrice: %d", p.ID, p.Name, p.Price)
	Info.Println("✅ Add a product successfully")
}