package local

import (
	"context"

	"google.golang.org/grpc"
)

// Server defines the local auth server
type Server struct{}

// SignUp implements Signup request
func (s *Server) SignUp(ctx context.Context, request *SignUpRequest) (*AuthResponse, error) {
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

// SignIn implements Signin request
func (s *Server) SignIn(ctx context.Context, request *SignInRequest) (*AuthResponse, error) {
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

// NewService returns a new Instantiated server
func NewService() *grpc.Server {
	grpcServer := grpc.NewServer()
	RegisterAuthServer(grpcServer, &Server{})
	return grpcServer
}
