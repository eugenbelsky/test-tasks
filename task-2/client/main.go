package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// type Foo struct {
// 	Bar string
// }

var myClient = &http.Client{Timeout: 10 * time.Second}

func main() {

	projects := getJson("http://localhost:8080/api/v1/programs")
	var projectsValid []map[string]interface{}
	for _, object := range projects {
		if object["plus_only"] == true && object["rating"].(float64) <= 16 {
			projectsValid = append(projectsValid, object)
		}
	}

	fmt.Println("--- Valid programs ---")
	for _, object := range projectsValid {
		fmt.Println(object)
	}

	fmt.Println("--- Valid programs` names ---")
	for _, object := range projectsValid {
		fmt.Println(object["name"])
	}

	fmt.Println("--- Total ---")
	fmt.Println("All programs: " + strconv.Itoa(len(projects)))
	fmt.Println("Valid programs: " + strconv.Itoa(len(projectsValid)))

}

func getJson(url string) []map[string]interface{} {
	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	b, _ := io.ReadAll(r.Body)
	var objects []map[string]interface{}
	err = json.Unmarshal(b, &objects)
	if err != nil {
		log.Fatal()
	}

	return objects

}
