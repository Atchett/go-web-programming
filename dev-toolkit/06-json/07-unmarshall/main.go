package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// you can choose to unmarshall
// only some of the JSON data
// data structure that only has
// fields for some of the data
type city struct {
	Latitude, Longitude float64
	City                string
}

type cities []city

func main() {

	var data cities
	rcvd := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	// unmarshalled gets put into the memory address of variable
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)

}
