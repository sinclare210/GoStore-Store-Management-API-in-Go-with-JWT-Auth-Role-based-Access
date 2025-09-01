package models

import (
	"errors"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/db"
)

type Product struct {
	Id          int64
	Name        string
	Description string
	Price       float64
	Quantity    int64
	User_Id     int64
}

func (product *Product) CreateProducts() error {
	query := `
	INSERT INTO products(Name,Description,Price,Quantity,User_Id)
	VALUES(?,?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("invalid statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Description, product.Price, product.Quantity, product.User_Id)
	if err != nil {
		return errors.New("invalid inputs")
	}
	return nil
}

func GetProducts() ([]Product, error) {
	query := `
	SELECT * FROM products
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New("invalid query")
	}
	var products []Product
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.User_Id)
		if err != nil {
			return nil, errors.New("invalid output")
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProductById(Id int64) (Product, error) {
	query := `
	SELECT * FROM products WHERE Id = ?
	`

	rows := db.DB.QueryRow(query, Id)
	var product Product

	err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.User_Id)
	if err != nil {
		return Product{}, errors.New("invalid output")
	}

	return product, nil
}
