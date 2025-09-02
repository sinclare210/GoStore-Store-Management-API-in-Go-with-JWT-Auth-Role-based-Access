package models

import (
	"errors"

	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/db"
)

type Order struct {
	Id         int64
	User_Id    int64
	Product_Id int64
	Product_Name string 
	Product_Price float64 
}

func (order *Order) CreateOrder() error {
	query := `
	INSERT INTO orders(User_Id,Product_Id,Product_Name,Product_Price)
	VALUES(?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("failed to prepare statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(order.User_Id, order.Product_Id,order.Product_Name,order.Product_Price)
	if err != nil {
		return errors.New("failed to create order")
	}

	return nil

}

func GetOrdersForUser(id int64) ([]Order,error){
	query := `
	SELECT * FROM orders WHERE User_Id = ?
	`

	rows,err := db.DB.Query(query,id)

	if err != nil {
		return nil, errors.New("invalid query")
	}
	var orders []Order
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.Id, &order.User_Id, &order.Product_Id,&order.Product_Name,&order.Product_Price)
		if err != nil {
			return nil, errors.New("invalid output")
		}
		orders = append(orders, order)
	}
	return orders, nil

}
