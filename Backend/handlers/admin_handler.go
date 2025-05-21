package handlers

import (
	"errors"
	"net/http"

	"meme-api-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AdminHandler содержит обработчики для админ-панели.
type AdminHandler struct {
	DB *gorm.DB
}

// NewAdminHandler создает новый экземпляр AdminHandler.
func NewAdminHandler(db *gorm.DB) *AdminHandler {
	return &AdminHandler{DB: db}
}

// GetAdminMemes возвращает список всех мемов для администратора.
func (h *AdminHandler) GetAdminMemes(c *gin.Context) {
	var memes []models.Meme
	result := h.DB.Find(&memes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, memes)
}

// DeleteAdminMeme удаляет мем по его ID для администратора.
func (h *AdminHandler) DeleteAdminMeme(c *gin.Context) {
	id := c.Param("id")
	result := h.DB.Delete(&models.Meme{}, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Мем не найден"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Мем успешно удален"})
}
