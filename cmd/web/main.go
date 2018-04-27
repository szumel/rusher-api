package main

import (
	"github.com/szumel/rusher-api/cmd/web/macro"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/macro", macro.Get)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
