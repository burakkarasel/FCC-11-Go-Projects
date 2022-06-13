package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/burakkarasel/bookstore/pkg/models"
	"github.com/burakkarasel/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

// Here we get all of our books from db by models package's getallbooks func and then we marshal them and send them as response
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// here we use the id given in the path and find the relevant book from db and return it as a response
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing:", err)
	}
	bookDetails, _ := models.GetBookById(ID) // we used _ because we dont need the return value as DB we just need relevant book

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Here we create a book according to the request we create a new book and parse the request into it then we marshal it and send back as a result
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{} // here we create a book by using models package
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook() // here we refer to our CreateBook receiver func inside models package

	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// here we take id from the request path and delete it with DeleteBook func from models package and then return as a resp
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing:", err)
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// here we created a new book by using Book struct from models pack
	var updateBook = &models.Book{}
	// here we parsed body of request and write it into updateBook
	utils.ParseBody(r, updateBook)
	// here we find the ID of our book from request's path
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing:", err)
	}

	// here we brought the book in the DB
	bookDetails, db := models.GetBookById(ID)
	// here we check if values we parsed from request's body are valid or not
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Author = updateBook.Publication
	}
	// here we save the updated version of our book to DB
	db.Save(&bookDetails)
	// and here we send updated version as a resp
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
