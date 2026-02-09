package memberships

import (
	"errors"
	"fmt"

	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
	"github.com/kautsarhasby/katalog-musik/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(request memberships.LoginRequest) (string, error) {
	user, err := s.repository.GetUser(request.Email, "", 0)
	fmt.Printf("user EMAIL: %+v", user)
	if err != nil {
		log.Error().Err(err).Msg("Error get user from database")
		return "", err
	}

	if user == nil {
		return "", errors.New("email not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	fmt.Println(err)
	if err != nil {
		return "", errors.New("Email and password invalid")
	}

	accessToken, err := jwt.CreateToken(int64(user.ID), user.Username, s.cfg.Service.SecretKey)
	if err != nil {
		log.Error().Err(err).Msg("Error Create token")
		return "", err
	}

	return accessToken, nil
}
