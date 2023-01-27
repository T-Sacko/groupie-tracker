package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Page struct {
	ArtistInfos []Info
}

type Info struct {
	Name         string
	Image        string
	DatesLocs    map[string][]string
	Members      []string
	CreationDate int
	FirstAlbum   string
}

type Artist struct {
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

type ting struct {
	Id        int                 `json:"id"`
	Dateslocs map[string][]string `json:"datesLocations"`
}

type loc struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

var (
	Infos     []Info
	Relations []ting
	Take      []Artist
	Locations loc
)

func main() {
	response, err := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		print(err)
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&Take)
	for i := 1; i <= 52; i++ {
		var next ting
		url := fmt.Sprintf("http://groupietrackers.herokuapp.com/api/relation/%d", i)
		res, er := http.Get(url)
		if err != nil {
			panic(er)
		}
		defer res.Body.Close()
		json.NewDecoder(res.Body).Decode(&next)
		Relations = append(Relations, next)

	}
	// var page Page

	for i := 0; i <= 51; i++ {
		var nextInfo Info = Info{
			Name:         Take[i].Name,
			Image:        Take[i].Image,
			DatesLocs:    Relations[i].Dateslocs,
			Members:      Take[i].Members,
			CreationDate: Take[i].CreationDate,
			FirstAlbum:   Take[i].FirstAlbum}
		Infos = append(Infos, nextInfo)
	}
	// page = Page{ArtistInfos: infos}
	fmt.Println(Relations[5])

	respons, er := http.Get("http://groupietrackers.herokuapp.com/api/relation/1")
	if er != nil {
		print(er)
	}
	defer respons.Body.Close()
	json.NewDecoder(respons.Body).Decode(&Relations)

	http.HandleFunc("/", groupie)
	http.ListenAndServe(":5505", nil)
}

func groupie(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	var err error

	if len(r.URL.Path) > 1 && len(r.URL.Path) <= 3 {
		n := r.URL.Path[1:]
		sn, _ := strconv.Atoi(n)
		fmt.Println(n)
		fmt.Println(len(r.URL.Path))
		if sn <= 52 && sn >= 0 {
			tmp, err := template.ParseFiles("id.html")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 Internal Server Error"))

			} else {
				tmp.Execute(w, Infos[sn-1])
			}
		}
	} else if r.URL.Path == "/" {
		tmpl, err = template.ParseFiles("index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 Internal Server Error"))
		} else {
			tmpl.Execute(w, Take)
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<html><body>404 error</body></html>")
	}
	// fmt.Println("gogo")
	nn := len(r.URL.Path) - 1
	n := r.URL.Path[nn:]
	sn, _ := strconv.Atoi(n)
	tmp, _ := template.ParseFiles("id.html")
	tmp.Execute(w, Infos[sn])

}
