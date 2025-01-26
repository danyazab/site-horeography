package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type ContactResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ContactHandler(c echo.Context) error {
	// Перевіряємо метод (в Echo це можна налаштувати вже на рівні маршруту, але залишимо перевірку)
	if c.Request().Method != http.MethodPost {
		return c.String(http.StatusMethodNotAllowed, "Only POST is allowed")
	}

	name := c.FormValue("name")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	message := c.FormValue("message")

	fmt.Printf("Надійшов контакт:\nІм'я: %s\nEmail: %s\nТелефон: %s\nПовідомлення: %s\n",
		name, email, phone, message)

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	adminsStr := os.Getenv("TELEGRAM_ADMIN_CHAT_IDS") // "123456789,987654321"
	if botToken == "" || adminsStr == "" {
		fmt.Println("Помилка: не задано TELEGRAM_BOT_TOKEN або TELEGRAM_ADMIN_CHAT_IDS")
		return c.JSON(http.StatusInternalServerError, ContactResponse{
			Status:  "error",
			Message: "Налаштовано не всі змінні оточення!",
		})
	}

	// Готуємо текст
	telegramText := fmt.Sprintf(
		"Нове повідомлення з форми:\nІм'я: %s\nEmail: %s\nТелефон: %s\nПовідомлення: %s",
		name, email, phone, message,
	)

	// Розділимо список чатів за комою
	chatIDs := strings.Split(adminsStr, ",")
	// Для кожного chat_id робимо POST
	for _, chatID := range chatIDs {
		chatID = strings.TrimSpace(chatID)
		telegramAPIURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

		data := url.Values{}
		data.Set("chat_id", chatID)
		data.Set("text", telegramText)

		resp, err := http.PostForm(telegramAPIURL, data)
		if err != nil {
			fmt.Println("Помилка надсилання повідомлення:", err)
		} else {
			resp.Body.Close()
		}
	}

	return c.JSON(http.StatusOK, ContactResponse{
		Status:  "success",
		Message: "Ваше повідомлення надіслано! Найближчим часом з вами зв’яжуться.",
	})
}
