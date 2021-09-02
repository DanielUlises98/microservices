package handlers

import (
	"fmt"
	"net/http"

	"github.com/DanielUlises98/microservices/data"
)

func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	fmt.Printf("DEBUG Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}
