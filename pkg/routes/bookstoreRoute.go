package routes

import (
	"github.com/gorilla/mux"
	"github.com/miladfallah/go-bookstore/pkg/controller"
)

var registerBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}/", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}/", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}/", controller.DeleteBook).Methods("DELETE")
}
