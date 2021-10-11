package service

import (
	"context"
	"reflect"
	"test-stockbit/pkg/model"
	"test-stockbit/pkg/repository"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_service_DetailMovieService(t *testing.T) {
	type fields struct {
		repo repository.Repository
		log  *logrus.Logger
	}
	type args struct {
		ctx     context.Context
		request *model.DetailMovieByIDRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &service{
				repo: tt.fields.repo,
				log:  tt.fields.log,
			}
			got, err := service.DetailMovieService(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.DetailMovieService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.DetailMovieService() = %v, want %v", got, tt.want)
			}
		})
	}
}
