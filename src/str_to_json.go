package main

import (
	"encoding/json"
	"fmt"
)

// Laptop has neccessary feilds
type Laptop struct {
	ID          int64   `json:"id"`
	Model       string  `json:"model"`
	Color       string  `json:"color"`
	Weight      float32 `json:"weight"`
	HealthCheck bool    `json:"healthcheck"`
}

// laptopJson assigns JSON string
func main() {
	laptopJson := `{
      "id" : 1000001 ,
      "model": "Asus Rog",
      "color" : "Black",
      "weight": 1.90,
      "healthcheck":true}`

	var laptopObj Laptop

	/* JSON string to JSON object conversion
	   After conversion, laptop stores the result
	*/
	err := json.Unmarshal([]byte(laptopJson), &laptopObj)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("JSON string to JSON object conversion:")
	fmt.Printf("Laptop ID: %d , Model: %s, Color: %s, Weight: %g, HealthCheck: %t \n",
		laptopObj.ID, laptopObj.Model, laptopObj.Color, laptopObj.Weight, laptopObj.HealthCheck)

	/* JSON object to JSON string conversion
	   laptopString holds the  JsonString which obtained after marshalling JSON object
	*/
	laptopString, _ := json.Marshal(laptopObj)
	fmt.Println("JSON object to JSON string conversion")
	fmt.Println(string(laptopString))

}
