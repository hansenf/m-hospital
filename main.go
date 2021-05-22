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

//Function delete Item
func deleteItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	_deleteItemAtUid(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

//Function delete Item by UID
func _deleteItemAtUid(uid string) {
	for index, item := range inventory {
		if item.UID == uid {
			//Delete item from slice
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
}

//Function update Item
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)
	//Delete the item at UID
	_deleteItemAtUid(params["uid"])
	//Create it with new data
	inventory = append(inventory, item)

	json.NewEncoder(w).Encode(inventory)

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
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory", getInventory).Methods("GET")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")
	
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
