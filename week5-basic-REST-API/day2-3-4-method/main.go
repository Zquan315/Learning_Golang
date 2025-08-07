package main

import ( 
	f "fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"

)
type User struct {
	ID string `json:"Id"`
	Name string `json:"Name"`
}

var users []User

func HomePage(w http.ResponseWriter, r *http.Request) {
	f.Fprint(w, "This is my Home Page")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Error in dencoding users", http.StatusInternalServerError)
	}

}
func GetUserWithID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/")
	for _, v := range users {
		if id == v.ID {
			w.Header().Set("Content-Type", "application/json")

			err := json.NewEncoder(w).Encode(v)
			if err != nil {
				http.Error(w, "Error in dencoding users", http.StatusInternalServerError)
			}
			break
		}
	}
}

func Init() {
	for i:=0; i < 3 ; i++ {
		s := strconv.Itoa(i)
		user := User{
			ID : s,
			Name: "Quan" + s,
		}
		users = append(users, user)
	}
}

func Handler() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/users", GetUsers)
	http.HandleFunc("/user/", GetUserWithID)
	f.Println("Server is listening on: http://localhost:3101")
	http.ListenAndServe(":3101", nil)
}
func main() {
	Init()
	Handler()
}