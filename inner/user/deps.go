package user

import (
	"context"
	genmodel "simple-client/generated/http/client"
)

type client interface {
	ListUserEndpoint(ctx context.Context, in *genmodel.EmptyRequest) (*genmodel.ProxyResponse, error)
	ListUserByEmailEndpoint(ctx context.Context, in *genmodel.ListUserByEmailRequest) (*genmodel.ProxyResponse, error)
	GetUserEndpoint(ctx context.Context, in *genmodel.GetUserRequest) (*genmodel.GetUserResponse, error)
	CreateUserEndpoint(ctx context.Context, in *genmodel.CreateUserRequest) (*genmodel.ProxyResponse, error)
	UpdateUserEndpoint(ctx context.Context, in *genmodel.UpdateUserRequest) (*genmodel.ProxyResponse, error)
	DeleteUserEndpoint(ctx context.Context, in *genmodel.DeleteUserRequest) (*genmodel.ProxyResponse, error)
}
