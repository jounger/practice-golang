package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

var lock = &sync.Mutex{}

type Wiki struct {
	Title string
	Content string
}

var singleInstance *Wiki

func getInstance() *Wiki {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &Wiki{}
			fmt.Println("Wiki has been created successful")
		} else {
			fmt.Println("Wiki has been created fail")
		}
	} else {
		fmt.Println("Wiki has been created fail")
	}
	return singleInstance
}

func (wi *Wiki) save() error {
	filename := wi.Title + ".txt"
	source := "../resources/" + filename
	return ioutil.WriteFile(source, []byte(wi.Content), 0600)
}

func (wi *Wiki) loadContent() (string, error){
	filename := wi.Title + ".txt"
	source := "../resources/" + filename
	content, err := ioutil.ReadFile(source)
	if err != nil {
		return "", err
	}
	return string(content), nil
}