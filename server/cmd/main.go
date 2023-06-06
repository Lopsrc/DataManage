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

// delete print(err.Error())
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "postgres"
)

var db *sql.DB //Вынести конфиги и объект db  в отдельную структуру ConnectionDB

type ListPayments struct {
	ID       int    `json:"id"`
	Date     string `json:"date"`
	PriceDay int    `json:"payments"`
	Token    string `json:"token"`
}
type ListPrices struct {
	ID    int    `json:"id"`
	Date  string `json:"date"`
	Price int    `json:"price"`
	Token string `json:"token"`
}
type Workspace struct {
	ID            int    `json:"id"`
	NameWorkspace string `json:"name_workspace"`
	Date          string `json:"date"`
	Price         int    `json:"price"`
	TimeWork      int    `json:"time_work"`
	Penalty       int    `json:"penalty"`
	Token         string `json:"token"`
}

type Token struct {
	Token string `json:"token"`
}

func main() {

	// Установка соединения с базой данных. Вынести в функцию, которая вернет ConnectionDB и Error
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
	router.HandleFunc("/workspace", createRecordFromWorkspace).Methods("POST")
	router.HandleFunc("/price", createRecordFromListPrice).Methods("POST")
	router.HandleFunc("/payments", createRecordFromListPayment).Methods("POST")

	router.HandleFunc("/workspace", updateRecordFromWorkspace).Methods("PUT")
	router.HandleFunc("/price", updateRecordFromListPrice).Methods("PUT")
	router.HandleFunc("/payments", updateRecordFromListPayments).Methods("PUT")

	router.HandleFunc("/workspace", deleteRecordFromWorkspace).Methods("DELETE")
	router.HandleFunc("/price", deleteRecordFromListPrice).Methods("DELETE")
	router.HandleFunc("/payments", deleteRecordFromListPayments).Methods("DELETE")

	router.HandleFunc("/workspace", getAllRecords).Methods("GET")                // Новый обработчик для получения всех записей
	router.HandleFunc("/price", getAllRecordsFromListPrice).Methods("GET")       // Обработчик для получения всех записей из таблицы "list_of_price"
	router.HandleFunc("/payments", getAllRecordsFromListPayments).Methods("GET") // Обработчик для получения всех записей из таблицы "list_of_payments"
	// Запуск сервера на порту 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createRecordFromWorkspace(w http.ResponseWriter, r *http.Request) {
	var record Workspace
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	print("w date = "+record.Date)
	print("w name = "+record.NameWorkspace)
	print("w token = "+record.Token)
	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = insertRecordToWorkspace(record)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func createRecordFromListPrice(w http.ResponseWriter, r *http.Request) {
	var record ListPrices
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		print(err.Error())
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
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func createRecordFromListPayment(w http.ResponseWriter, r *http.Request) {
	var record ListPayments
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		print(err.Error())
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
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateRecordFromWorkspace(w http.ResponseWriter, r *http.Request) {
	var record Workspace
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		print(err.Error())
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
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateRecordFromListPrice(w http.ResponseWriter, r *http.Request) {
	var record ListPrices
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		print(err.Error())
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
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateRecordFromListPayments(w http.ResponseWriter, r *http.Request) {
	var record ListPayments
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		print(err.Error())
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
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteRecordFromWorkspace(w http.ResponseWriter, r *http.Request) {
	var record Workspace
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = deleteRecordInWorkspace(record.ID)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteRecordFromListPrice(w http.ResponseWriter, r *http.Request) {
	var record ListPrices
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = deleteRecordInListPrice(record.ID)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteRecordFromListPayments(w http.ResponseWriter, r *http.Request) {
	var record ListPayments
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(record.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = deleteRecordInListPayment(record.ID)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Функция getAllRecords возвращает все записи из таблицы "workspace" базы данных
func getAllRecords(w http.ResponseWriter, r *http.Request) {
	var token Token
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Проверка валидности токена
	if !isValidToken(token.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Выполнение SQL-запроса для получения всех записей из таблицы "workspace"
	rows, err := db.Query("SELECT * FROM workspace")
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()

	// Создание среза для хранения записей
	var records []Workspace

	// Итерация по результатам запроса
	for rows.Next() {
		var record Workspace
		err := rows.Scan(&record.ID, &record.NameWorkspace, &record.Date, &record.Price, &record.TimeWork, &record.Penalty)
		if err != nil {
			print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		records = append(records, record)
	}

	// Проверка ошибок при итерации
	err = rows.Err()
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Преобразование записей в JSON
	jsonData, err := json.Marshal(records)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Отправка JSON-ответа клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
func getAllRecordsFromListPrice(w http.ResponseWriter, r *http.Request) {
	var token Token
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Проверка валидности токена
	if !isValidToken(token.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Получение всех записей из таблицы "list_of_price"
	rows, err := db.Query("SELECT * FROM list_of_price")
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Создание среза для хранения всех записей
	var entries []ListPrices

	// Итерация по результатам запроса и добавление записей в срез
	for rows.Next() {
		var entry ListPrices
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
	var token Token
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Проверка валидности токена
	if !isValidToken(token.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Получение всех записей из таблицы "list_of_payments"
	rows, err := db.Query("SELECT * FROM list_of_payments")
	if err != nil {
		print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Создание среза для хранения всех записей
	var entries []ListPayments

	// Итерация по результатам запроса и добавление записей в срез
	for rows.Next() {
		var entry ListPayments
		err := rows.Scan(
			&entry.ID,
			&entry.Date,
			&entry.PriceDay,
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

func insertRecordToWorkspace(record Workspace) error {
	stmt, err := db.Prepare("INSERT INTO workspace (name_workspace, work_date, price, time_work, penalty) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	print(" date = "+record.Date)
	print(" name = "+record.NameWorkspace)
	print(" token = "+record.Token)
	// Выполнение SQL-запроса для вставки записи
	_, err = stmt.Exec(record.NameWorkspace, record.Date, record.Price, record.TimeWork, record.Penalty)
	if err != nil {
		return err
	}

	return nil
}
func insertRecordToListPrice(record ListPrices) error {
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
func insertRecordToListPayment(record ListPayments) error {
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
	stmt, err := db.Prepare("UPDATE workspace SET name_workspace = $2, work_date = $3, price = $4, time_work = $5, penalty = $6 WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполнение SQL-запроса для обновления записи
	_, err = stmt.Exec(record.ID, record.NameWorkspace, record.Date, record.Price, record.TimeWork, record.Penalty)
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
	stmt, err := db.Prepare("DELETE FROM workspace WHERE id = $1")
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
