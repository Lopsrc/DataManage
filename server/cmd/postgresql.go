package main

func insertRecordToWorkspace(record Workspace) error{
	stmt, err := db.Prepare("INSERT INTO maloesareevo (work_date, price, time_work, penalty) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для вставки записи
	_, err = stmt.Exec(record.ID, record.Date, record.Price, record.TimeWork, record.Penalty)
	if err != nil {
		return err
	}

	return nil
}
func insertRecordToListPrice(record ListPrices) error{
	stmt, err := db.Prepare("INSERT INTO list_of_price (date_change, price_day) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для вставки записи
	_, err = stmt.Exec(record.ID, record.Date, record.Price)
	if err != nil {
		return err
	}

	return nil
}
func insertRecordToListPayment(record ListPayments) error{
	stmt, err := db.Prepare("INSERT INTO list_of_payments (date_change, price_day) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для вставки записи
	_, err = stmt.Exec(record.ID, record.Date, record.PriceDay)
	if err != nil {
		return err
	}

	return nil
}
func updateRecordInWorkspace(record Workspace) error {
	// Подготовка SQL-запроса для обновления записи
	stmt, err := db.Prepare("UPDATE maloesareevo SET work_date = $2, price = $3, time_work = $4, penalty = $5 WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для обновления записи
	_, err = stmt.Exec(record.ID, record.Date, record.Price, record.TimeWork, record.Penalty)
	if err != nil {
		return err
	}

	return nil
}
func updateRecordInListPrice(record ListPrices) error {
	// Подготовка SQL-запроса для обновления записи
	stmt, err := db.Prepare("UPDATE list_of_price SET date_change = $2, price_day = $3 WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для обновления записи
	_, err = stmt.Exec(record.ID, record.Date, record.Price)
	if err != nil {
		return err
	}

	return nil
}
func updateRecordInListPayment(record ListPayments) error {
	// Подготовка SQL-запроса для обновления записи
	stmt, err := db.Prepare("UPDATE list_of_payments SET payment_date = $2, price = $3 WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для обновления записи
	_, err = stmt.Exec(record.ID, record.Date, record.PriceDay)
	if err != nil {
		return err
	}

	return nil
}
func deleteRecordInWorkspace(recordID int) error {
	// Подготовка SQL-запроса для удаления записи
	stmt, err := db.Prepare("DELETE FROM maloesareevo WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для удаления записи
	_, err = stmt.Exec(recordID)
	if err != nil {
		return err
	}

	return nil
}

func deleteRecordInListPrice(recordID int) error {
	// Подготовка SQL-запроса для удаления записи
	stmt, err := db.Prepare("DELETE FROM list_of_price WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для удаления записи
	_, err = stmt.Exec(recordID)
	if err != nil {
		return err
	}

	return nil
}

func deleteRecordInListPayment(recordID int) error {
	// Подготовка SQL-запроса для удаления записи
	stmt, err := db.Prepare("DELETE FROM list_of_payments WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для удаления записи
	_, err = stmt.Exec(recordID)
	if err != nil {
		return err
	}

	return nil
}



