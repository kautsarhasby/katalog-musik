package memberships

import (
	"errors"

	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignUp(request memberships.SignUpRequest) error {
	existingUser, err := s.repository.GetUser(request.Email, request.Username, 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("Error get user from database")
		return err
	}

	if existingUser != nil {
		return errors.New("Email or username already exist")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("error hash password")
		return err
	}

	model := memberships.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  string(pass),
		CreatedBy: request.Email,
		UpdateBy:  request.Email,
	}

	return s.repository.CreateUser(model)
}
