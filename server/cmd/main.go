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
	PriceDay int    `json:"price_day"`
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
		log.Fatal(err)
	}
	defer db.Close()
	
	// Создание таблицы в базе данных (если она еще не существует)
	// createTable()

	// Инициализация маршрутизатора
	router := mux.NewRouter()

	// Установка обработчиков запросов
	router.HandleFunc("/entry", createEntry).Methods("POST")
	router.HandleFunc("/entry", updateEntry).Methods("PUT")
	router.HandleFunc("/entry", deleteEntry).Methods("DELETE")

	// Запуск сервера на порту 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

// func createTable() {
// 	_, err := db.Exec(`
// 		CREATE TABLE IF NOT EXISTS MaloeSareevo (
// 			id SERIAL PRIMARY KEY,
// 			date VARCHAR(10) NOT NULL,
// 			price FLOAT NOT NULL,
// 			time_work FLOAT NOT NULL,
// 			penalty FLOAT NOT NULL,
// 			price_day FLOAT NOT NULL,
// 			payments FLOAT[] NOT NULL
// 		)
// 	`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func createEntry(w http.ResponseWriter, r *http.Request) {
	// Чтение и декодирование JSON из запроса
	
	var entry Entry
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(entry.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	// Вставка данных в базу данных
	err = insertEntryToDB(entry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	print("succesfull")
	w.WriteHeader(http.StatusOK)
}

func insertEntryToDB(entry Entry) error {
	print("begin")
	stmt, err := db.Prepare("INSERT INTO maloesareevo (work_date, price, time_work, penalty) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
		// return err
	}
	defer stmt.Close()
	print("such")
	_, err = stmt.Exec(entry.Date, entry.Price, entry.TimeWork, entry.Penalty)
	if err != nil {
		print(err.Error())
		return err
	}

	return nil
}

func updateEntry(w http.ResponseWriter, r *http.Request) {
	// Чтение и декодирование JSON из запроса
	var entry Entry
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(entry.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Обновление данных в базе данных
	err = updateEntryInDB(entry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func updateEntryInDB(entry Entry) error {
	stmt, err := db.Prepare("UPDATE maloesareevo SET work_date = $1, price = $2, time_work = $3, penalty = $4 WHERE id = $5")
	if err != nil {
		print(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(entry.Date, entry.Price, entry.TimeWork, entry.Penalty, entry.ID)
	if err != nil {
		print(err.Error())
		return err
	}

	return nil
}

func deleteEntry(w http.ResponseWriter, r *http.Request) {
	// Чтение и декодирование JSON из запроса
	var entry Entry
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка валидности токена
	if !isValidToken(entry.Token) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Удаление данных из базы данных
	err = deleteEntryFromDB(entry.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteEntryFromDB(entryID int) error {
	stmt, err := db.Prepare("DELETE FROM MaloeSareevo WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(entryID)
	if err != nil {
		return err
	}

	return nil
}

func isValidToken(token string) bool {
	// Реализуйте проверку валидности токена здесь
	// Возвращайте true, если токен валиден, и false в противном случае
	// Это может включать проверку в базе данных или другую логику проверки
	// В данном примере функция всегда возвращает true
	return true
}
