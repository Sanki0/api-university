package main

import (
	"github.com/Sanki0/api-university/api-routes"
	"github.com/Sanki0/api-university/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	utils.InitDB()

	apiroutes.InitApiRoutes(r)

}
