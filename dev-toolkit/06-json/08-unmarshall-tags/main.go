package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// autogenerated
// but names changed
type city struct {
	Pre string  `json:"precision"`
	Lat float64 `json:"Latitude"`
	Lon float64 `json:"Longitude"`
	Adr string  `json:"Address"`
	Cit string  `json:"City"`
	Sta string  `json:"State"`
	Zip string  `json:"Zip"`
	Con string  `json:"Country"`
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
	fmt.Println(data[1].Lat)

}
