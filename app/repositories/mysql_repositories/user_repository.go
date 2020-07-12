package mysql_repositories

import (
	"github.com/jedi4z/minesweeper-api/app/models"
	"github.com/jedi4z/minesweeper-api/app/repositories"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepositoryInterface {
	return UserRepository{
		DB: db,
	}
}

func (r UserRepository) Insert(u *models.User) error {
	return r.DB.Create(&u).Error
}

func (r UserRepository) IsUnique(u *models.User) (bool, error) {
	type exists struct {
		Result bool
	}

	var e exists
	if err := r.DB.
		Raw("SELECT EXISTS(SELECT * FROM users WHERE email = ?) AS result", u.Email).
		Scan(&e).Error; err != nil {
		return e.Result, err
	}

	return !e.Result, nil
}

func (r UserRepository) FindOne(id string) (*models.User, error) {
	user := &models.User{}
	if err := r.DB.Where("ID = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) FindOneByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := r.DB.Where("Email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
