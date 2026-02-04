package repositories

import (
	"context"
	"errors"

	"github.com/huda7077/gin-and-gorm-boilerplate/models"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(ctx context.Context, search, orderBy, sort string, offset, limit int) ([]models.User, int64)
	FindById(ctx context.Context, id int) (models.User, error)
	FindByEmail(ctx context.Context, email string) (models.User, error)
	Create(ctx context.Context, user models.User) (models.User, error)
	Update(ctx context.Context, id int, updatedUser models.User) (models.User, error)
	Delete(ctx context.Context, id int) error
}

type userRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) FindAll(ctx context.Context, search, orderBy, sort string, offset, limit int) ([]models.User, int64) {
	var users []models.User
	var total int64

	query := r.DB.WithContext(ctx).Model(&models.User{})

	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return []models.User{}, 0
	}

	if orderBy != "" {
		if sort != "asc" && sort != "desc" {
			sort = "desc"
		}
		query = query.Order(orderBy + " " + sort)
	}

	if err := query.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return []models.User{}, 0
	}

	return users, total
}

func (r *userRepositoryImpl) FindById(ctx context.Context, id int) (models.User, error) {
	var user models.User
	if err := r.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, exceptions.NotFoundError{Message: "user not found"}
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, exceptions.NotFoundError{Message: "email not found"}
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepositoryImpl) Create(ctx context.Context, user models.User) (models.User, error) {
	err := r.DB.WithContext(ctx).Create(&user).Error
	return user, err
}

func (r *userRepositoryImpl) Update(ctx context.Context, id int, updatedUser models.User) (models.User, error) {
	var user models.User
	if err := r.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, exceptions.NotFoundError{Message: "user not found"}
		}
		return models.User{}, err
	}

	if err := r.DB.WithContext(ctx).Model(&user).Updates(updatedUser).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id int) error {
	var user models.User

	if err := r.DB.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return exceptions.NotFoundError{Message: "user not found"}
		}
		return err
	}

	if err := r.DB.WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
