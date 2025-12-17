package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"simple-backend/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	// it will log the line number
	log.SetReportCaller(true);

	var r *chi.Mux = chi.NewRouter();

	handlers.Handler(r);

	fmt.Println("Starting GO API Services....");
	err := http.ListenAndServe("localhost:8000", r);

	if (err != nil) {
		log.Error(err);
	}
}