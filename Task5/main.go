package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

func IsOpened(workerID int, fanIn <-chan string, timeOut int, result chan<- Validation) {
	for j := range fanIn {
		timeout := time.Duration(timeOut) * time.Second
		conn, err := net.DialTimeout("tcp", j, timeout)
		var flag bool = false
		if err != nil {
			flag = false
		}
		if conn != nil {
			conn.Close()
			flag = true
		}
		fmt.Println("worker", workerID, " > ", j, " -> ", strconv.FormatBool(flag))
		result <- Validation{j, flag}
	}
}

func WriteToFile(workerID int, fanOut <-chan Validation, output chan<- bool) {
	var Filepath string
	for j := range fanOut {
		if j.Status {
			Filepath = "valid.yml"
			file, err := os.OpenFile(Filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("File Error: ", err.Error())
				os.Exit(1)
			}
			data, err := yaml.Marshal(j.Sites)
			if err != nil {
				fmt.Println("Marshal Error: ", err.Error())
				os.Exit(1)
			}
			_, err = file.Write(data)
			if err != nil {
				fmt.Println("File Write Error", err.Error())
				os.Exit(1)
			}
		} else {
			Filepath = "invalid.yml"
			file, err := os.OpenFile(Filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("File Error: ", err.Error())
				os.Exit(1)
			}
			data, err := yaml.Marshal(j.Sites)
			if err != nil {
				fmt.Println("Marshal Error: ", err.Error())
				os.Exit(1)
			}
			_, err = file.Write(data)
			if err != nil {
				fmt.Println("File Write Error", err.Error())
				os.Exit(1)
			}
		}
		fmt.Println("worker#", workerID, " ", j.Sites, " -> ", j.Status)
		output <- j.Status
	}
}

func main() {
	yamlFile, err := ioutil.ReadFile("scan.yml")
	if err != nil {
		fmt.Println("File Read Error: ", err.Error())
		os.Exit(1)
	}
	var config Config
	var address Address
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Println("Unmarshal Error: ", err.Error())
		os.Exit(1)
	}
	if config.Workers == 0 {
		//sysctl -n hw.ncpu
		config.Workers = 2 * runtime.NumCPU()
	}
	yamlFile, err = ioutil.ReadFile(config.Filepath)
	if err != nil {
		fmt.Println("File Read Error: ", err.Error())
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, &address)
	if err != nil {
		fmt.Println("Unmarshal Error: ", err.Error())
		os.Exit(1)
	}
	fanIn := make(chan string, len(address.Sites))
	result := make(chan Validation, len(address.Sites))
	for w := 1; w <= config.Workers; w++ {
		go IsOpened(w, fanIn, config.Timeout, result)
	}

	for _, val := range address.Sites {
		fanIn <- val
	}
	close(fanIn)
	var validation []Validation
	for i := 1; i <= len(address.Sites); i++ {
		validation = append(validation, <-result)
	}
	fanOut := make(chan Validation, len(validation))
	output := make(chan bool, len(validation))
	for w := 1; w <= config.Workers; w++ {
		go WriteToFile(w, fanOut, output)
	}
	for _, val := range validation {
		fanOut <- val
	}
	for i := 1; i <= len(validation); i++ {
		<-output
	}
	close(fanOut)
	fmt.Println("Completed...")
}

type Config struct {
	Timeout  int    `yaml:"timeout"`
	Workers  int    `yaml:"workers"`
	Filepath string `yaml:"filepath"`
}

type Address struct {
	Sites []string `yaml:"address"`
}

type Validation struct {
	Sites  string
	Status bool
}
