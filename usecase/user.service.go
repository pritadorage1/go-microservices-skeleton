package usecase

import (
	context "context"
	fmt "fmt"

	"go_microservices_skeleton/repository"
)

type Service struct {
	Repo *repository.UserRepository
}

func (s *Service) CreateUser(ctx context.Context, req *UserCreateRequest, res *UserCreateResponse) error {
	fmt.Println("user created", req)
	return nil
}
