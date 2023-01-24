/*
	1. Create a File
		Conditions for File Creation
			File not present -> so create the file
			File present -> so use existing file only
			Insufficient memory for file creation  -> Exit(1)
	2. Read/Write Operation on File
	3.
*/

package task2

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const FilePath = "log.txt"
const URL = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("1.Create File\n2.View File\n3.Perform Operations\n4.Quit\nEnter Your Choice.. ")
	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	CheckNilError(err)
	input, err := strconv.ParseInt(strings.TrimSpace(choice), 10, 0)
	InvalidChoice(err)
	switch input {
	case 1:
		//check if the file exist or not
		_, err := os.Stat(FilePath)
		if errors.Is(err, os.ErrNotExist) {
			//file does't exist
			CreateFile()
		}
	case 2:
		//view file
		file, err := os.Open(FilePath)
		FileError(err)
		defer file.Close()
		content, err := ioutil.ReadFile(FilePath)
		CheckNilError(err)
		fmt.Println(string(content))
	case 3:
		//writing into the file
		fmt.Println("Select from the Option below\n1.All Post\n2.Fetch Desired Post\n3.Random Post")
		choice, err := reader.ReadString('\n')
		CheckNilError(err)
		input, err := strconv.ParseInt(strings.TrimSpace(choice), 10, 0)
		InvalidChoice(err)
		switch input {
		case 1:
			//all post
			fileData := "User Selected Option 1: Number of Post is " + LoadPageData("") + "\n"
			WriteOnFile(fileData)
		case 2:
			//desired post
			fmt.Println("Enter the PostID: ")
			choice, err := reader.ReadString('\n')
			CheckNilError(err)
			fileData := "User Selected Option 2: " + LoadPageData(strings.TrimSpace(choice)) + "\n"
			WriteOnFile(fileData)
		case 3:
			//random post
			rand.Seed(time.Now().UnixNano())
			randomNumber := rand.Intn(100) + 1
			fileData := "User Selected Option 3: " + LoadPageData(strconv.Itoa(randomNumber)) + "\n"
			WriteOnFile(fileData)
		default:
			fmt.Println("Invalid Option Selected...")
			os.Exit(0)
		}
	case 4:
		fmt.Println("logging Off...")
		os.Exit(0)
	default:
		fmt.Println("Invalid Option Selected...")
		os.Exit(0)
	}
}

func CheckNilError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}

func FileError(err error) {
	if err != nil {
		fmt.Println("File Error: ", err.Error())
		os.Exit(1)
	}
}
func InvalidChoice(err error) {
	if err != nil {
		fmt.Println(err.Error(), "\nInvalid Choice Entered by the User...")
		os.Exit(1)
	}
}

func CreateFile() {
	file, err := os.Create(FilePath)
	CheckNilError(err)
	defer file.Close()
	fmt.Println("File Created Successfully...")
}
func WriteOnFile(fileData string) {
	file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_WRONLY, 0655)
	FileError(err)
	defer file.Close()
	_, err = file.WriteString(fileData)
	CheckNilError(err)
}

func LoadPageData(postNO string) string {
	if postNO == "" {
		resp, err := http.Get(URL)
		CheckNilError(err)
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		CheckNilError(err)
		var data []Data
		er := json.Unmarshal(content, &data)
		CheckNilError(er)
		return strconv.Itoa(len(data))
	} else {
		newURL := URL + "/" + postNO
		resp, err := http.Get(newURL)
		CheckNilError(err)
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		CheckNilError(err)
		var data Data
		err = json.Unmarshal(content, &data)
		CheckNilError(err)
		if data.Id == 0 {
			txt := "ID: " + postNO + "Invalid post requested by the User"
			return txt
		}
		logData := "ID: " + strconv.Itoa(data.Id) + " Title: " + data.Title
		return logData
	}
}

type Data struct {
	UserId int    `json : "userId"`
	Id     int    `json : "id"`
	Title  string `json : "title"`
	Body   string `json : "body"`
}
