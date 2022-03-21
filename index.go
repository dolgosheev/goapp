package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
)

type ViewData struct {
	I8Title string
	I8Min   int8
	I8Max   int8

	I16Title string
	I16Min   int16
	I16Max   int16
}

func main() {

	http.HandleFunc("/integers", func(w http.ResponseWriter, r *http.Request) {

		data := ViewData{
			I8Title:  "INT8 Interval",
			I8Min:    math.MinInt8,
			I8Max:    math.MaxInt8,
			I16Title: "INT16 Interval",
			I16Min:   math.MinInt16,
			I16Max:   math.MaxInt16,
		}

		tmpl, _ := template.ParseFiles("www/integers.html")
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "www/index.html")
	})

	/* console */
	fmt.Println("Server is listening...")

	/* settings */
	http.ListenAndServe("0.0.0.0:3001", nil)
}
