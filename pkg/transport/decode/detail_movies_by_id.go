package decode

import (
	"context"
	"net/http"
	"test-stockbit/pkg/model"
)

func DecodeDetailMoviesByIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	paramID := r.URL.Query().Get("i")

	return model.DetailMovieByIDRequest{
		ID: paramID,
	}, nil
}
