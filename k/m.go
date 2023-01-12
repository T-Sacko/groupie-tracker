package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	fmt.Println(struct{ People []Person }{people})
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, struct{ People []Person }{people})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8000", nil)
}
