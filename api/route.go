package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func init() {
	http.Handle("/", router)
}
