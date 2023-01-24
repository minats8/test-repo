package task1

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const dummyURL string = "https://jsonplaceholder.typicode.com/posts"

// take user input for the below choice:
// 1. all posts
// 2. single posts -> take the post id
// greater than 100
// less than 100

func main() {
	fmt.Println("1.All Post\n2.Single Post\nEnter Your Choice..")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	CheckNilErr(err)
	choice, err := strconv.ParseInt(strings.TrimSpace(input), 10, 0) // if error occurs then exit
	if err != nil {
		fmt.Println("Invalid Choice entered by the User...")
		os.Exit(1)
	}
	switch choice {
	case 1:
		//all post
		LoadPageData("")
	case 2:
		//single post
		fmt.Println("Enter the Post Number you want to fetch : ")
		input, err := reader.ReadString('\n')
		CheckNilErr(err)
		LoadPageData(strings.TrimSpace(input))

	default:
		fmt.Println("No Such Choice Exist..")
	}

}
func CheckNilErr(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error()) // gives the error in string format
		os.Exit(1)
	}
}
func LoadPageData(postNO string) {
	if postNO == "" {
		resp, err := http.Get(dummyURL)
		CheckNilErr(err)
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(content))
		CheckNilErr(err)
		var data []Data
		er := json.Unmarshal(content, &data)
		CheckNilErr(er)
		for _, val := range data {
			fmt.Println(val.UserId, "\t", val.Id, "\t", val.Title)
		}
	} else {
		newURL := dummyURL + "/" + postNO
		resp, err := http.Get(newURL)
		CheckNilErr(err)
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(content))
		CheckNilErr(err)
		var data Data
		err = json.Unmarshal(content, &data)
		CheckNilErr(err)
		if data.Id == 0 {
			// {0 0    } format
			fmt.Println("No Such Post Exist...")
		} else {
			fmt.Println(data.UserId, "\t", data.Id, "\t", data.Title)
		}
	}
}

type Data struct {
	UserId int    `json : "userId"`
	Id     int    `json : "id"`
	Title  string `json : "title"`
	Body   string `json : "body"`
}

/*
	force stop code
	os.Exit(and number other than zero) for unexpected execution
	os.Exit(0) execution went as expected
*/
