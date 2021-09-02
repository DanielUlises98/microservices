package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DanielUlises98/microservices/data"
	"github.com/gorilla/mux"
)

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

var ErrInvalidProductPath = fmt.Errorf("Invalid path, path should be /products/[id]")

func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id

}
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(prod)
}

type KeyProduct struct{}
