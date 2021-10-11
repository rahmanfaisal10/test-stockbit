package endpoint

import (
	"context"
	"test-stockbit/pkg/model"
	"test-stockbit/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ListMovies       endpoint.Endpoint
	DetailMoviesByID endpoint.Endpoint
}

func MakeEndpoints(svc service.Service) Endpoints {
	return Endpoints{
		ListMovies:       MakeListMovies(svc),
		DetailMoviesByID: MakeDetailMoviesByID(svc),
	}
}

func MakeListMovies(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.ListMovieRequest)
		return svc.ListMovieService(ctx, &req)
	}
}

func MakeDetailMoviesByID(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.DetailMovieByIDRequest)
		return svc.DetailMovieService(ctx, &req)
	}
}
