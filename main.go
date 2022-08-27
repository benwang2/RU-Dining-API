package main

import (
	"fmt"
	"net/http"

	"github.com/benwang2/ru_dining_api/middleware"
)

func main() {

	http.HandleFunc("/api/menu", middleware.GetMenu)
	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Println(err)
	}

	// middleware.FetchMenuFromURI("")
}
