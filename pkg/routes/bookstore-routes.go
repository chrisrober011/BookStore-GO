/* this help to visualize the routes where the users will hit 
from their front end and postman and the user will hits this routes */
package routes

import (
	"github.com/gorilla/mux"
	"github.com/chrisrober011/BookStore-GO/pkg/controllers"

)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	
}