package repositories

import "gorm.io/gorm"

// Repositories holds all repository instances
type Repositories struct {
	DB               *gorm.DB
	User             UserRepository
	VerificationCode VerificationCodeRepository
	// Add more repositories here as needed
	// Product ProductRepository
	// Order OrderRepository
}

// NewRepositories creates a new Repositories instance with all repositories initialized
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		DB:               db,
		User:             NewUserRepository(db),
		VerificationCode: NewVerificationCodeRepository(db),
		// Initialize more repositories here
		// Product: NewProductRepository(db),
		// Order: NewOrderRepository(db),
	}
}

// WithTx creates a new Repositories instance with transaction-aware repositories
// Use this when you need to perform multiple operations in a single transaction
func (r *Repositories) WithTx(tx *gorm.DB) *Repositories {
	return &Repositories{
		DB:               tx,
		User:             NewUserRepository(tx),
		VerificationCode: NewVerificationCodeRepository(tx),
		// Initialize more repositories with transaction here
		// Product: NewProductRepository(tx),
		// Order: NewOrderRepository(tx),
	}
}
