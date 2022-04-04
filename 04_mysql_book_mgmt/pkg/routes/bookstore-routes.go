package routes

import (
	"github.com/gorilla/mux"
	"github.com/rubbercable/mysql-book-mgmt/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router mux.Router) {
	router.HandleFunc("/book/", controllers.createBook).Methods("POST")
	router.HandleFunc("/book/", controllers.getBook).Methods("GET")
	router.HandleFunc("/book/", controllers.getBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.getBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.updateBook).Method("PUT")
	router.HandleFunc("/book/{bookId}", controllers.deleteBooks).Method("DELETE")

}
