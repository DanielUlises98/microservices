package data

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func GetProducts() Products {
	return productList
}
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductId(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}
	return productList[i], nil
}
func UpdateProduct(p Product) error {
	ind := findIndexByProductId(p.ID)
	if ind == -1 {
		return ErrProductNotFound
	}
	productList[ind] = &p
	return nil
}

func DeleteProduct(id int) error {
	index := findIndexByProductId(id)
	if index == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:index], productList[index+1:]...)
	return nil
}
func findIndexByProductId(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

var ErrProductNotFound = fmt.Errorf("Product Not Found")

func AddProduct(p Product) {
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)

}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}
