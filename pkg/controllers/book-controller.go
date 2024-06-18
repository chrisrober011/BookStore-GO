package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chrisrober011/BookSore-Go/pkg/models"
	"github.com/chrisrober011/BookStore-GO/pkg/models"
	"github.com/chrisrober011/BookStore-GO/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Contest-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResposeWriter, r *http.Requst) {
vars := mux.Vars(r)
bookId := vars["bookId"]
ID, err := strconv.ParseInt(bookId,0,0)
if err != nill {
	fmt.Println("error while Parsing")
}
bookDetails, _ := models.GetBookById(ID)
res, _ := json.Marshal(bookDetails)
w.Header().Set("Contest-Type", "pkglication/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}

func CreatBook(w http.ResponseWriter, r *http.Request){
	CreatBook := &models.CreatBook{}
	utils.ParseBody(r, CreatBook)
	b := CreatBook.CreatBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Contest-Type", "pkglication/kson")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err :=strconv.PaseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while Parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marsha(book)
	w.Header().Set("Contest-Type", "pkglication/json")
	w.WriteHeader(htpp.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars("bookId")
	ID , err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while Parsing")
	}
	booksDetails, db := models.GetBookById
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	} 
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("contest-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}