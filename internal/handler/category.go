package handler

import (
	"fmt"
	"net/http"
)

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Category Page!!")
}
