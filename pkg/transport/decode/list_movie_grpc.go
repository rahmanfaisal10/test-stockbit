package decode

import (
	"context"
	"errors"

	"test-stockbit/pkg/model"
	"test-stockbit/pkg/proto"
)

func DecodeListMoviesGRPC(ctx context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*proto.ListMovieRequest)
	if !ok {
		return nil, errors.New("invalid type")
	}
	return &model.ListMovieRequest{
		Apikey:     req.Apikey,
		SearchWord: req.SearchWord,
		Pagination: req.Pagination,
	}, nil
}
