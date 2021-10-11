package util

import (
	"context"
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Success bool `json:"success"`
}

// EncodeLegacyError is used to mirror MPG1's error response format
func EncodeLegacyError(_ context.Context, err error, w http.ResponseWriter) {
	parsedErr, ok := err.(*Error)
	if ok {
		if parsedErr.Code/100 == http.StatusUnauthorized {
			w.Header().Set("Connection", `close`)
			w.Header().Set("X-Content-Type-Options", `nosniff`)
			w.Header().Set("Www-Authenticate", `Basic realm="Authorization Required"`)
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorized"))
			return
		}
	}

	w.Header().Set(CONTENTTYPEKEY, CONTENTTYPEJSONANDCHARSETUTF8VALUE)

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func EncodeSuccessResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(CONTENTTYPEKEY, CONTENTTYPEJSONANDCHARSETUTF8VALUE)
	resp := successResponse{Success: true}
	return json.NewEncoder(w).Encode(resp)
}

func EncodeResponseWithData(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(CONTENTTYPEKEY, CONTENTTYPEJSONANDCHARSETUTF8VALUE)
	return json.NewEncoder(w).Encode(response)
}

func EncodeResponseNoData(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func EncodeResponseWithCount(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(CONTENTTYPEKEY, CONTENTTYPEJSONANDCHARSETUTF8VALUE)
	return json.NewEncoder(w).Encode(response)
}
