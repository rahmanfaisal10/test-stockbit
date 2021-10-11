package main

import (
	"fmt"
	"io"
	"net/http"
	"test-stockbit/config"
	"test-stockbit/pkg/transport"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.Get()
	fmt.Println(cfg)

	log.Info().Msgf("Start listening on port : %s", cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), corsHandler())
	if err != nil {
		log.Fatal().Err(err).Msg("Error when starting!!")
	}
}

func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})

	v1 := r.PathPrefix("/api/v1").Subrouter()
	transport.MakeHandler(v1)

	return r
}

func corsHandler() http.Handler {
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(header, methods, origins)(initRouter())
}
