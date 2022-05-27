package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Struct for Versions
type Versions struct {
	Version     string
	Date        string
	Description string
}

//Struct for Data
type Data struct {
	Versions []Versions
}

//Struct for msg
type Message struct {
	Code int
	Msg  string
	Data Data
	Time uint64
}

func main() {
	var content []byte
	var err error
	var fileName string

	for {
		//Determine JSON filename from user input
		fmt.Println("Please input a JSON file.")

		fmt.Scan(&fileName)

		//Read JSON File
		if len(fileName) > 4 {
			if fileName[len(fileName)-5:] == ".txt" {
				content, err = ioutil.ReadFile(fileName)
			} else {
				content, err = ioutil.ReadFile(fileName + ".txt")
			}
		} else {
			content, err = ioutil.ReadFile(fileName + ".txt")
		}
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	//Unpackage the JSON message based on struct created
	var msg Message
	json.Unmarshal([]byte(content), &msg)

	//String Formatting
	Code := msg.Code
	MessageContent := msg.Msg
	Version := msg.Data.Versions[0].Version
	Date := msg.Data.Versions[0].Date
	Description := msg.Data.Versions[0].Description
	Time := msg.Time

	fmt.Printf("Code: %v\nMessageContent: %v\n Version: %v\n Date: %v\n Description: %v\nTime: %v\n", Code, MessageContent, Version, Date, Description, Time)

}
