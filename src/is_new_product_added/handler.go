package count_products_in_discount

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CountProductsInDiscountHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]int{"discounted_products": h.Service.CountProductsInDiscount()})
}

func (h *Handler) IsNewProductAddedHandler(w http.ResponseWriter, r *http.Request) {

	if h.Service.AddProductService.IsProductoAgregado() {

		h.Service.AddProductService.ResetProductoAgregado()
		json.NewEncoder(w).Encode(map[string]bool{"new_product": true})
	} else {
		json.NewEncoder(w).Encode(map[string]bool{"new_product": false})
	}
}
