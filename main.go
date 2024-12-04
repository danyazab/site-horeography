package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	// Ініціалізація Echo
	e := echo.New()

	// Увімкнення middleware для логів
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Обслуговування статичних файлів з папки "static"
	e.Static("/", "static")

	// Визначення порту для Heroku або локального запуску
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // За замовчуванням
	}

	// Запуск сервера
	log.Printf("Сервер запущено на http://localhost:%s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatal("Помилка запуску сервера:", err)
	}
}
