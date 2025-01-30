package service

import (
	"Test_task1/onlineStore"
	"Test_task1/onlineStore/repository"
)

type Seller interface {
	CreateSeller(seller onlineStore.Seller) (int, error)
	GetSellerById(id int) (onlineStore.Seller, error)
	DeleteSellerById(id int) error
	GetAllSellers() ([]onlineStore.Seller, error)
	UpdateSellerById(id int, updateData onlineStore.UpdateSellerRequest) (onlineStore.Seller, error)
}
type Service struct {
	Seller
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Seller: NewSellerService(repo.Seller),
	}
}
