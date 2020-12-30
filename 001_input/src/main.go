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
	var filePath string
	// E:\viettel-training\001_input\resources\config.json
	var data []byte
	var err error
	for {
		fmt.Println("Hello, please input config path:")
		fmt.Scanln(&filePath);
		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("The system cannot find the file specified.")
		} else {
			break
		}
	}
	var init = Wrap{}
	var res Wrap

	if err := json.Unmarshal([]byte(data), &res); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(init)
	fmt.Println(res)
	for _, el := range res.Instances {
		fmt.Println(el.Name, el.Counts)
	}
}