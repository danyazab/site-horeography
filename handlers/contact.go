package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContactResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ContactHandler - обробка POST-форми з контактними даними
func ContactHandler(c echo.Context) error {
	// Перевіряємо метод (в Echo це можна налаштувати вже на рівні маршруту, але залишимо перевірку)
	if c.Request().Method != http.MethodPost {
		return c.String(http.StatusMethodNotAllowed, "Only POST is allowed")
	}

	// Отримуємо дані з форми
	name := c.FormValue("name")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	message := c.FormValue("message")

	// Друкуємо у консоль для тесту
	fmt.Printf("Надійшов контакт:\nІм'я: %s\nEmail: %s\nТелефон: %s\nПовідомлення: %s\n",
		name, email, phone, message)

	// Повертаємо JSON-відповідь
	resp := ContactResponse{
		Status:  "success",
		Message: "Ваше повідомлення надіслано! Найближчим часом з вами зв’яжуться.",
	}
	return c.JSON(http.StatusOK, resp)
}
