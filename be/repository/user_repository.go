// repositories/user_repository.go
package repositories

import (
	"SingleService-Labpro/initializers"
	model "SingleService-Labpro/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByUsername(username string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	db, err := initializers.GetDBInstance()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
