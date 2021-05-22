package main

import (
	"fmt"
	//"m-hospital/calculation"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Item struct {
	UID string `json:"UID"`
	Name string`json:"Name"`
	Desc string `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

/*
func Multiply(num1 int, num2 int) int {
	return mul(num1, num2)
}

func mul(num1 int, num2 int) int {
	return num1 * num2
}
*/

//Batasan endpoint
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint called: homePage()")
}

//Function get Inventory
func getInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println("Function Called: getInventory")

	json.NewEncoder(w).Encode(inventory)
}

//Function menambahkan Item baru
func createItem (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var item Item 
	
	_ = json.NewDecoder(r.Body).Decode(&item)

	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(item)
}

//Memanggil page http port :8000
func handlerRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory",getInventory).Methods("GET")
	router.HandleFunc("/inventory",createItem).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//Halaman Utama
func main() {
	inventory = append(inventory, Item{
		UID: "0",
		Name: "Bodrex",
		Desc: "A fine medicine of headache, flu, and cough.",
		Price: 3000,
	})

	inventory = append(inventory, Item{
		UID: "1",
		Name: "Daktarin",
		Desc: "Anti fungiderm, clean up your skin.",
		Price: 7500,
	})

	handlerRequest()
}
