package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3030", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	course := Course{
		Name:        "Curso de Go",
		Description: "Aulas do curso de go",
		Price:       1000,
	}
	json.NewEncoder(w).Encode(course)
}
