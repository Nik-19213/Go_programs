package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	//Unmarshalling JSON
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)

	//Unmarshalling Unstructured Data
	str := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	var reading2 map[string]interface{}
	err = json.Unmarshal([]byte(str), &reading2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading2)

	//Handling Unstructured Lists in Go
	str2 := `[{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}]`

	var reading3 []map[string]interface{}
	err = json.Unmarshal([]byte(str2), &reading3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading3)

}
