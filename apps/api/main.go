package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type addressBook struct {
	Firstname string
	Lastname  string
	Code      int
	Phone     string
}

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Chaiyarin",
		Lastname:  "Niamsuwan",
		Code:      1993,
		Phone:     "0870940955",
	}
	json.NewEncoder(w).Encode(addBook) // (1)
}

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getAddress", getAddressBookAll) // เพิ่มบรรทัดนี้
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
