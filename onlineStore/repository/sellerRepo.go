package repository

import (
	"Test_task1/onlineStore"
	"database/sql"
	"fmt"
	"strings"
)

type SellerRepo struct {
	db *sql.DB
}

func NewSellerRepo(db *sql.DB) *SellerRepo {
	return &SellerRepo{db: db}
}

func (r *SellerRepo) Create(seller onlineStore.Seller) (int, error) {
	query := `INSERT INTO sellers (name, phone) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(query, seller.Name, seller.Phone).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create seller: %w", err)
	}
	return id, nil
}

func (r *SellerRepo) GetAll() ([]onlineStore.Seller, error) {
	query := `SELECT id, name, phone FROM sellers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get sellers: %w", err)
	}
	defer rows.Close()

	var sellers []onlineStore.Seller
	for rows.Next() {
		var seller onlineStore.Seller
		if err := rows.Scan(&seller.ID, &seller.Name, &seller.Phone); err != nil {
			return nil, fmt.Errorf("failed to scan seller: %w", err)
		}
		sellers = append(sellers, seller)
	}

	return sellers, nil
}

func (r *SellerRepo) GetById(id int) (onlineStore.Seller, error) {
	query := `SELECT id, name, phone FROM sellers WHERE id = $1`
	var seller onlineStore.Seller
	err := r.db.QueryRow(query, id).Scan(&seller.ID, &seller.Name, &seller.Phone)
	if err != nil {
		return seller, fmt.Errorf("failed to get seller by id: %w", err)
	}

	return seller, nil
}

func (r *SellerRepo) Delete(id int) error {
	query := `DELETE FROM sellers WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete seller: %w", err)
	}
	return nil
}

func (r *SellerRepo) Update(id int, seller onlineStore.UpdateSellerRequest) (onlineStore.Seller, error) {
	setQuery, args := buildUpdateQuery(seller)

	if len(setQuery) == 0 {
		return onlineStore.Seller{}, fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf(`UPDATE sellers SET %s WHERE id = $%d RETURNING id, name, phone`, setQuery, len(args)+1)
	args = append(args, id)

	row := r.db.QueryRow(query, args...)

	var updatedSeller onlineStore.Seller
	if err := row.Scan(&updatedSeller.ID, &updatedSeller.Name, &updatedSeller.Phone); err != nil {
		return onlineStore.Seller{}, fmt.Errorf("failed to update seller: %w", err)
	}

	return updatedSeller, nil
}

func buildUpdateQuery(seller onlineStore.UpdateSellerRequest) (string, []interface{}) {
	var setValues []string
	var args []interface{}
	argID := 1

	if seller.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name = $%d", argID))
		args = append(args, seller.Name)
		argID++
	}

	if seller.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone = $%d", argID))
		args = append(args, seller.Phone)
		argID++
	}

	return strings.Join(setValues, ", "), args
}
