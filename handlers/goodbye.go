package handlers

import (
	"log"
	"net/http"
)

//Goodbye is
type Goodbye struct {
	l *log.Logger
}

//NewGoodbye returns new var Goodbye
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

//ServeHTTP serves our handle
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byee"))
}
