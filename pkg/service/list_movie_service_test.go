package service

import (
	"context"
	"reflect"
	"test-stockbit/pkg/model"
	"test-stockbit/pkg/repository"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_service_ListMovieService(t *testing.T) {
	type fields struct {
		repo repository.Repository
		log  *logrus.Logger
	}
	type args struct {
		ctx     context.Context
		request *model.ListMovieRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ListMovieResponse
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
			got, err := service.ListMovieService(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ListMovieService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.ListMovieService() = %v, want %v", got, tt.want)
			}
		})
	}
}
