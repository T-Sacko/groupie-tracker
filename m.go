package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Info struct {
	Name         string
	Image        string
	DatesLocs    map[string][]string
	Members      []string
	CreationDate int
	FirstAlbum   string
}

type Artist []struct {
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
	Take      Artist
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

		var nextInfo Info = Info{
			Name:         Take[i-1].Name,
			Image:        Take[i-1].Image,
			DatesLocs:    Relations[i-1].Dateslocs,
			Members:      Take[i-1].Members,
			CreationDate: Take[i-1].CreationDate,
			FirstAlbum:   Take[i-1].FirstAlbum,
		}
		Infos = append(Infos, nextInfo)

	}

	fmt.Println(len(Relations))

	http.HandleFunc("/", groupie)
	http.ListenAndServe(":5505", nil)
}

func groupie(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	var err error

	if len(r.URL.Path) > 1 && len(r.URL.Path) <= 3 {
		n := r.URL.Path[1:]
		sn, err := strconv.Atoi(n)
		if err != nil || sn > 52 || sn < 1 {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "<html><body>404 error</body></html>")
		} else if sn <= 52 && sn >= 1 {
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
}
