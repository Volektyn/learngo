package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 8. Create structs to host the request data.
// 9. Create structs to parse the response data.
// 10. Create an addNameAndParseResponse function that posts a name to the server.
// 11. Create a getDataAndParseResponse function that parses the server response.
// 12. Send a POST request to the server, to add names.
// 13. Send a GET request to the server.
// 14. Parse the response into a struct.
// 15. Loop through the struct and print the names.

var url = "http://localhost:8080"

type ReqData struct {
	Name string `json:"name"`
}

type ResData struct {
	Names []string `json:"names"`
}

type Resp struct {
	OK bool `json:"ok"`
}

func addNameAndParseResponse(nameToAdd string) error {
	name := ReqData{Name: nameToAdd}
	nameBytes, err := json.Marshal(name)
	if err != nil {
		return err
	}
	response, err := http.Post(fmt.Sprintf("%s/addName", url), "text/json", bytes.NewReader(nameBytes))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	respData := Resp{}
	err = json.Unmarshal(data, &respData)
	if err != nil {
		return err
	}
	if !respData.OK {
		return errors.New("response not ok")
	} else {
		fmt.Println("name added")
	}
	return nil
}

func getDataAndParseResponse() []string {
	r, err := http.Get(fmt.Sprintf("%s/", url))
	if err != nil {
		log.Fatal(err)
	}
	// get data from the response body
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	respNames := ResData{}
	err = json.Unmarshal(data, &respNames)
	if err != nil {
		log.Fatal(err)
	}
	return respNames.Names
}

func main() {
	err := addNameAndParseResponse("Electric")
	if err != nil {
		log.Fatal(err)
	}
	err = addNameAndParseResponse("Boogaloo")
	if err != nil {
		log.Fatal(err)
	}
	names := getDataAndParseResponse()
	if len(names) > 0 {
		for _, name := range names {
			log.Println(name)
		}
	} else {
		fmt.Println("no names")
	}

}
