package main

func (s *Store) BulkUpdateStock(updates map[int]int) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("UPDATE products SET quantity = quantity + ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for id, qty := range updates {
		_, err := stmt.Exec(qty, id)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
