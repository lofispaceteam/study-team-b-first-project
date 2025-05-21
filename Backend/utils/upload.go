package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadFile загружает файл и возвращает URL.
func UploadFile(c *gin.Context, fieldname string) (string, error) {
	file, err := c.FormFile(fieldname)
	if err != nil {
		return "", fmt.Errorf("не удалось получить файл формы: %v", err)
	}

	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	dst := "./uploads/" + filename

	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		return "", fmt.Errorf("не удалось создать директорию: %v", err)
	}

	if err := c.SaveUploadedFile(file, dst); err != nil {
		return "", fmt.Errorf("не удалось сохранить файл: %v", err)
	}

	url := fmt.Sprintf("http://localhost:8080/uploads/%s", filename)
	return url, nil
}
