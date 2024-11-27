package main

import (
	"log"
	"net/http"
)

func main() {
	// Обслуговування статичних файлів з папки "static"
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Запуск сервера на порті 8080
	log.Println("Сервер запущено на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
