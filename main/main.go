package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User Model
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Main Page Handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello! Welcome to Go Server")
}

// Json Data Handler
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post Method supported !", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Unvalid JSON", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("User received: %s <%s>", user.Name, user.Email)
	fmt.Fprintln(w, response)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/user", userHandler)

	fmt.Println("Server running on port 8080 ..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
