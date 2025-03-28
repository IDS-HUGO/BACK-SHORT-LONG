package is_new_product_added

import "api-productos/src/add_product"

type Service struct {
	AddProductService *add_product.Service
}

func NewService(addProductService *add_product.Service) *Service {
	return &Service{AddProductService: addProductService}
}

func (s *Service) IsNewProductAdded() bool {
	return len(s.AddProductService.GetAllProducts()) > 0
}
