package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Listings struct {
	Numlistings   int       `json:"numlistings"`
	City          string    `json:"city"`
	Neighbourhood string    `json:"neighbourhood"`
	Area          string    `json:"area"`
	Listings      []Listing `json:"listings"`
}

type Listing struct {
	MLS_num     string `json:"mlsnum"`
	Price       string `json:"price"`
	Address     string `json:"address"`
	Bedrooms    string `json:"bedrooms"`
	Bathrooms   string `json:"bathrooms"`
	Parking     string `json:"parking"`
	Size        string `json:"size"`
	Timelisted  string `json:"timelisted"`
	Brokerage   string `json:"brokerage"`
	TimeScraped string `json:"timescraped"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("test.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["Listings"])
	fmt.Println(result["num_listings"])
	//birdJson := `[{ "num_listings": "36 ","City": "Toronto","Neighbourhood": "Downtown","Area": "Moss Park"},{"num_listings":"39","City":"Toronto","Neighbourhood":"Downtown","Area":"The Waterfront"}]`
	//var listing []Listing
	//json.Unmarshal([]byte(birdJson), &listing)
	//fmt.Printf("Listing %+v", listing)

}
