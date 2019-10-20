package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"encoding/json"
	"net/http"
)

// below is the struct that represents one section of the data
type usefulChunk struct {
		Time string `json:"time"`
		Latitude string `json:"latitude"`
		Latitude_dir string `json:"latitude_dir"`
		Longitude string `json:"longitude"`
		Longitude_dir string `json:"longitude_dir"`
		Nsats string `json:"n_sats"`
		Altitude string `json:"altitude"`
	}

func takeItAllLOL() []usefulChunk {
	// reads all of the data
	f, _ := os.Open("smallcsv.txt")
	reader := csv.NewReader(f)
	var finalSlice []usefulChunk
	line, err := reader.Read()
	fmt.Println(line)
	remainder, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fmt.Println(remainder)
	for _, line := range remainder {
		fmt.Println(line)
		fmt.Println(line[1])
		pieceOfData := usefulChunk{Time: line[2], Latitude: line[6], Latitude_dir: line[7], Longitude: line[8], Longitude_dir: line[9], Nsats: line[11], Altitude: line[13]}
		fmt.Println(pieceOfData)
		finalSlice = append(finalSlice, pieceOfData)
	}
	fmt.Println(err)
	fmt.Println("Here's another thing")
	fmt.Println(finalSlice)
	return finalSlice
}
/*

func pickAPeck (index uint64) usefulChunk {
	// allows a user to retrieve the data at a specific index
	f, _ := os.Open("smallcsv.txt")
	reader := csv.NewReader(f)
	line, err := reader.Read()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i <= index; i++ {
		line, err := reader.Read()
		if err != nil {
			fmt.Println(err)
		}
	defer f.Close()
	pieceOfData := usefulChunk{line[2], line[6], line[7], line[8], line[9], line[11], line[13]}
	fmt.Println(pieceOfData)
	return pieceOfData
	}
} */

func handlePigs(w http.ResponseWriter, r *http.Request) {
	// allows the user to retrieve everything
	thingOfStuff := takeItAllLOL()
	if err := json.NewEncoder(w).Encode(thingOfStuff); err != nil {
		fmt.Println(err)
	}
}
/*
func handleSnobs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pickAPeck())
}*/
	
func main() {
	fmt.Println("Here's a thing")
	http.HandleFunc("/iwantitall", handlePigs)
	http.HandleFunc("/", handlePigs)
	http.ListenAndServe(":8000", nil)
}