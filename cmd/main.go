package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"site-horeography/handlers"
)

var tmpl *template.Template

func main() {
	// Ініціалізуємо Echo
	e := echo.New()

	// (Опційно) використання middleware (логування, відновлення тощо)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Завантажуємо усі .html-шаблони з папки "templates"
	tmpl = template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

	// Налаштовуємо статичні файли (CSS, JS, зображення тощо)
	e.Static("/static", "static")

	// Реєструємо маршрути:
	e.GET("/", homeHandler)
	// Новий маршрут для сторінки "Адміністрація"
	e.GET("/administration", adminHandler)
	e.GET("/support", supportHandler)
	e.GET("/graduates", graduatesHandler)
	// Обробка форми з контактами (POST)
	e.POST("/contact", handlers.ContactHandler)

	fmt.Println("Сервер запущено на http://localhost:8089")
	log.Fatal(e.Start(":8089"))
}

// homeHandler - замість homeHandler(w http.ResponseWriter, r *http.Request)
func homeHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		"Quote":      "«Унікальний осередок мистецтва та професіоналізму, де кожна дитина має шанс відкрити для себе чарівний світ танцю.»",
		"Description": `Ми пропонуємо всебічну підготовку, поєднуючи класичну балетну техніку 
з сучасними танцювальними стилями, що дозволяє нашим учням розвивати не лише фізичні, 
але й творчі здібності. Наші учні не лише опановують складну танцювальну техніку, але й 
вчаться працювати в команді, виступати на сцені та досягати поставлених цілей, що є незамінним 
досвідом для їхнього майбутнього. Заняття балетом з нами сформують у вашої дитини дисципліну, 
витривалість, почуття ритму та артистизм, що є важливими навичками для будь-якої життєвої сфери.`,
		"Phone1":       "+380 (97) 986 49 05",
		"Phone2":       "+380 (63) 309 32 34",
		"FacebookURL":  "https://facebook.com/yourpage",
		"InstagramURL": "https://instagram.com/yourpage",
		"EmailURL":     "ballet.school.lviv@ukr.net ",
		"Address":      "м. Львів, вул. Дорошенка 63",
	}

	// Рендеримо шаблон home.html
	err := tmpl.ExecuteTemplate(c.Response().Writer, "home.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// adminHandler - новий маршрут (GET /administration)
// Приклад рендеру сторінки "Адміністрація" з даними, якщо потрібно
func adminHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		// Можливо, додаткові поля, потрібні для admin.html,
		// наприклад, список осіб / інформації про адміністрацію.
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "admin.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func supportHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		// Можливо, додаткові поля, потрібні для admin.html,
		// наприклад, список осіб / інформації про адміністрацію.
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "support.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func graduatesHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		// Можливо, додаткові поля, потрібні для admin.html,
		// наприклад, список осіб / інформації про адміністрацію.
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "graduates.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}
