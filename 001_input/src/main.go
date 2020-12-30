package main

import (
	"fmt";
	"io/ioutil";
	"encoding/json";
	"log"
)

type Instance struct {
	Name string `json:"name"`
	VCPU int `json:"vCPU"`
	VRam int `json:"vRam"`
	Counts int `json:"counts"`
}

type Wrap struct {
	Instances []Instance `json:"instances"`
}

func main() {
	fmt.Println("Hello, playground")

	var filePath string
	fmt.Scanln(&filePath);

	// E:\viettel-training\001_input\resources\config.json

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	var res Wrap

	if err := json.Unmarshal([]byte(data), &res); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(res)
	fmt.Println(res.Instances[0].Counts)
}