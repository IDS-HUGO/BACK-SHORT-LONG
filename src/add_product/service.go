package add_product

import "sync"

type Service struct {
	Productos        []Producto
	ProductoAgregado bool
	mu               sync.Mutex
}

func NewService() *Service {
	return &Service{Productos: []Producto{}, ProductoAgregado: false}
}

func (s *Service) AddProduct(p Producto) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Productos = append(s.Productos, p)
	s.ProductoAgregado = true
}

func (s *Service) GetAllProducts() []Producto {
	s.mu.Lock()
	defer s.mu.Unlock()
	productos := make([]Producto, len(s.Productos))
	copy(productos, s.Productos)
	return productos
}

func (s *Service) IsProductoAgregado() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ProductoAgregado
}

func (s *Service) ResetProductoAgregado() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ProductoAgregado = false
}
