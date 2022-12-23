package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type T []struct {
	id           int    `json: "id"`
	image        string `json: "image"`
	name         string `json: "name"`
	members      string `json: "members"`
	creationDate int    `json: "creationDate"`
	firstAlbum   string `json: "firstAlbum"`
	locations    string `json: "locations"`
	concertDates string `json: "concertDates"`
	relations    string `json: "relations"`
}

var Take T

func Data() {
	response, _ := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	defer response.Body.Close()
	file, _ := ioutil.ReadAll(response.Body)
	h, _ := json.Unmarshal(file, T)
	return h
}

func main() {
	fmt.Println(Data())
}
