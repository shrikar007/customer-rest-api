package crud_interface

import (
	"net/http"
)

type Database interface {
	CreateCustomer(w http.ResponseWriter,r *http.Request)
	GetId(w http.ResponseWriter,r *http.Request)
	GetAll(w http.ResponseWriter,r *http.Request)
	CustomerCtx(next http.Handler) http.Handler
}
