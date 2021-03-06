package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Vokabellists(w http.ResponseWriter, r *http.Request) {
	// return all available Vokabelset titles
	data := loadLists()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func Vokabelsets(w http.ResponseWriter, r *http.Request) {
	// return all data in one defined Vokabelset(set) as JSON
	vars := mux.Vars(r)
	setRaw := string(vars["list"])
	set, err := strconv.Atoi(setRaw)
	//fmt.Printf("Chose Set: %d\n", set)
	if err != nil {
		fmt.Fprintf(w, "Error: Bad List Number: %s", setRaw)
		return
	}
	//fmt.Printf("Checked for Vokabelset %d\n", set)
	data := loadVokabel(set)
	if data == nil {
		fmt.Fprint(w, "Dataset empty")
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
