package service

import (
	"Test_task1/onlineStore"
	"Test_task1/onlineStore/repository"
	"errors"
)

type SellerService struct {
	repo repository.Seller
}

func NewSellerService(repo repository.Seller) *SellerService {
	return &SellerService{repo: repo}
}

func (s *SellerService) CreateSeller(seller onlineStore.Seller) (int, error) {
	if !isValidPhone(seller.Phone) {
		return 0, errors.New("invalid phone format")
	}

	return s.repo.Create(seller)
}

func isValidPhone(phone string) bool {
	return len(phone) >= 10
}

func (s *SellerService) GetAllSellers() ([]onlineStore.Seller, error) {
	return s.repo.GetAll()
}

func (s *SellerService) GetSellerById(id int) (onlineStore.Seller, error) {
	return s.repo.GetById(id)
}

func (s *SellerService) UpdateSellerById(id int, seller onlineStore.UpdateSellerRequest) (onlineStore.Seller, error) {
	if seller.Name == nil && seller.Phone == nil {
		return onlineStore.Seller{}, errors.New("update input cannot be empty: at least one field must be provided")
	}
	return s.repo.Update(id, seller)
}

func (s *SellerService) DeleteSellerById(id int) error {
	return s.repo.Delete(id)
}
