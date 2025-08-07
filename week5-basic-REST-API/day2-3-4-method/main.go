package main

import ( 
	f "fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"strings"

)
type User struct {
	ID string `json:"ID"`
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

func PostUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This is method Post.", http.StatusInternalServerError)
		return
	}

	var user User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		f.Fprint(w, "Error in decode json to User")
		return
	}

	users = append(users, user)
	f.Fprint(w, "Add successfully a user with ID:", user.ID, " and Name: ", user.Name)
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
	http.HandleFunc("/add-user", PostUser)
	http.HandleFunc("/user/", GetUserWithID)
	f.Println("Server is listening on: http://localhost:3101")
	http.ListenAndServe(":3101", nil)
}
func main() {
	Init()
	Handler()
}