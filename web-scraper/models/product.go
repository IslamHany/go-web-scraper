package models

import "web-scraper/db"

type Product struct {
	Name, Price string
}

func (p *Product) Save() error {
	query := `INSERT INTO products(name, price)
		VALUES(?, ?)
	`

	_, err := db.DB.Exec(query, p.Name, p.Price)

	return err
}
