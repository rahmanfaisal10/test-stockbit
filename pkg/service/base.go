package service

import (
	"context"
	"test-stockbit/pkg/model"
	"test-stockbit/pkg/repository"

	"github.com/sirupsen/logrus"
)

type service struct {
	repo repository.Repository
	log  *logrus.Logger
}

func NewService(repo repository.Repository, log *logrus.Logger) service {
	return service{
		repo: repo,
		log:  log,
	}
}

type Service interface {
	ListMovieService(ctx context.Context, request *model.ListMovieRequest) (*model.ListMovieResponse, error)
}
