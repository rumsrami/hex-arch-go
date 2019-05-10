package local

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

// Repository defines the storage service
type Repository interface {
	AddUser(*SignUpRequest)
	GetUser(*SignInRequest) *AuthResponse
}


// Service defines the local auth server
type Service struct{
	userRepo Repository
}

// SignUp implements Signup request
func (s *Service) SignUp(ctx context.Context, request *SignUpRequest) (*AuthResponse, error) {
	return &AuthResponse{
		User: &User{
			Username: request.Username,
			Email:    request.Email,
		},
		Token: "1234",
		Error: &ErrorMessage{
			Message: "",
			Code:    0,
		},
	}, nil
}

// SignIn implements Signin request
func (s *Service) SignIn(ctx context.Context, request *SignInRequest) (*AuthResponse, error) {
	return &AuthResponse{
		User: &User{
			Username: "ramih",
			Email:    "rami@rami.com",
		},
		Token: "1234",
		Error: &ErrorMessage{
			Message: "",
			Code:    0,
		},
	}, nil
}

// NewGRPCService returns a new Instantiated server
func NewGRPCService(lis net.Listener) {
	grpcServer := grpc.NewServer()
	RegisterAuthServer(grpcServer, &Service{})
	grpcServer.Serve(lis)
}


