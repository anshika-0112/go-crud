package main

import (
	"net/http"

	"github.com/anshika-0112/go-crud/database"
	"github.com/anshika-0112/go-crud/product"
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {
	database.SetUpDatabase()
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":8000", nil)
}
