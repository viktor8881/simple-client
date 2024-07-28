package user

import (
	"context"
	"errors"
	simplehttp "github.com/viktor8881/service-utilities/http/client"
	"github.com/viktor8881/service-utilities/http/server"
	"net/http"
	genmodel "simple-client/generated/http/client"
	generated "simple-client/generated/http/server"
)

type UserService struct {
	client client
}

func NewService(client client) *UserService {
	return &UserService{client: client}
}

func (s *UserService) ListUser(ctx context.Context, in *generated.EmptyRequest) (*generated.ProxyResponse, error) {
	dtoOut, err := s.client.ListUserEndpoint(ctx, &genmodel.EmptyRequest{})
	if err != nil {
		var clientNot200Resp *simplehttp.ClientResponseNot200Error
		if errors.As(err, &clientNot200Resp) {
			if clientNot200Resp.ClientResponseCode > http.StatusOK && clientNot200Resp.ClientResponseCode < http.StatusMultipleChoices {
				return nil, &server.CustomError{
					HttpCode:    clientNot200Resp.ClientResponseCode,
					HttpMessage: clientNot200Resp.ClientResponseBody,
					Err:         err,
				}
			}
		}

		return nil, &server.CustomError{
			HttpCode: http.StatusBadGateway,
		}
	}

	resp := generated.ProxyResponse(dtoOut)
	return &resp, nil
}

func (s *UserService) ListUserByEmail(ctx context.Context, in *generated.ListUserByEmailRequest) (*generated.ProxyResponse, error) {
	dtoOut, err := s.client.ListUserByEmailEndpoint(ctx, &genmodel.ListUserByEmailRequest{Email: in.Email})
	if err != nil {
		var clientNot200Resp *simplehttp.ClientResponseNot200Error
		if errors.As(err, &clientNot200Resp) {
			if clientNot200Resp.ClientResponseCode > http.StatusOK && clientNot200Resp.ClientResponseCode < http.StatusMultipleChoices {
				return nil, &server.CustomError{
					HttpCode:    clientNot200Resp.ClientResponseCode,
					HttpMessage: clientNot200Resp.ClientResponseBody,
					Err:         err,
				}
			}
		}

		return nil, &server.CustomError{
			HttpCode: http.StatusBadGateway,
		}
	}

	resp := generated.ProxyResponse(dtoOut)
	return &resp, nil
}

func (s *UserService) GetUser(ctx context.Context, in *generated.GetUserRequest) (*generated.GetUserResponse, error) {
	dtoOut, err := s.client.GetUserEndpoint(ctx, &genmodel.GetUserRequest{ID: in.ID})
	if err != nil {
		var clientNot200Resp *simplehttp.ClientResponseNot200Error
		if errors.As(err, &clientNot200Resp) {
			return nil, &server.ProxyError{
				Code: clientNot200Resp.ClientResponseCode,
				Body: clientNot200Resp.ClientResponseBody,
				Err:  err,
			}
		}

		return nil, &server.CustomError{
			HttpCode: http.StatusBadGateway,
		}
	}

	return &generated.GetUserResponse{ID: dtoOut.ID, Name: dtoOut.Name, Email: dtoOut.Email}, nil
}

func (s *UserService) CreateUser(ctx context.Context, in *generated.CreateUserRequest) (*generated.GetUserResponse, error) {
	dtoOut, err := s.client.CreateUserEndpoint(ctx, &genmodel.CreateUserRequest{Name: in.Name, Email: in.Email})
	if err != nil {
		var clientNot200Resp *simplehttp.ClientResponseNot200Error
		if errors.As(err, &clientNot200Resp) {
			if clientNot200Resp.ClientResponseCode > http.StatusOK && clientNot200Resp.ClientResponseCode < http.StatusMultipleChoices {
				return nil, &server.CustomError{
					HttpCode:    clientNot200Resp.ClientResponseCode,
					HttpMessage: clientNot200Resp.ClientResponseBody,
					Err:         err,
				}
			}
		}

		return nil, &server.CustomError{
			HttpCode: http.StatusBadGateway,
		}
	}

	return сonvertProxyResponseToGetUserResponse(dtoOut)
}

func (s *UserService) UpdateUser(ctx context.Context, in *generated.UpdateUserRequest) (*generated.GetUserResponse, error) {
	dtoOut, err := s.client.UpdateUserEndpoint(ctx, &genmodel.UpdateUserRequest{ID: in.ID, Name: in.Name, Email: in.Email})
	if err != nil {
		var clientNot200Resp *simplehttp.ClientResponseNot200Error
		if errors.As(err, &clientNot200Resp) {
			if clientNot200Resp.ClientResponseCode > http.StatusOK && clientNot200Resp.ClientResponseCode < http.StatusMultipleChoices {
				return nil, &server.CustomError{
					HttpCode:    clientNot200Resp.ClientResponseCode,
					HttpMessage: clientNot200Resp.ClientResponseBody,
					Err:         err,
				}
			}
		}

		return nil, &server.CustomError{
			HttpCode: http.StatusBadGateway,
		}
	}

	return сonvertProxyResponseToGetUserResponse(dtoOut)
}

func (s *UserService) DeleteUser(ctx context.Context, in *generated.DeleteUserRequest) (*generated.EmptyResponse, error) {
	_, err := s.client.DeleteUserEndpoint(ctx, &genmodel.DeleteUserRequest{ID: in.ID})
	if err != nil {
		var clientNot200Resp *simplehttp.ClientResponseNot200Error
		if errors.As(err, &clientNot200Resp) {
			if clientNot200Resp.ClientResponseCode > http.StatusOK && clientNot200Resp.ClientResponseCode < http.StatusMultipleChoices {
				return nil, &server.CustomError{
					HttpCode:    clientNot200Resp.ClientResponseCode,
					HttpMessage: clientNot200Resp.ClientResponseBody,
					Err:         err,
				}
			}
		}

		return nil, &server.CustomError{
			HttpCode: http.StatusBadGateway,
		}
	}

	var dest generated.EmptyResponse
	return &dest, nil
}
