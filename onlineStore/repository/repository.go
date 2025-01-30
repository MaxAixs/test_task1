package repository

import (
	"Test_task1/onlineStore"
	"database/sql"
)

type Seller interface {
	Create(seller onlineStore.Seller) (int, error)
	GetById(id int) (onlineStore.Seller, error)
	Delete(id int) error
	GetAll() ([]onlineStore.Seller, error)
	Update(id int, seller onlineStore.UpdateSellerRequest) (onlineStore.Seller, error)
}

type Repository struct {
	Seller
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Seller: NewSellerRepo(db),
	}
}
