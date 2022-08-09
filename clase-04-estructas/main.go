package main

import (
	"errors"
	"fmt"
)

var Products []Product

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	fmt.Println(Products)
}

func getById(id int) (Product, error) {
	for _, v := range Products {
		if v.ID == id {
			return v, nil
		}
	}
	return Product{}, errors.New("product Id doesn't exist")

	/*Product{
		ID:          0,
		Name:        "None",
		Price:       0,
		Description: "None",
		Category:    "None",
	}*/
}

func main() {
	prodcunt1 := Product{
		ID:          25,
		Name:        "Banana",
		Price:       190.00,
		Description: "Juicy banana",
		Category:    "Fruit",
	}
	prodcunt2 := Product{
		ID:          2,
		Name:        "Apple",
		Price:       250.00,
		Description: "Juicy red apple",
		Category:    "Fruit",
	}

	prodcunt1.Save()
	prodcunt2.Save()
	prodcunt2.GetAll()

	var success, fail = getById(25)
	if fail == nil {
		fmt.Println(success)
	} else {
		fmt.Println(fail)
	}

}
