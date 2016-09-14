package main

import (
	"fmt"
	"net/http"

	"github.com/i2bskn/shade"
)

var layout *shade.Layout

func init() {
	layout = shade.Default()
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := layout.Render("index.html", nil)
	fmt.Fprint(w, body)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := layout.Render("index.html", "Hello World!")
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, body)
	})
	http.ListenAndServe(":3000", nil)
}
