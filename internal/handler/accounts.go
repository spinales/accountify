package handler

import (
	"fmt"
	"net/http"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Accounts Page!!")
}
