package count_products_in_discount

import "api-productos/src/add_product"

type Service struct {
	AddProductService *add_product.Service
}

func NewService(addProductService *add_product.Service) *Service {
	return &Service{AddProductService: addProductService}
}

func (s *Service) CountProductsInDiscount() int {
	count := 0
	for _, p := range s.AddProductService.GetAllProducts() {
		if p.Descuento {
			count++
		}
	}
	return count
}
func (s *Service) IsNewProductAdded() bool {
	return s.AddProductService.IsProductoAgregado()
}
