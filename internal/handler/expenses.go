package handler

import (
	"fmt"
	"net/http"
)

func ExpensesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Expenses Page!!")
}
