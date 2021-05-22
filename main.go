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

//Memanggil page http port :8000
func handlerRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")

	router.HandleFunc("/inventory",getInventory).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

//Halaman Utama
func main() {
	inventory = append(inventory, Item{
		UID: "0",
		Name: "Cheese",
		Desc: "A fine block of cheese",
		Price: 4.99,
	})

	handlerRequest()
}
