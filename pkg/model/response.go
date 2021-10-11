package model

type ListMovieResponse struct {
	Search       []interface{} `json:"Search"`
	TotalResults string        `json:"totalResults"`
	Response     string        `json:"Response"`
}
