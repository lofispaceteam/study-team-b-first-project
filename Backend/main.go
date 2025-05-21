package main

import (
	"log"
	"meme-api-backend/config"
	"meme-api-backend/handlers"
	"meme-api-backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Запуск приложения...")

	r := gin.Default()

	// Подключение к базе данных
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Автомиграция схемы базы данных
	log.Println("Начало автоматической миграции схемы базы данных...")
	db.AutoMigrate(&models.Meme{})
	log.Println("Автомиграция схемы базы данных завершена успешно.")

	// Инициализация обработчиков
	log.Println("Инициализация обработчиков...")
	memeHandler := handlers.NewMemeHandler(db)
	log.Println("Обработчики инициализированы успешно.")

	// Маршрутизация
	api := r.Group("/api/v1")
	{
		// Маршруты для мемов
		api.POST("/memes", memeHandler.CreateMeme)
		api.GET("/memes", memeHandler.GetMemes)
		api.GET("/memes/:id", memeHandler.GetMemeByID)
		api.PUT("/memes/:id", memeHandler.UpdateMeme)
		api.DELETE("/memes/:id", memeHandler.DeleteMeme)
		api.GET("/memes/random", memeHandler.RandomMeme)
		api.POST("/memes/upload", memeHandler.UploadImage)
	}

	// Сервирование статических файлов для загруженных изображений
	r.Static("/uploads", "./uploads")

	log.Println("Сервер запущен на :8080")
	r.Run(":8080")
}
