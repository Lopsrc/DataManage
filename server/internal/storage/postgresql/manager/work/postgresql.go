package postgresqlwork

import (
	"context"

	models "server/server/internal/models/work"

	"github.com/jackc/pgx/v4/pgxpool"
)


type Repository struct{
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (rep *Repository) Create(
	ctx context.Context,
	rec models.CreateWork,
)(bool, error){
	return true, nil
}
func (rep *Repository) Update(
	ctx context.Context,
	rec models.UpdateWork,
)(bool, error){
	return true, nil
}
func (rep *Repository) GetAll(
	ctx context.Context,
	rec []*models.Work,
)error{
	return nil
}
func (rep *Repository) GetAllByEmail(
	ctx context.Context,
	rec []*models.Work,
)error{
	return nil
}
func (rep *Repository) Delete(
	ctx context.Context,
	id int64,
)(bool, error){
	return true, nil
}
// func insertRecordToWorkspace(record Workspace) error{
// 	stmt, err := db.Prepare("INSERT INTO workspace (name_workspace, work_date, price, time_work, penalty) VALUES ($1, $2, $3, $4, $5)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Выполнение SQL-запроса для вставки записи
// 	_, err = stmt.Exec(record.ID, record.NameWorkspace, record.Date, record.Price, record.TimeWork, record.Penalty)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func insertRecordToListPrice(record ListPrices) error{
// 	stmt, err := db.Prepare("INSERT INTO list_of_price (date_change, price_day) VALUES ($1, $2)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Выполнение SQL-запроса для вставки записи
// 	_, err = stmt.Exec(record.ID, record.Date, record.Price)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func insertRecordToListPayment(record ListPayments) error{
// 	stmt, err := db.Prepare("INSERT INTO list_of_payments (date_change, price_day) VALUES ($1, $2)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Выполнение SQL-запроса для вставки записи
// 	_, err = stmt.Exec(record.ID, record.Date, record.PriceDay)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func updateRecordInWorkspace(record Workspace) error {
// 	// Подготовка SQL-запроса для обновления записи
// 	stmt, err := db.Prepare("UPDATE workspace SET name_workspace = $2, work_date = $3, price = $4, time_work = $5, penalty = $6 WHERE id = $1")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Выполнение SQL-запроса для обновления записи
// 	_, err = stmt.Exec(record.ID, record.NameWorkspace, record.Date, record.Price, record.TimeWork, record.Penalty)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func updateRecordInListPrice(record ListPrices) error {
// 	// Подготовка SQL-запроса для обновления записи
// 	stmt, err := db.Prepare("UPDATE list_of_price SET date_change = $2, price_day = $3 WHERE id = $1")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Выполнение SQL-запроса для обновления записи
// 	_, err = stmt.Exec(record.ID, record.Date, record.Price)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func updateRecordInListPayment(record ListPayments) error {
// 	// Подготовка SQL-запроса для обновления записи
// 	stmt, err := db.Prepare("UPDATE list_of_payments SET payment_date = $2, price = $3 WHERE id = $1")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Выполнение SQL-запроса для обновления записи
// 	_, err = stmt.Exec(record.ID, record.Date, record.PriceDay)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func deleteRecordInWorkspace(recordID int) error {
// 	// Подготовка SQL-запроса для удаления записи
// 	stmt, err := db.Prepare("DELETE FROM workspace WHERE id = $1")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Выполнение SQL-запроса для удаления записи
// 	_, err = stmt.Exec(recordID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func deleteRecordInListPrice(recordID int) error {
// 	// Подготовка SQL-запроса для удаления записи
// 	stmt, err := db.Prepare("DELETE FROM list_of_price WHERE id = $1")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Выполнение SQL-запроса для удаления записи
// 	_, err = stmt.Exec(recordID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func deleteRecordInListPayment(recordID int) error {
// 	// Подготовка SQL-запроса для удаления записи
// 	stmt, err := db.Prepare("DELETE FROM list_of_payments WHERE id = $1")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Выполнение SQL-запроса для удаления записи
// 	_, err = stmt.Exec(recordID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
