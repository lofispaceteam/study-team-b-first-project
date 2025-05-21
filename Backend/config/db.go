package config

import (
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=24682468 dbname=project_api port=7777 sslmode=disable"
	log.Printf("Попытка подключения к базе данных с DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Ошибка при подключении к базе данных: %v", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Ошибка при получении базы данных SQL: %v", err)
		return nil, err
	}

	// Установка максимального количества пустых соединений
	sqlDB.SetMaxIdleConns(10)

	// Установка максимального количества открытых соединений
	sqlDB.SetMaxOpenConns(100)

	// Установка времени жизни соединений
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Проверка подключения с простым запросом SELECT 1
	var one int
	err = sqlDB.QueryRow("SELECT 1").Scan(&one)
	if err != nil {
		log.Printf("Ошибка при выполнении проверочного запроса: %v", err)
		return nil, err
	}

	log.Println("Проверочный запрос выполнен успешно. Соединение с базой данных активно.")
	log.Println("Успешное подключение к базе данных")
	return db, nil
}
