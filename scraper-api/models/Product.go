package models

import (
	"database/sql"
	"scraper-api/db"
)

type Product struct {
	ID          int64
	Name, Price string
}

func GetAllProducts(limit, page int64) ([]Product, error) {
	query := "SELECT * FROM products LIMIT ? OFFSET ? "

	var rows *sql.Rows
	var err error

	if limit == 0 {
		rows, err = db.DB.Query("SELECT * FROM products")
	} else {
		rows, err = db.DB.Query(query, limit, page*limit)
	}

	// defer rows.Close()

	if err != nil {
		return nil, err
	}

	var products []Product

	for rows.Next() {
		var p Product

		err = rows.Scan(&p.ID, &p.Name, &p.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
