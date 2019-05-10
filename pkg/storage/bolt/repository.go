package bolt

import (
	"github.com/rhass99/katz/pkg/auth/local"
)

// Storage into a bbolt database
type Storage struct {
	user User
}

// AddUser adds a user to the database
func (s *Storage) AddUser(u *local.SignUpRequest) {}

// GetUser returns a user from database or error if doesnt exist
func (s *Storage) GetUser(u *local.SignInRequest) *local.AuthResponse{
	authResponse := &local.AuthResponse{}
	return authResponse
}