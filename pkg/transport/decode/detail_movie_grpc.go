package decode

import (
	"context"
	"errors"
	"test-stockbit/pkg/model"
	"test-stockbit/pkg/proto"
)

func DecodeDetailMoviesGRPC(ctx context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*proto.DetailMovieByIDRequest)
	if !ok {
		return nil, errors.New("invalid type")
	}
	return &model.DetailMovieByIDRequest{
		Apikey: req.Apikey,
		ID:     req.ID,
		Plot:   req.Plot,
	}, nil
}
