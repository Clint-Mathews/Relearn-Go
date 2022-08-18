package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON")
	encodeJSON()
	decodeJSON()
}

func encodeJSON() {
	courses := []course{
		{"React Js", 100, "C.com", "123", []string{"web", "js"}},
		{"Vue Js", 200, "Vue.com", "c123", []string{"google", "js"}},
		{"Angular Js", 300, "Angular.com", "j123", nil}}

	// Package the data as JSON

	finalJSON, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJSON)
}

func decodeJSON() {
	jsonData := []byte(`
		{
                "coursename": "React Js",
                "price": 100,
                "website": "C.com",
                "tags": [
                        "web",
                        "js"
                ]
        }
	`)
	var courseData course

	checkData := json.Valid(jsonData)
	if checkData {
		fmt.Println("JSON valid")
		err := json.Unmarshal(jsonData, &courseData)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", courseData)
	} else {
		fmt.Println("JSON in-valid")
	}

	// There are some cases where you just want to add data to key value

	var onlineData map[string]interface{}

	err := json.Unmarshal(jsonData, &onlineData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", onlineData)

	for k, v := range onlineData {
		fmt.Printf("Key %v and Value is %v and Type of value %T\n", k, v, v)
	}

}
