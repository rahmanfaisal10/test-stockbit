package decode

import (
	"context"
	"net/http"
	"test-stockbit/pkg/model"
)

func DecodeListMoviesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	paramSearch := r.URL.Query().Get("s")
	paramPage := r.URL.Query().Get("page")

	return model.ListMovieRequest{
		SearchWord: paramSearch,
		Pagination: paramPage,
	}, nil
}
