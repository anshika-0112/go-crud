package main

import (
	"net/http"

	"github.com/anshika-0112/go-crud/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":8000", nil)
}
