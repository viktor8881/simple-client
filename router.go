package main

import (
	"context"
	"github.com/spf13/viper"
	simpleClient "github.com/viktor8881/service-utilities/http/client"
	"github.com/viktor8881/service-utilities/http/server"
	"go.uber.org/zap"
	"net/http"
	genclient "simple-client/generated/http/client"
	generated "simple-client/generated/http/server"
	"simple-client/inner/user"
	"simple-client/middleware"
	"time"
)

func RegisterRoutes(ctx context.Context, mux *http.ServeMux, logger *zap.Logger,
) []func() {
	tr := server.NewTransport(mux)

	client, cFunc := GetExternalClient(logger)

	userService := user.NewService(client)

	generated.ListUserEndpoint(
		tr,
		userService.ListUser,
		logger,
		server.LoggerMiddleware(logger),
	)

	generated.ListUserByEmailEndpoint(
		tr,
		userService.ListUserByEmail,
		logger,
		server.LoggerMiddleware(logger),
	)

	generated.GetUserEndpoint(
		tr,
		userService.GetUser,
		logger,
		server.LoggerMiddleware(logger),
	)

	generated.CreateUserEndpoint(
		tr,
		userService.CreateUser,
		logger,
		server.LoggerMiddleware(logger),
	)

	generated.UpdateUserEndpoint(
		tr,
		userService.UpdateUser,
		logger,
		server.LoggerMiddleware(logger),
	)

	generated.DeleteUserEndpoint(
		tr,
		userService.DeleteUser,
		logger,
		server.LoggerMiddleware(logger),
	)

	return []func(){
		cFunc,
	}
}

func GetExternalClient(logger *zap.Logger) (*genclient.Client, func()) {
	transport := &http.Transport{
		IdleConnTimeout: 90 * time.Second,
	}

	sClient := simpleClient.NewSimpleClient(
		viper.GetString("external_services.user.base_url"),
		viper.GetDuration("external_services.user.timeout")*time.Second,
		middleware.NewJwtRoundTripper(
			simpleClient.NewMetricsRoundTripper(
				simpleClient.NewLoggingRoundTripper(transport, logger, true),
			),
			logger,
			viper.GetString("external_services.user.api_token"),
		),
	)

	closeFunc := func() {
		sClient.Close()
		logger.Info("external client closed successfully")
	}

	return genclient.NewClient(sClient), closeFunc
}
