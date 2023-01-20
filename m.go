package main

import (
	"encoding/json"
	"fmt"
	"html/template"
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

type loc struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

func main() {
	response, err := http.Get("http://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		print(err)
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&Take)
	fmt.Println(Take[0].Name)
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
