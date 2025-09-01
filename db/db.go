package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

}

func createTables() {
	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
    Id INTEGER PRIMARY KEY AUTOINCREMENT, 
    Name TEXT NOT NULL,
    Description TEXT NOT NULL,
    Price REAL NOT NULL CHECK (Price > 0),
    Quantity INTEGER NOT NULL CHECK (Quantity >= 0),
	User_Id INTEGER,
	FOREIGN KEY(User_Id) REFERENCES users(Id)
);
`
	_, err := DB.Exec(createProductsTable)
	if err != nil {
		panic("Could not create products table")
	}

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
    Name TEXT NOT NULL,
    Email TEXT UNIQUE NOT NULL,
    Password TEXT NOT NULL,
    Role TEXT DEFAULT 'user'
);
`

	_, err = DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createOrdersTable := `
CREATE TABLE IF NOT EXISTS orders (
    Id INTEGER PRIMARY KEY AUTOINCREMENT,
	User_Id INTEGER NOT NULL,
	Product_Id INTEGER NOT NULL,
	FOREIGN KEY(User_Id) REFERENCES users(Id),
	FOREIGN KEY(Product_Id) REFERENCES products(Id)
);
`
	_, err = DB.Exec(createOrdersTable)
	if err != nil {
		panic("Could not create order table")
	}
}
