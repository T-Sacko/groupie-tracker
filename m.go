package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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

var (
	Take      T
	Locations loc
)

type loc []struct {
	Id        int    `json:"id"`
	Dates     int    `json:"dates"`
	Locations string `json:"locations"`
}

func main() {
	response, err := http.Get("http://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		print(err)
	}
	defer response.Body.Close()
	file, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(file, &Take)
	fmt.Println(Take[0].Image)
	http.HandleFunc("/", groupie)
	http.ListenAndServe(":5505", nil)

}

func groupie(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	if r.URL.Path == "/" {
		tmpl, _ = template.ParseFiles("index.html")
		tmpl.Execute(w, struct{ Take T }{Take})

	} else {
		fmt.Println("gogo")
		tmp, _ := template.ParseFiles("id.html")
		tmp.Execute(w, nil)
		// locs, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
		// if err != nil {
		// 	print(err)
		// }
		// defer locs.Body.Close()
		// file, _ := ioutil.ReadAll(locs.Body)
		// _ = json.Unmarshal(file, &Locations)
		// dates, errr := http.Get("https://groupietrackers.herokuapp.com/api/dates")
		// if errr != nil {
		// 	print(err)
		// }
		// defer locs.Body.Close()
		// filedates, _ := ioutil.ReadAll(dates.Body)
		// _ = json.Unmarshal(filedates, &Locations)
		// fmt.Println(Locations)
	}

}

// func GetData() []U {

// 	// make sure to manage the error and return and Http.Status(Internalservererror)
// 	for _, v := range Take {
// 		Data = append(Data, U{
// 			Id:           v.Id,
// 			Image:        v.Image,
// 			Name:         v.Name,
// 			Members:      v.Members,
// 			CreationDate: v.CreationDate,
// 			FirstAlbum:   v.FirstAlbum,
// 			Locations:    v.Locations,
// 			ConcertDates: v.ConcertDates,
// 			Relations:    v.Relations,
// 		})
// 	}
// 	return (Data)
// }
