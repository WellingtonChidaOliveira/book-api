package main

import (
	"github.com/welligtonchida/book-api/internal/http/api"
	"github.com/welligtonchida/book-api/internal/http/routes"
)

func main() {
	h := routes.Handlers()
	err := api.Start("8080", h)
	if err != nil {
		panic(err)
	}

}
