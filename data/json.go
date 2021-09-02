package data

import (
	"encoding/json"
	"io"
)

func (p *Products) ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func (p *Product) FromJson(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}
