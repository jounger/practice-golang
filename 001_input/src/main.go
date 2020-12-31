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
	// ../resources/config.json
	// ../resources/config_2.json
	var init = Wrap{}
	for i := 0; i < 5; i++ {
		var res Wrap
		data, _ := input()
		if err := json.Unmarshal([]byte(data), &res); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		compare(res, &init)
	}
}

func input() ([]byte, error) {
	var filePath string
	for {
		fmt.Println("Hello, please input config path:")
		fmt.Scanln(&filePath)
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("The system cannot find the file specified.")
		} else {
			return data, err
		}
	}
}

func compare(res Wrap, init *Wrap) {
	for _, el := range res.Instances {
		// fmt.Println(el.Name, el.Counts)
		if i, found := find(init.Instances, el); found {
			if el.Counts > init.Instances[i].Counts {
				// add more
				fmt.Println(el.Name, "provision", el.Counts - init.Instances[i].Counts)
			} else if el.Counts < init.Instances[i].Counts {
				// remove
				fmt.Println(el.Name, "delete", init.Instances[i].Counts - el.Counts)
			} else {
				fmt.Println(el.Name, "N/A", init.Instances[i].Counts)
			}
			(*init).Instances[i].Counts = el.Counts
		} else {
			// add to slice
			init.Instances = append(init.Instances, el)
			fmt.Println(el.Name, "provision", el.Counts)
		}
	}
}

func find(arr []Instance, e Instance) (int, bool){
	for i, el := range arr {
		if el.Name == e.Name {
			return i, true
		}
	}
	return -1, false
}