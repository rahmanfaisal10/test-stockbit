package transport

import (
	"fmt"
	"net/http"
	"test-stockbit/config"
	"test-stockbit/pkg/endpoint"
	"test-stockbit/pkg/repository"
	"test-stockbit/pkg/service"
	"test-stockbit/pkg/transport/decode"
	"test-stockbit/util"

	kithttp "github.com/go-kit/kit/transport/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func MakeHandler(r *mux.Router) http.Handler {
	log := logrus.New()

	cfg := config.Get()

	connection := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBConnection,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sqlx.Connect("mysql", connection)
	if err != nil {
		log.Error(err)
	}

	repo := repository.NewRepository(db, log)
	svc := service.NewService(repo, log)
	endpoint := endpoint.MakeEndpoints(&svc)

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(util.EncodeLegacyError),
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
	}

	listMovieEndpoint := endpoint.ListMovies
	listMovieHandler := kithttp.NewServer(
		listMovieEndpoint,
		decode.DecodeListMoviesRequest,
		util.EncodeResponseWithData,
		opts...,
	)

	detailMovieByIdEndpoint := endpoint.DetailMoviesByID
	detailMovieByHandler := kithttp.NewServer(
		detailMovieByIdEndpoint,
		decode.DecodeDetailMoviesByIDRequest,
		util.EncodeResponseWithData,
		opts...,
	)

	r.Handle("/omdbapi/list-movie", listMovieHandler).Methods("GET")
	r.Handle("/omdbapi/detail-movie-by-id", detailMovieByHandler).Methods("GET")
	return r
}
