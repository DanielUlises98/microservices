package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DanielUlises98/microservices/data"
	"github.com/gorilla/mux"
)

type KeyProduct struct{}
type Products struct {
	l *log.Logger
	v *data.Validation
}
type GenericError struct {
	Message string `json:"message"`
}
type ValidationError struct {
	Messages []string `json:"messages"`
}

func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

var ErrInvalidProductPath = fmt.Errorf("invalid path, path should be /products/[id]")

func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id

}
