package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/DanielUlises98/microservices/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}
type GenericError struct {
	Message string `json:"message"`
}
type ValidationError struct {
	Message string `json:"message"`
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
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

func (p Products) MiddleWareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "post Unable to marshal json", http.StatusBadRequest)
			return
		}

		//validate the product

		err = prod.Validate()
		if err != nil {
			p.l.Println("Error", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
