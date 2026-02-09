package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"type:varchar(100)" json:"name"`
	Price     int            `gorm:"check:price >= 0 AND price <= 1000000" json:"price"`
	Stock     *int           `gorm:"check:stock >= 0 AND stock <= 100" json:"stock"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (u *Product) TableName() string {
	return "products"
}
