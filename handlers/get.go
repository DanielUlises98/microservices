package handlers

import (
	"net/http"

	"github.com/DanielUlises98/microservices/data"
)

func (p *Products) ListALL(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("DEBUG get all records")

	prods := data.GetProducts()
	if err != nil {
		p.l.Println("ERROR serializing product", err)
	}
}
