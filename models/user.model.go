package models

import "time"

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

type User struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Name       string     `gorm:"size:255;not null" json:"name"`
	Email      string     `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password   string     `gorm:"size:255;not null" json:"-"`
	Image      *string    `gorm:"size:255;" json:"image"`
	Role       Role       `gorm:"type:varchar(20);not null;default:'USER'" json:"role"`
	VerifiedAt *time.Time `json:"verifiedAt"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`

	VerificationCodes []VerificationCode `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"tokens,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}
