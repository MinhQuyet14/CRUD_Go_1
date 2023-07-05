package main

import (
	"go-web/config"
	"go-web/controllers/categorycontroller"
	"go-web/controllers/homecontrollers"
	"go-web/controllers/productcontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	// HOME PAGE
	http.HandleFunc("/", homecontrollers.Welcome)

	// CATEGORIES
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)
	//PRODUCTS
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)
	//RUN SERVER
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)

}
