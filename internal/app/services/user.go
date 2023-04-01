package services

import (
	"hmdp/internal/domain"
)

type UserService struct {
	users domain.UserRepository
}

type UserConfiguration func(os *UserService) error

// NewUserService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewUserService(cfgs ...UserConfiguration) (*UserService, error) {
	// Create the user service
	os := &UserService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func (s *UserService) SendCode() {

}

func (s *UserService) CreateUser(phone string) {
	//s.users.CreateUser(phone)
}
