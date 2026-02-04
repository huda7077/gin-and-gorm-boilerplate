package models

import "time"

type VerificationCode struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	Code      string    `gorm:"size:6;not null" json:"code"`
	Purpose   string    `gorm:"size:50;not null" json:"purpose"` // e.g. "EMAIL_VERIFICATION", "RESET_PASSWORD"
	ExpiredAt time.Time `json:"expiredAt"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (VerificationCode) TableName() string {
	return "verification_codes"
}
