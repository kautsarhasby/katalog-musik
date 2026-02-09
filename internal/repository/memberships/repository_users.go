package memberships

import (
	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
)

func (r *repository) CreateUser(model memberships.User) error {
	return r.db.Create(&model).Error
}

func (r *repository) GetUser(email, username string, id uint) (*memberships.User, error) {
	user := memberships.User{}
	result := r.db.Where("email = ?", email).Or("username = ?", username).Or("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
