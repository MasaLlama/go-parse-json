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
	jsonFile, err := os.Open("example.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened example.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["Listings"])
	fmt.Println(result["num_listings"])
	//`[{ "num_listings": "36 ","City": "Toronto","Neighbourhood": "Downtown","Area": "Moss Park"},{"num_listings":"39","City":"Toronto","Neighbourhood":"Downtown","Area":"The Waterfront"}]`
	//
	//
	//Example output
	//[map[Address:1007 - 219 Dundas  St E Bathrooms:2BA Bedrooms:3BD Brokerage:LIVING REALTY INC., BROKERAGE MLS_num:C5467882 Parking:0 Parking Price:3,200 Size:800-899 sqft timeListed:11 hours timeScraped:01/10/2022-00:13:38]
	//map[Address:2208 - 219 Dundas St E Bathrooms:1BA Bedrooms:1+1BD Brokerage:LIVING REALTY INC., BROKERAGE MLS_num:C5467667 Parking:0 Parking Price:2,400 Size:500-599 sqft timeListed:1 day timeScraped:01/10/2022-00:13:38]]
	//

}
