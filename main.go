package main

import (
	"api-productos/src/add_product"
	"api-productos/src/count_products_in_discount"
	"api-productos/src/is_new_product_added"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	addProductService := add_product.NewService()
	isNewProductService := is_new_product_added.NewService(addProductService)
	countDiscountService := count_products_in_discount.NewService(addProductService)

	addProductHandler := add_product.NewHandler(addProductService)
	isNewProductHandler := is_new_product_added.NewHandler(isNewProductService)
	countDiscountHandler := count_products_in_discount.NewHandler(countDiscountService)

	r := mux.NewRouter()
	r.HandleFunc("/addProduct", addProductHandler.AddProductHandler).Methods("POST")
	r.HandleFunc("/isNewProductAdded", isNewProductHandler.IsNewProductAddedHandler).Methods("GET")
	r.HandleFunc("/countProductsInDiscount", countDiscountHandler.CountProductsInDiscountHandler).Methods("GET")
	r.HandleFunc("/getAllProducts", addProductHandler.GetAllProductsHandler).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
