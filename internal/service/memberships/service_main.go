package memberships

import (
	"github.com/kautsarhasby/katalog-musik/internal/configs"
	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
)

//go:generate mockgen -source=service_main.go -destination=service_main_mock_test.go -package=memberships
type repository interface {
	GetUser(email, username string, id uint) (*memberships.User, error)
	CreateUser(model memberships.User) error
}

type service struct {
	cfg        *configs.Config
	repository repository
}

func NewService(cfg *configs.Config, repository repository) *service {
	return &service{
		cfg:        cfg,
		repository: repository,
	}
}
