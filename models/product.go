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

	stmt.Exec(product.Name, product.Description, product.Price, product.Quantity, product.User_Id)
	if err != nil {
		return errors.New("invalid inputs")
	}
	return nil
}
