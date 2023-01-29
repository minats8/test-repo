package filesystem

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const FilePath string = "/Users/sidhantdas/Desktop/Task4/fileSystem/log.txt"

func FetchData() error {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("File Error: ", err.Error())
		return err
	}
	defer file.Close()
	content, err := ioutil.ReadFile(FilePath)
	if err != nil {
		fmt.Println("Read File Error: ", err.Error())
		return err
	}
	fmt.Println(string(content))
	return nil
}

func CreateFile() {
	_, err := os.Stat(FilePath)
	if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(FilePath)
		if err != nil {
			fmt.Println("File Error: ", err.Error())
			os.Exit(1)
		}
		defer file.Close()
	}
}

func WriteOnFile(fileData string) error {
	CreateFile()
	file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		fmt.Println("File Opening Error: ", err.Error())
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fileData)
	if err != nil {
		fmt.Println("File Write Error: ", err.Error())
		return err
	}
	return nil
}
