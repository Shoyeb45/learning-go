package handlers

import (
	"encoding/json"
	"net/http"
	"simple-backend/api"
	"simple-backend/internal/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{};
	var decoder *schema.Decoder = schema.NewDecoder();
	
	var err error;

	err = decoder.Decode(&params, r.URL.Query());

	if err != nil {
		log.Error(err);
		api.InternalErrorHandler(w);
		return;
	}

	var database *tools.DatabaseInterface;
	database, err = tools.NewDatabase();

	if err != nil {
		log.Error(err);
		api.InternalErrorHandler(w);
		return;
	}

	var tokenDetails *tools.CoinDetails = (*database).GetUserCoins(params.Username);

	if tokenDetails == nil {
		log.Error(err);
		api.InternalErrorHandler(w);
		return;
	}

	// set the value to the response struct

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code: http.StatusOK,
	}

	// write the response to the response writer

	w.Header().Set("Content-Type", "application/json");
	err = json.NewEncoder(w).Encode(response);

	if err != nil {
		log.Error(err);
		api.InternalErrorHandler(w);
		return;
	}

}