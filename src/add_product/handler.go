package add_product

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

func (h *Handler) AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var producto Producto

	if err := json.NewDecoder(r.Body).Decode(&producto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	h.Service.AddProduct(producto)
	w.WriteHeader(http.StatusCreated)
}
func (h *Handler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	productos := h.Service.GetAllProducts()

	json.NewEncoder(w).Encode(productos)
}
