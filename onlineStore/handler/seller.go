package handler

import (
	"Test_task1/onlineStore"
	"net/http"
)

func (h *Handler) CreateSeller(w http.ResponseWriter, r *http.Request) {
	if !CheckContentType(w, r) {
		return
	}

	var seller onlineStore.Seller

	if !ParseJSONBody(w, r, seller) {
		return
	}

	if err := ValidateSeller(seller.Name, seller.Phone); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	IdSeller, err := h.services.Seller.CreateSeller(seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newSuccessResponse(w, map[string]interface{}{"id": IdSeller})

}

func (h *Handler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	if !CheckContentType(w, r) {
		return
	}

	sellerId, err := getSellerID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var seller onlineStore.UpdateSellerRequest
	if !ParseJSONBody(w, r, &seller) {
		return
	}

	updatedSeller, err := h.services.Seller.UpdateSellerById(sellerId, seller)
	if err != nil {
		http.Error(w, "Failed to update seller", http.StatusInternalServerError)
		return
	}

	newSuccessResponse(w, map[string]interface{}{
		"UpdateSeller": updatedSeller,
	})
}

func (h *Handler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	if !CheckContentType(w, r) {
		return
	}

	sellerId, err := getSellerID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.services.Seller.DeleteSellerById(sellerId)
	if err != nil {
		http.Error(w, "Failed to delete seller", http.StatusInternalServerError)
		return
	}

	newSuccessResponse(w, map[string]interface{}{"data": "seller successfully deleted"})

}

func (h *Handler) GetSellerById(w http.ResponseWriter, r *http.Request) {
	if !CheckContentType(w, r) {
		return
	}

	sellerId, err := getSellerID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	seller, err := h.services.Seller.GetSellerById(sellerId)
	if err != nil {
		http.Error(w, "Seller not found", http.StatusNotFound)
		return
	}

	newSuccessResponse(w, map[string]interface{}{"id": seller.ID, "name": seller.Name, "phone": seller.Phone})
}

func (h *Handler) GetAllSellers(w http.ResponseWriter, r *http.Request) {
	if !CheckContentType(w, r) {
		return
	}

	sellers, err := h.services.Seller.GetAllSellers()
	if err != nil {
		http.Error(w, "Failed to get sellers", http.StatusInternalServerError)
		return
	}

	newSuccessResponse(w, map[string]interface{}{
		"AllSellers": sellers,
	})
}
