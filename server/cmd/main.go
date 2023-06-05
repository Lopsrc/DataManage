package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)
//delete print(err.Error())
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "postgres"
)

type Entry struct {
	ID       int    `json:"id"`
	ID_table int 	`json:"id_table"`
	Token    string `json:"token"`
	Date     string `json:"date"`
	Price    int    `json:"price"`
	TimeWork int    `json:"time_work"`
	Penalty  int    `json:"penalty"`
	Payments int    `json:"payments"`
	
}

var db *sql.DB

func main() {
	
	// Установка соединения с базой данных
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		print(err.Error())
		log.Fatal(err)
	}
	defer db.Close()

	// Инициализация маршрутизатора
	router := mux.NewRouter()

	// Установка обработчиков запросов
	router.HandleFunc("/entry", createRecordFromWorkspace).Methods("POST")
	router.HandleFunc("/price", createRecordFromListPrice).Methods("POST")
	router.HandleFunc("/payments", createRecordFromListPayment).Methods("POST")

	router.HandleFunc("/entry", updateRecordFromWorkspace).Methods("PUT")
	router.HandleFunc("/price", updateRecordFromListPrice).Methods("PUT")
	router.HandleFunc("/payments", updateRecordFromListPayments).Methods("PUT")

	router.HandleFunc("/entry", deleteRecordFromWorkspace).Methods("DELETE")
	router.HandleFunc("/price", deleteRecordFromListPrice).Methods("DELETE")
	router.HandleFunc("/payments", deleteRecordFromListPayments).Methods("DELETE")

	router.HandleFunc("/entry", getAllRecords).Methods("GET") // Новый обработчик для получения всех записей
	router.HandleFunc("/price", getAllRecordsFromListPrice).Methods("GET")       // Обработчик для получения всех записей из таблицы "list_of_price"
	router.HandleFunc("/payments", getAllRecordsFromListPayments).Methods("GET") // Обработчик для получения всех записей из таблицы "list_of_payments"
	// Запуск сервера на порту 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}


// func createEntry(w http.ResponseWriter, r *http.Request) {
// 	// Чтение и декодирование JSON из запроса
	
// 	var entry Entry
// 	err := json.NewDecoder(r.Body).Decode(&entry)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Проверка валидности токена
// 	if !isValidToken(entry.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
	
// 	switch entry.ID_table {
// 		case 1: // Добавление данных в бд таблица 1
// 			err = insertEntryToDB(entry)
// 			if err != nil {
// 				print(err.Error())
// 				w.WriteHeader(http.StatusInternalServerError)
// 				return
// 			}
// 		case 2: // Добавление данных в бд таблица 2
// 		err = insertEntryToDBtoListPrices(entry)
// 		if err != nil {
// 			print(err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 			}
// 		case 3: // Добавление данных в бд таблица 3
// 		err = insertEntryToDBtoListPayments(entry)
// 		if err != nil {
// 			print(err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 			}
// 		}


// 	w.WriteHeader(http.StatusOK)
// }

// func insertEntryToDB(entry Entry) error {
// 	print("begin")
// 	stmt, err := db.Prepare("INSERT INTO maloesareevo (work_date, price, time_work, penalty) VALUES ($1, $2, $3, $4)")
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}
// 	defer stmt.Close()
// 	print("such")
// 	_, err = stmt.Exec(entry.Date, entry.Price, entry.TimeWork, entry.Penalty)
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}

// 	return nil
// }

// func insertEntryToDBtoListPrices(entry Entry) error {
// 	print("begin")
// 	stmt, err := db.Prepare("INSERT INTO list_of_price (date_change, price_day) VALUES ($1, $2)")
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}
// 	defer stmt.Close()
// 	print("such")
// 	_, err = stmt.Exec(entry.Date, entry.Price)
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}

// 	return nil
// }

// func insertEntryToDBtoListPayments(entry Entry) error {
// 	print("begin")
// 	stmt, err := db.Prepare("INSERT INTO list_of_payments (payment_date, price) VALUES ($1, $2)")
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}
// 	defer stmt.Close()
// 	print("such")
// 	_, err = stmt.Exec(entry.Date, entry.Price)
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}

// 	return nil
// }

// func updateEntry(w http.ResponseWriter, r *http.Request) {
// 	// Чтение и декодирование JSON из запроса
// 	var entry Entry
// 	err := json.NewDecoder(r.Body).Decode(&entry)
// 	if err != nil {
// 		print(err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Проверка валидности токена
// 	if !isValidToken(entry.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}

// 	// Обновление данных в базе данных
// 	switch entry.ID_table {
// 		case 1: // Обновление данных в бд таблица 1
// 			err = updateEntryInDB(entry)
// 			if err != nil {
// 				w.WriteHeader(http.StatusInternalServerError)
// 				print(err.Error())
// 				return
// 			}
// 		case 2: // Обновление данных в бд таблица 2
// 		err = updateEntryInDBtoListPrices(entry)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			print(err.Error())
// 			return
// 			}
// 		case 3: // Обновление данных в бд таблица 3
// 		err = updateEntryInDBtoListPayments(entry)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			print(err.Error())
// 			return
// 			}
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func updateEntryInDB(entry Entry) error {
// 	stmt, err := db.Prepare("UPDATE maloesareevo SET work_date = $1, price = $2, time_work = $3, penalty = $4 WHERE id = $5")
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(entry.Date, entry.Price, entry.TimeWork, entry.Penalty, entry.ID)
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}

// 	return nil
// }

// func updateEntryInDBtoListPrices(entry Entry) error {
// 	stmt, err := db.Prepare("UPDATE list_of_price SET date_change = $1, price_day = $2 WHERE id = $3")
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(entry.Date, entry.Price, entry.ID)
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}

// 	return nil
// }

// func updateEntryInDBtoListPayments(entry Entry) error {
// 	stmt, err := db.Prepare("UPDATE list_of_payments SET payment_date = $1, price = $2 WHERE id = $3")
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(entry.Date, entry.Price, entry.ID)
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}

// 	return nil
// }

// func deleteEntry(w http.ResponseWriter, r *http.Request) {
// 	// Чтение и декодирование JSON из запроса
// 	var entry Entry
// 	err := json.NewDecoder(r.Body).Decode(&entry)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		print(err.Error())
// 		return
// 	}

// 	// Проверка валидности токена
// 	if !isValidToken(entry.Token) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
	
// 	switch entry.ID_table {
// 	case 1: // Удаление данных из базы данных записи по id в таблице 1
// 		err = deleteEntryFromDB(entry.ID)
// 		if err != nil {
// 			print(err.Error())
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 	case 2: // Удаление данных из базы данных записи по id в таблице 1
// 	err = deleteRecordFromDBtoListPrices(entry.ID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		print(err.Error())
// 		return
// 		}
// 	case 3: // Удаление данных из базы данных записи по id в таблице 1
// 	err = deleteRecordFromDBtoListPayments(entry.ID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		print(err.Error())
// 		return
// 		}
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

// func deleteEntryFromDB(entryID int) error {
// 	stmt, err := db.Prepare("DELETE FROM maloesareevo WHERE id = $1")
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(entryID)
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}

// 	return nil
// }

// func deleteRecordFromDBtoListPrices(entryID int) error {
// 	stmt, err := db.Prepare("DELETE FROM list_of_price WHERE id = $1")
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(entryID)
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}

// 	return nil
// }

// func deleteRecordFromDBtoListPayments(entryID int) error {
// 	stmt, err := db.Prepare("DELETE FROM list_of_payments WHERE id = $1")
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(entryID)
// 	if err != nil {
// 		print(err.Error())
// 		return err
// 	}

// 	return nil
// }

func createRecordFromWorkspace(w http.ResponseWriter, r *http.Request) {
	var record Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = insertRecordToWorkspace(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func createRecordFromListPrice(w http.ResponseWriter, r *http.Request) {
	var record ListPriceRecord
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = insertRecordToListPrice(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func createRecordFromListPayment(w http.ResponseWriter, r *http.Request) {
	var record ListPaymentRecord
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = insertRecordToListPayment(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateRecordFromWorkspace(w http.ResponseWriter, r *http.Request) {
	var record Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = updateRecordInWorkspace(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateRecordFromListPrice(w http.ResponseWriter, r *http.Request) {
	var record ListPriceRecord
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = updateRecordInListPrice(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateRecordFromListPayments(w http.ResponseWriter, r *http.Request) {
	var record ListPaymentRecord
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = updateRecordInListPayment(record)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteRecordFromWorkspace(w http.ResponseWriter, r *http.Request) {
	var record Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = deleteRecordFromWorkspace(record.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteRecordFromListPrice(w http.ResponseWriter, r *http.Request) {
	var record ListPriceRecord
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = deleteRecordFromListPrice(record.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteRecordFromListPayments(w http.ResponseWriter, r *http.Request) {
	var record ListPaymentRecord
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = deleteRecordFromListPayment(record.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getAllRecords(w http.ResponseWriter, r *http.Request) {
	// Получение всех записей из базы данных
	rows, err := db.Query("SELECT * FROM maloesareevo")
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Создание среза для хранения всех записей
	var entries []Entry

	// Итерация по результатам запроса и добавление записей в срез
	for rows.Next() {
		var entry Entry
		err := rows.Scan(
			&entry.ID,
			&entry.ID_table,
			&entry.Token,
			&entry.Date,
			&entry.Price,
			&entry.TimeWork,
			&entry.Penalty,
			&entry.Payments,
		)
		if err != nil {
			print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		entries = append(entries, entry)
	}

	// Проверка наличия ошибок при выполнении запроса
	if err = rows.Err(); err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Кодирование среза записей в формат JSON и отправка клиенту
	err = json.NewEncoder(w).Encode(entries)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func getAllRecordsFromListPrice(w http.ResponseWriter, r *http.Request) {
	// Получение всех записей из таблицы "list_of_price"
	rows, err := db.Query("SELECT * FROM list_of_price")
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Создание среза для хранения всех записей
	var entries []Entry

	// Итерация по результатам запроса и добавление записей в срез
	for rows.Next() {
		var entry Entry
		err := rows.Scan(
			&entry.ID,
			&entry.Date,
			&entry.Price,
		)
		if err != nil {
			print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		entries = append(entries, entry)
	}

	// Проверка наличия ошибок при выполнении запроса
	if err = rows.Err(); err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Кодирование среза записей в формат JSON и отправка клиенту
	err = json.NewEncoder(w).Encode(entries)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getAllRecordsFromListPayments(w http.ResponseWriter, r *http.Request) {
	// Получение всех записей из таблицы "list_of_payments"
	rows, err := db.Query("SELECT * FROM list_of_payments")
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Создание среза для хранения всех записей
	var entries []Entry

	// Итерация по результатам запроса и добавление записей в срез
	for rows.Next() {
		var entry Entry
		err := rows.Scan(
			&entry.ID,
			&entry.Date,
			&entry.Price,
		)
		if err != nil {
			print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		entries = append(entries, entry)
	}

	// Проверка наличия ошибок при выполнении запроса
	if err = rows.Err(); err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Кодирование среза записей в формат JSON и отправка клиенту
	err = json.NewEncoder(w).Encode(entries)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func isValidToken(token string) bool {
	// Реализуйте проверку валидности токена здесь
	// Возвращайте true, если токен валиден, и false в противном случае
	// Это может включать проверку в базе данных или другую логику проверки
	// В данном примере функция всегда возвращает true
	return true
}
