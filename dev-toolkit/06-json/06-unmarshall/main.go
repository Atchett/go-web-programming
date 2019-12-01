package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// from JSON to Go site
// type img struct {
// 	Width     int    `json:"Width"`
// 	Height    int    `json:"Height"`
// 	Title     string `json:"Title"`
// 	Thumbnail struct {
// 		URL    string `json:"Url"`
// 		Height int    `json:"Height"`
// 		Width  int    `json:"Width"`
// 	} `json:"Thumbnail"`
// 	Animated bool  `json:"Animated"`
// 	IDs      []int `json:"IDs"`
// }

// manual

type thumbnail struct {
	URL           string
	Height, Width int
}

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

func main() {
	var data img
	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	// sending the value back to the memory address of data
	// the value pointed to
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}

	fmt.Println(data)

	for i, v := range data.IDs {
		fmt.Println(i, v)
	}

	fmt.Println(data.Thumbnail.URL)

}
