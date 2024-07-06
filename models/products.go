package models

import "alura-store/db"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func SearchAllProducts() []Product {
	db := db.ConnectDB()
	selectAllProducts, err := db.Query("select * from products order by id desc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDB()

	insertDataDB, err := db.Prepare("insert into products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDataDB.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteProduct, err := db.Prepare("delete from products where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func ProductById(id string) Product {
	db := db.ConnectDB()

	productDB, err := db.Query("select * from products where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for productDB.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productDB.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}

	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id, quantity int, name, description string, price float64) {
	db := db.ConnectDB()
	UpdateProcut, err := db.Prepare("update products set name = $1, description = $2, price = $3, quantity = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}

	UpdateProcut.Exec(name, description, price, quantity, id)
	defer db.Close()
}
