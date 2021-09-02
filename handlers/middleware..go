package handlers

import (
	"context"
	"net/http"

	"github.com/DanielUlises98/microservices/data"
)

func (p Products) MiddleWareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := data.FromJson(prod, r.Body)
		if err != nil {
			p.l.Println("ERROR desearializing product ", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		//validate the product

		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("ERRO validating product", errs)

			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
