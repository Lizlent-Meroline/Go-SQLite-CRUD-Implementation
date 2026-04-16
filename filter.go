package main

func (s *Store) ListProducts(minPrice float64, nameFilter string) ([]Product, error) {
	query := "SELECT id, name, price, quantity FROM products WHERE 1=1"
	args := []any{}

	if minPrice > 0 {
		query += " AND price >= ?"
		args = append(args, minPrice)
	}

	if nameFilter != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+nameFilter+"%")
	}

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
