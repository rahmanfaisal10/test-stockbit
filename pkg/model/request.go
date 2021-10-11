package model

type ListMovieRequest struct {
	Apikey     string `json:"apikey"`
	SearchWord string `json:"s"`
	Pagination string `json:"page"`
}

type DetailMovieByIDRequest struct {
	Apikey string `json:"apikey"`
	ID     string `json:"i"`
	Plot   string `json:"plot"`
}

type DetailMovieByTitleRequest struct {
	Apikey string `json:"apikey"`
	Title  string `json:"t"`
	Year   string `json:"y"`
	Plot   string `json:"plot"`
}
