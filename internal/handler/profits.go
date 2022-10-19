package handler

import (
	"fmt"
	"net/http"
)

func ProfitsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Profits Page!!")
}
