package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	ArtistInfos []Info
}

type Info struct {
	Name         string
	Image        string
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	DatesLocs    map[string][]string
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
		res, _ := http.Get(url)
		defer res.Body.Close()
		json.NewDecoder(res.Body).Decode(&next)
		Relations = append(Relations, next)

	}
	var page Page
	var infos []Info
	for i := 0; i <= 51; i++ {
		var nextInfo Info = Info{Name: Take[i].Name, Image: Take[i].Image, DatesLocs: Relations[i].Dateslocs}
		infos = append(infos, nextInfo)
	}
	page = Page{ArtistInfos: infos}
	fmt.Println(Relations[5])

	respons, er := http.Get("http://groupietrackers.herokuapp.com/api/relation/1")
	if er != nil {
		print(er)
	}
	defer respons.Body.Close()
	json.NewDecoder(respons.Body).Decode(&Relations)
	fmt.Println(Relations)

	http.HandleFunc("/", groupie)
	http.ListenAndServe(":5505", nil)
}

func groupie(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	var err error

	// nn := len(r.URL.Path) - 1
	// n := r.URL.Path[nn:]
	// // sn, _ := strconv.Atoi(n)

	// if r.URL.Path == "/"+n {
	// 	tmpl, err = template.ParseFiles("id.html")

	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		w.Write([]byte("500 Internal Server Error"))
	// 	} else {
	// 		err = tmpl.Execute(w,Dates, Locations, Take)
	// 	}
	// }

	if r.URL.Path == "/" {
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
	// nn := len(r.URL.Path) - 1
	// n := r.URL.Path[nn:]
	// sn, _ := strconv.Atoi(n)
	// locs, err := http.Get(Take[sn-1].Locations)
	// if err != nil {
	// 	print(err)
	// }
	// defer locs.Body.Close()
	// file, _ := ioutil.ReadAll(locs.Body)
	// _ = json.Unmarshal(file, &Locations)
	// tmp, _ := template.ParseFiles("id.html")
	// tmp.Execute(w, Locations)
	// fmt.Println(Locations.Id)

	// }
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
