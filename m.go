package main

import (
	"encoding/json"
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

type U struct {
	Id           int    `json: "id"`
	Image        string `json: "image"`
	Name         string `json: "name"`
	Members      string `json: "members"`
	CreationDate int    `json: "creationDate"`
	FirstAlbum   string `json: "firstAlbum"`
	Locations    string `json: "locations"`
	ConcertDates string `json: "concertDates"`
	Relations    string `json: "relations"`
}

var (
	Take T
	Art  []U
)

func Data() {
	response, _ := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	defer response.Body.Close()
	file, _ := ioutil.ReadAll(response.Body)
	h := json.Unmarshal(file, &Take)
	// make sure to manage the error and return and Http.Status(Internalservererror)
	for _, v := range Take {
		Art = append(Art, U{
			Id:           v.ID,
			Image:        v.Image,
			Name:         v.Name,
			Members:      v.Members,
			CreationDate: v.CreationDate,
			FirstAlbum:   v.FirstAlbum,
			Locations:    v.Locations,
			ConcertDates: v.ConcertDates,
			Relations:    v.Relations
		})
	}
}

func main() {
	// fmt.Println(Data())
}
