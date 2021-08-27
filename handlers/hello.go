package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger

}

func NewHello(l *log.Logger)*Hello{
	return &Hello{l}
}

func (h *Hello) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world")
	d, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	//log.Printf("%+v\n", string(d))

	fmt.Fprintf(rw, "Hello %s", d)
}
