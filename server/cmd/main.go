package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	grpc_ "server/server/internal/app"
	"server/server/internal/config"
	"server/server/pkg/client/postgresql"
	"syscall"
)

const (
	pathConfig = "server/config/local.yaml"
	envLocal   			= "local"
	envDev     			= "dev"
	envProd    			= "prod"
)

func main() {
	
	cfg := config.GetConfig(pathConfig)

	log := setupLogger(cfg.Env)

	postgreSQLClient, err := postgresql.NewClient(context.Background(), 3, cfg.Storage)
	if err!= nil {
        panic(err)
    }

	application := grpc_.New(log, cfg.Listen.Port, postgreSQLClient)

	go func() {
		application.GRPCServer.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("gRPC seever stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
// func createRecordFromWorkspace(w http.ResponseWriter, r *http.Request) {
// 	var record Workspace
// 	err := json.NewDecoder(r.Body).Decode(&record)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	// Проверка валидности токена
// 	if !isValidToken(record.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	err = insertRecordToWorkspace(record)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func createRecordFromListPayment(w http.ResponseWriter, r *http.Request) {
// 	var record ListPayments
// 	err := json.NewDecoder(r.Body).Decode(&record)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Проверка валидности токена
// 	if !isValidToken(record.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	err = insertRecordToListPayment(record)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func updateRecordFromWorkspace(w http.ResponseWriter, r *http.Request) {
// 	var record Workspace
// 	err := json.NewDecoder(r.Body).Decode(&record)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Проверка валидности токена
// 	if !isValidToken(record.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	err = updateRecordInWorkspace(record)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func updateRecordFromListPayments(w http.ResponseWriter, r *http.Request) {
// 	var record ListPayments
// 	err := json.NewDecoder(r.Body).Decode(&record)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	// Проверка валидности токена
// 	if !isValidToken(record.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	err = updateRecordInListPayment(record)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func deleteRecordFromWorkspace(w http.ResponseWriter, r *http.Request) {
// 	var record Workspace
// 	err := json.NewDecoder(r.Body).Decode(&record)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Проверка валидности токена
// 	if !isValidToken(record.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	err = deleteRecordInWorkspace(record.ID)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func deleteRecordFromListPayments(w http.ResponseWriter, r *http.Request) {
// 	var record ListPayments
// 	err := json.NewDecoder(r.Body).Decode(&record)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Проверка валидности токена
// 	if !isValidToken(record.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	err = deleteRecordInListPayment(record.ID)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// // Функция getAllRecords возвращает все записи из таблицы "workspace" базы данных
// func getAllRecords(w http.ResponseWriter, r *http.Request) {
// 	var token Token
// 	err := json.NewDecoder(r.Body).Decode(&token)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	// Проверка валидности токена
// 	if !isValidToken(token.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	// Выполнение SQL-запроса для получения всех записей из таблицы "workspace"
// 	rows, err := db.Query("SELECT * FROM workspace")
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Println(err)
// 		return
// 	}
// 	defer rows.Close()

// 	// Создание среза для хранения записей
// 	var records []Workspace

// 	// Итерация по результатам запроса
// 	for rows.Next() {
// 		var record Workspace
// 		err := rows.Scan(&record.ID, &record.NameWorkspace, &record.Date, &record.Price, &record.TimeWork, &record.Penalty)
// 		if err != nil {
// 			print(err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			log.Println(err)
// 			return
// 		}
// 		records = append(records, record)
// 	}

// 	// Проверка ошибок при итерации
// 	err = rows.Err()
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Println(err)
// 		return
// 	}

// 	// Преобразование записей в JSON
// 	jsonData, err := json.Marshal(records)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Println(err)
// 		return
// 	}

// 	// Отправка JSON-ответа клиенту
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonData)
// }

// func getAllRecordsFromListPayments(w http.ResponseWriter, r *http.Request) {
// 	var token Token
// 	err := json.NewDecoder(r.Body).Decode(&token)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	// Проверка валидности токена
// 	if !isValidToken(token.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	// Получение всех записей из таблицы "list_of_payments"
// 	rows, err := db.Query("SELECT * FROM list_of_payments")
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	// Создание среза для хранения всех записей
// 	var entries []ListPayments

// 	// Итерация по результатам запроса и добавление записей в срез
// 	for rows.Next() {
// 		var entry ListPayments
// 		err := rows.Scan(
// 			&entry.ID,
// 			&entry.NameWorkspace,
// 			&entry.Date,
// 			&entry.PriceDay,
// 		)
// 		if err != nil {
// 			print(err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		entries = append(entries, entry)
// 	}

// 	// Проверка наличия ошибок при выполнении запроса
// 	if err = rows.Err(); err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	// Кодирование среза записей в формат JSON и отправка клиенту
// 	err = json.NewEncoder(w).Encode(entries)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }
// func isValidToken(token string) bool {
// 	if token!=TOKEN {return false}
// 	return true
// }

// func insertRecordToWorkspace(record Workspace) error {
// 	stmt, err := db.Prepare("INSERT INTO workspace (name_workspace, work_date, price, time_work, penalty) VALUES ($1, $2, $3, $4, $5)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	// Выполнение SQL-запроса для вставки записи
// 	_, err = stmt.Exec(record.NameWorkspace, record.Date, record.Price, record.TimeWork, record.Penalty)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func insertRecordToListPayment(record ListPayments) error {
// 	stmt, err := db.Prepare("INSERT INTO list_of_payments (name_workspace, payment_date, price) VALUES ($1, $2, $3)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	// Выполнение SQL-запроса для вставки записи
// 	_, err = stmt.Exec(record.NameWorkspace, record.Date, record.PriceDay)
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
// func updateRecordInListPayment(record ListPayments) error {
// 	// Подготовка SQL-запроса для обновления записи
// 	stmt, err := db.Prepare("UPDATE list_of_payments SET name_workspace = $2, payment_date = $3, price = $4 WHERE id = $1")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	print("hhhhh : "+record.Date)
// 	print(record.NameWorkspace)
// 	print(record.ID)
// 	print(record.PriceDay)
// 	// Выполнение SQL-запроса для обновления записи
// 	_, err = stmt.Exec(record.ID, record.NameWorkspace, record.Date, record.PriceDay)
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
