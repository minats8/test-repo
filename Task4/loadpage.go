package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	structure "task4/dataStructure"
	"task4/models"
)

const URL string = "https://jsonplaceholder.typicode.com/posts"

func LoadPageData(postNO string) (models.Data, error) {
	newURL := URL + "/" + postNO
	var data models.Data
	resp, err := http.Get(newURL)
	if err != nil {
		fmt.Println("HTTP Error: ", err.Error())
		return data, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error: ", err.Error())
		return data, err
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Println("Parsing Error: ", err.Error())
		return data, err
	}
	if data.Id == 0 {
		fmt.Println("Unable to load this Request")
		os.Exit(1)
	}
	return data, nil

}

func Fetch(op structure.Operations) error {
	return op.GetData()
}

func WrtieOnStorage(op structure.Operations, data models.Data) error {
	return op.PutData(data)
}
