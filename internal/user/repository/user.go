package repository

import (
	"apigo/internal/user/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func MakeUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u *UserRepository) Create(user models.User) (*models.User, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := models.User{}
	if err := u.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindByID(usr models.User) (*models.User, error) {
	user := models.User{}
	if err := u.DB.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
