package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type Store struct {
	DB *sql.DB
}

// Creating
func (s *Store) CreateProduct(p *Product) (int64, error) {
	stmt, err := s.DB.Prepare("INSERT INTO products(name, price, quantity) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(p.Name, p.Price, p.Quantity)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// Reading
func (s *Store) GetProduct(id int) (*Product, error) {
	row := s.DB.QueryRow("SELECT id, name, price, quantity FROM products WHERE id = ?", id)

	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}

	return &p, nil
}

// Updating
func (s *Store) UpdateProduct(p *Product) error {
	stmt, err := s.DB.Prepare("UPDATE products SET name=?, price=?, quantity=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(p.Name, p.Price, p.Quantity, p.ID)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no product updated")
	}

	return nil
}

// Deleting
func (s *Store) DeleteProduct(id int) error {
	res, err := s.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}
