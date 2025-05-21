package handlers

import (
	"errors"
	"log"
	"meme-api-backend/models"
	"meme-api-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MemeHandler struct {
	DB *gorm.DB
}

func NewMemeHandler(db *gorm.DB) *MemeHandler {
	return &MemeHandler{DB: db}
}

func (h *MemeHandler) CreateMeme(c *gin.Context) {
	var meme models.Meme
	if err := c.ShouldBindJSON(&meme); err != nil {
		log.Printf("Ошибка при привязке JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.DB.Create(&meme)
	if result.Error != nil {
		log.Printf("Ошибка при создании мема: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	log.Printf("Мем создан: %+v", meme)
	c.JSON(http.StatusCreated, meme)
}

func (h *MemeHandler) GetMemes(c *gin.Context) {
	var memes []models.Meme
	result := h.DB.Find(&memes)
	if result.Error != nil {
		log.Printf("Ошибка при получении мемов: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	log.Printf("Получено мемов: %d", len(memes))
	c.JSON(http.StatusOK, memes)
}

func (h *MemeHandler) GetMemeByID(c *gin.Context) {
	id := c.Param("id")
	var meme models.Meme
	result := h.DB.First(&meme, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Мем с ID %s не найден", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Мем не найден"})
			return
		}
		log.Printf("Ошибка при получении мема с ID %s: %v", id, result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	log.Printf("Получен мем: %+v", meme)
	c.JSON(http.StatusOK, meme)
}

func (h *MemeHandler) UpdateMeme(c *gin.Context) {
	id := c.Param("id")
	var meme models.Meme
	result := h.DB.First(&meme, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Мем с ID %s не найден", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Мем не найден"})
			return
		}
		log.Printf("Ошибка при получении мема с ID %s: %v", id, result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var input models.Meme
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Ошибка при привязке JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	meme.Title = input.Title
	meme.ImageURL = input.ImageURL

	result = h.DB.Save(&meme)
	if result.Error != nil {
		log.Printf("Ошибка при обновлении мема с ID %s: %v", id, result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	log.Printf("Мем обновлён: %+v", meme)
	c.JSON(http.StatusOK, meme)
}

func (h *MemeHandler) DeleteMeme(c *gin.Context) {
	id := c.Param("id")
	result := h.DB.Delete(&models.Meme{}, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Мем с ID %s не найден", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Мем не найден"})
			return
		}
		log.Printf("Ошибка при удалении мема с ID %s: %v", id, result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	log.Printf("Мем с ID %s успешно удалён", id)
	c.JSON(http.StatusOK, gin.H{"message": "Мем успешно удален"})
}

func (h *MemeHandler) RandomMeme(c *gin.Context) {
	var memes []models.Meme
	result := h.DB.Order("RANDOM()").Limit(1).Find(&memes)
	if result.Error != nil {
		log.Printf("Ошибка при получении случайного мема: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if len(memes) == 0 {
		log.Println("Мемы не найдены")
		c.JSON(http.StatusNotFound, gin.H{"error": "Мемы не найдены"})
		return
	}

	log.Printf("Случайный мем: %+v", memes[0])
	c.JSON(http.StatusOK, memes[0])
}

func (h *MemeHandler) UploadImage(c *gin.Context) {
	url, err := utils.UploadFile(c, "image")
	if err != nil {
		log.Printf("Ошибка при загрузке изображения: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Изображение загружено: %s", url)
	c.JSON(http.StatusOK, gin.H{"message": "Изображение успешно загружено", "url": url})
}
