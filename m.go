package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var (
	Take T
	Data []U
)

type T []struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type U struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func groupie(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, struct{ Take T }{Take})

}

func main() {
	response, err := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		print(err)
	}
	defer response.Body.Close()
	file, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(file, &Take)
	fmt.Println(GetData()[0].Image)
	http.HandleFunc("/", groupie)
	http.ListenAndServe(":5505", nil)

}

func GetData() []U {

	// make sure to manage the error and return and Http.Status(Internalservererror)
	for _, v := range Take {
		Data = append(Data, U{
			Id:           v.Id,
			Image:        v.Image,
			Name:         v.Name,
			Members:      v.Members,
			CreationDate: v.CreationDate,
			FirstAlbum:   v.FirstAlbum,
			Locations:    v.Locations,
			ConcertDates: v.ConcertDates,
			Relations:    v.Relations,
		})
	}
	return (Data)
}
