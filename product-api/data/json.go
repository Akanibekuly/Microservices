package data

import (
	"encoding/json"
	"io"
)

// ToJSON serialize the given interface into a string based JSON format
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(p)
}

// FromJSON deserialize the object from JSON string
// in an io.Reader to the given interface
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
