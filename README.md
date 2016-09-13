# shade

[![Build Status](https://travis-ci.org/i2bskn/shade.svg?branch=master)](https://travis-ci.org/i2bskn/shade)

shade is thin wrapper of `html/template`.

Example
-------

```go
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
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
```