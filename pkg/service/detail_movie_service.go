package service

import (
	"context"
	"fmt"
	"test-stockbit/pkg/model"
	"test-stockbit/util"
	"time"

	"github.com/sirupsen/logrus"
)

func (service *service) DetailMovieService(ctx context.Context, request *model.DetailMovieByIDRequest) (interface{}, error) {
	logg := &model.Logger{
		ID:        0,
		Transport: logrus.FieldKeyFunc,
		Level:     logrus.GetLevel().String(),
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
	}

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=faf7e5bb&i=%s", request.ID)
	resp := new(interface{})

	err := util.SendGetRequest(url, resp)
	if err != nil {
		logg.Logging = "gagal get url " + url
		logg.Level = service.log.GetLevel().String()
		service.repo.InsertLoggerRepository(logg)
		return nil, err
	}

	logg.Level = service.log.GetLevel().String()
	err = service.repo.InsertLoggerRepository(logg)
	if err != nil {
		service.log.Error(err)
		return nil, err
	}
	return resp, nil
}
