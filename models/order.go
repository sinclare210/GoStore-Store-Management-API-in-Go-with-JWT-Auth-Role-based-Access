package models

import (
	"errors"

	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/db"
)

type Order struct {
	Id         int64
	User_Id    int64
	Product_Id int64
}

func (order *Order) CreateOrder() error {
	query := `
	INSERT INTO orders(User_Id,Product_Id)
	VALUES(?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.User_Id, order.Product_Id)
	if err != nil {
		return errors.New("failed to create order")
	}

	return nil

}
