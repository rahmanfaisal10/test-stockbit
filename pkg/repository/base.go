package repository

import (
	"test-stockbit/pkg/model"

	"github.com/jmoiron/sqlx"
	logrus "github.com/sirupsen/logrus"
)

type Repository interface {
	InsertLoggerRepository(logger *model.Logger) error
}

type repository struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewRepository(db *sqlx.DB, log *logrus.Logger) Repository {
	return &repository{
		db:  db,
		log: log,
	}
}
