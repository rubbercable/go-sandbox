package main

import (
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rubbercable/mysql-book-mgmt/pkg/routes"
	_ "github.com/rubbercable/mysql-book-mgmt/pkg/routes"
	"log"
	_ "log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("Localhist:9010", r))
}
