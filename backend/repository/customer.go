package repository

import (
	"database/sql"

	deremsmodels "der-ems/models/der-ems"
)

// CustomerRepository godoc
type CustomerRepository interface {
	GetCustomerByCustomerID(customerID int) (*deremsmodels.Customer, error)
}

type defaultCustomerRepository struct {
	db *sql.DB
}

// NewCustomerRepository godoc
func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &defaultCustomerRepository{db}
}

// GetCustomerByCustomerID godoc
func (repo defaultCustomerRepository) GetCustomerByCustomerID(customerID int) (*deremsmodels.Customer, error) {
	return deremsmodels.FindCustomer(repo.db, customerID)
}
