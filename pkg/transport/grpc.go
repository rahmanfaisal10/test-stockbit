package transport

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"test-stockbit/config"
	"test-stockbit/pkg/endpoint"
	"test-stockbit/pkg/model"
	"test-stockbit/pkg/proto"
	"test-stockbit/pkg/repository"
	"test-stockbit/pkg/service"
	"test-stockbit/pkg/transport/decode"
	"test-stockbit/util"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

//GrpcServer ....
type GrpcServer struct {
	DetailMoviesByID grpctransport.Handler
	ListMovies       grpctransport.Handler
}

//GRPCServerRun run grpc server
func GRPCServerRun(
	addr string,
	recruitmentSrv proto.ServiceServer,
) {
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}

	var opts []grpc.ServerOption
	// opts = apmgrpc.GetElasticAPMServerOptions()

	go func() {
		baseServer := grpc.NewServer(opts...)

		proto.RegisterServiceServer(baseServer, recruitmentSrv)

		logrus.Info("🚀 Server recruitment started successfully 🚀 - Running on", addr)
		baseServer.Serve(grpcListener)
	}()

	logrus.Error("exit", <-errs)
}

func NewGRPCServer() *GrpcServer {
	log := logrus.New()

	cfg := config.Get()

	connection := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBConnection,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sqlx.Connect("mysql", connection)
	if err != nil {
		log.Error(err)
	}

	repo := repository.NewRepository(db, log)
	service := service.NewService(repo, log)

	endpoint := endpoint.MakeEndpoints(&service)

	handlerOpt := []grpctransport.ServerOption{
		// grpctransport.ServerBefore(jwt.GRPCToContext()),
	}

	return &GrpcServer{
		ListMovies:       grpctransport.NewServer(endpoint.ListMovies, decode.DecodeListMoviesGRPC, encodeListMoviesResponse, handlerOpt...),
		DetailMoviesByID: grpctransport.NewServer(endpoint.DetailMoviesByID, decode.DecodeDetailMoviesGRPC, encodeDetailMoviesResponse, handlerOpt...),
	}
}

func encodeListMoviesResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(model.ListMovieResponse)
	if !ok {
		return nil, errors.New("invalid type")
	}

	search, err := util.GetBytes(resp.Search)
	if err != nil {
		return nil, err
	}

	return &proto.ListMovieResponse{
		Search:       search,
		TotalResults: resp.TotalResults,
		Response:     resp.Response,
	}, nil
}

func encodeDetailMoviesResponse(ctx context.Context, response interface{}) (interface{}, error) {
	search, err := util.GetBytes(response)
	if err != nil {
		return nil, err
	}

	return search, nil
}
