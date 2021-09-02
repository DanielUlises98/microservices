package handlers

import (
	"net/http"

	"github.com/DanielUlises98/microservices/data"
)

func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id, err := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	http.Error(rw, "Unable to convert ID ", http.StatusBadRequest)
	// }

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("DEBUG updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("ERROR product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found int database"}, rw)
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}
