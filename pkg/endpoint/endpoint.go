package endpoint

import (
	"context"
	"test-stockbit/pkg/model"
	"test-stockbit/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ListMovies endpoint.Endpoint
}

func MakeEndpoints(svc service.Service) Endpoints {
	return Endpoints{
		ListMovies: MakeListMovies(svc),
	}
}

func MakeListMovies(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.ListMovieRequest)
		return svc.ListMovieService(ctx, &req)
	}
}
