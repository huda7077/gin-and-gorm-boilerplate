package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/huda7077/gin-and-gorm-boilerplate/models"
	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
	"gorm.io/gorm"
)

type VerificationCodeRepository interface {
	Create(ctx context.Context, code models.VerificationCode) (models.VerificationCode, error)
	FindByUserAndPurpose(ctx context.Context, userID uint, purpose string) (models.VerificationCode, error)
	FindValidCode(ctx context.Context, userID uint, code, purpose string) (models.VerificationCode, error)
	DeleteByUser(ctx context.Context, userID uint, purpose string) error
	DeleteExpired(ctx context.Context) error
}

type verificationCodeRepositoryImpl struct {
	db *gorm.DB
}

func NewVerificationCodeRepository(db *gorm.DB) VerificationCodeRepository {
	return &verificationCodeRepositoryImpl{db: db}
}

func (r *verificationCodeRepositoryImpl) Create(ctx context.Context, code models.VerificationCode) (models.VerificationCode, error) {
	if err := r.db.WithContext(ctx).Create(&code).Error; err != nil {
		return models.VerificationCode{}, err
	}
	return code, nil
}

func (r *verificationCodeRepositoryImpl) FindByUserAndPurpose(ctx context.Context, userID uint, purpose string) (models.VerificationCode, error) {
	var code models.VerificationCode
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND purpose = ?", userID, purpose).
		Order("created_at DESC").
		First(&code).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.VerificationCode{}, exceptions.NewNotFoundError("verification code not found")
		}
		return models.VerificationCode{}, err
	}

	return code, nil
}

func (r *verificationCodeRepositoryImpl) FindValidCode(ctx context.Context, userID uint, code, purpose string) (models.VerificationCode, error) {
	var verificationCode models.VerificationCode
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND code = ? AND purpose = ? AND expired_at > ?", userID, code, purpose, time.Now()).
		First(&verificationCode).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.VerificationCode{}, exceptions.NewNotFoundError("invalid or expired verification code")
		}
		return models.VerificationCode{}, err
	}

	return verificationCode, nil
}

func (r *verificationCodeRepositoryImpl) DeleteByUser(ctx context.Context, userID uint, purpose string) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND purpose = ?", userID, purpose).
		Delete(&models.VerificationCode{}).Error
}

func (r *verificationCodeRepositoryImpl) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).
		Where("expired_at < ?", time.Now()).
		Delete(&models.VerificationCode{}).Error
}
