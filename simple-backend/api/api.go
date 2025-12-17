package api

import (
	"encoding/json"
	"net/http"
)


type CoinBalanceParams struct {
	Username string 
}

type CoinBalanceResponse struct {
	Code int // success code
	Balance int64 // account balance
}

type Error struct {
	Code int // error code response 
	Message string // string
}


// helper function to format the error into nice format
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(code); // set the resposne code

	json.NewEncoder(w).Encode(resp); // serialize the error into json 
}


var (
	RequestErrorHandler = func (w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest);
	}

	InternalErrorHandler = func (w http.ResponseWriter) {
		writeError(w, "An unexpected error occurred.", http.StatusInternalServerError);
	} 
)