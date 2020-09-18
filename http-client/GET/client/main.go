package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Names struct {
	Names []string `json:"names"`
}

func getDataAndParseResponse() (int, int) {
	r, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	rData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	names := Names{}
	err = json.Unmarshal(rData, &names)
	if err != nil {
		log.Fatal(err)
	}

	electricCount := 0
	boogalooCount := 0

	for _, v := range names.Names {
		if strings.Trim(v, " ") == "Electric" {
			electricCount++
		} else if strings.Trim(v, " ") == "Boogaloo" {
			boogalooCount++
		}
	}
	return electricCount, boogalooCount
}
func main() {
	electricCount, boogalooCount := getDataAndParseResponse()
	fmt.Println("Electric Count: ", electricCount)
	fmt.Println("Boogaloo Count: ", boogalooCount)
}
