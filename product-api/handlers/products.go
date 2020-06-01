package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Akanibekuly/Microservices/product-api/data"
	"github.com/gorilla/mux"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// Products handler for getting and updating products
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts creates a products handler with the given logger and validation
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid path, path should be path/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationErorr is a collection of validation error messages
type ValidationErorr struct {
	Messages []string `json:"messages"`
}

// getProductID returns returns the product id from the URL
// Panics if cannot convert the id into integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	//parse the product ID from the URL
	vars := mux.Vars(r)

	//convert the id into the integer and return
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id)
	if err != nil {
		//should never happen
		panic(err)
	}
	return id
}
