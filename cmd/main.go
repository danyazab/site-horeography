package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"site-horeography/backend/bot"
	"site-horeography/backend/handlers"
)

var tmpl *template.Template

func main() {
	// Ініціалізуємо Echo
	e := echo.New()

	// (Опційно) використання middleware (логування, відновлення тощо)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Запускаємо Телеграм-бота у горутині
	go bot.StartBot()

	// Визначаємо шляхи до статичних файлів і шаблонів
	templatesPath := filepath.Join("..", "frontend", "templates", "*.html")
	staticPath := filepath.Join("..", "frontend")

	// Завантажуємо всі .html шаблони
	files, err := filepath.Glob("./frontend/templates/*.html")
	if err != nil || len(files) == 0 {
		log.Fatalf("No templates found in path: %s", templatesPath)
	}
	tmpl = template.Must(template.ParseGlob(templatesPath))
	log.Printf("Templates loaded from: %s", templatesPath)

	// Налаштування статичних файлів
	staticAbsPath, err := filepath.Abs(staticPath)
	if err != nil {
		log.Fatalf("Failed to resolve static path: %v", err)
	}
	log.Printf("Serving static files from: %s", staticAbsPath)
	e.Static("/frontend", staticAbsPath)

	// Реєструємо маршрути
	e.GET("/", homeHandler)
	e.GET("/administration", adminHandler)
	e.GET("/support", supportHandler)
	e.GET("/graduates", graduatesHandler)
	e.GET("/cost-education", costEducationHandler)
	e.GET("/admission-requirements", admissionRequirementsHandler)
	e.GET("/educational-programs", educationalProgramsHandler)
	e.GET("/history", historyHandler)
	e.GET("/structural-divisions", structuralDivisionsHandler)

	e.GET("/contact", contactHandler)
	e.POST("/form-contact", handlers.ContactHandler)

	fmt.Println("Сервер запущено на http://localhost:8089")
	log.Fatal(e.Start(":8089"))
}

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
		"EmailURL":     "ballet.school.lviv@ukr.net",
		"Address":      "м. Львів, вул. Дорошенка 63",
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "home.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Інші обробники залишаються без змін

// adminHandler - новий маршрут (GET /administration)
// Приклад рендеру сторінки "Адміністрація" з даними, якщо потрібно
func adminHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		// Можливо, додаткові поля, потрібні для admin.html,
		// наприклад, список осіб / інформації про адміністрацію.
		"Phone1":       "+380 (97) 986 49 05",
		"Phone2":       "+380 (63) 309 32 34",
		"FacebookURL":  "https://facebook.com/yourpage",
		"InstagramURL": "https://instagram.com/yourpage",
		"EmailURL":     "ballet.school.lviv@ukr.net ",
		"Address":      "м. Львів, вул. Дорошенка 63",
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
		"Phone1":       "+380 (97) 986 49 05",
		"Phone2":       "+380 (63) 309 32 34",
		"FacebookURL":  "https://facebook.com/yourpage",
		"InstagramURL": "https://instagram.com/yourpage",
		"EmailURL":     "ballet.school.lviv@ukr.net ",
		"Address":      "м. Львів, вул. Дорошенка 63",
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
		"Phone1":       "+380 (97) 986 49 05",
		"Phone2":       "+380 (63) 309 32 34",
		"FacebookURL":  "https://facebook.com/yourpage",
		"InstagramURL": "https://instagram.com/yourpage",
		"EmailURL":     "ballet.school.lviv@ukr.net ",
		"Address":      "м. Львів, вул. Дорошенка 63",
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "graduates.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func costEducationHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		// Можливо, додаткові поля, потрібні для admin.html,
		// наприклад, список осіб / інформації про адміністрацію.
		"Phone1":       "+380 (97) 986 49 05",
		"Phone2":       "+380 (63) 309 32 34",
		"FacebookURL":  "https://facebook.com/yourpage",
		"InstagramURL": "https://instagram.com/yourpage",
		"EmailURL":     "ballet.school.lviv@ukr.net ",
		"Address":      "м. Львів, вул. Дорошенка 63",
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "cost-education.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}
func admissionRequirementsHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		// Можливо, додаткові поля, потрібні для admin.html,
		// наприклад, список осіб / інформації про адміністрацію.
		"Phone1":       "+380 (97) 986 49 05",
		"Phone2":       "+380 (63) 309 32 34",
		"FacebookURL":  "https://facebook.com/yourpage",
		"InstagramURL": "https://instagram.com/yourpage",
		"EmailURL":     "ballet.school.lviv@ukr.net ",
		"Address":      "м. Львів, вул. Дорошенка 63",
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "admission-requirements.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}
func educationalProgramsHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		// Можливо, додаткові поля, потрібні для admin.html,
		// наприклад, список осіб / інформації про адміністрацію.
		"Phone1":       "+380 (97) 986 49 05",
		"Phone2":       "+380 (63) 309 32 34",
		"FacebookURL":  "https://facebook.com/yourpage",
		"InstagramURL": "https://instagram.com/yourpage",
		"EmailURL":     "ballet.school.lviv@ukr.net ",
		"Address":      "м. Львів, вул. Дорошенка 63",
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "educational-programs.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}
func historyHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		// Можливо, додаткові поля, потрібні для admin.html,
		// наприклад, список осіб / інформації про адміністрацію.
		"Phone1":       "+380 (97) 986 49 05",
		"Phone2":       "+380 (63) 309 32 34",
		"FacebookURL":  "https://facebook.com/yourpage",
		"InstagramURL": "https://instagram.com/yourpage",
		"EmailURL":     "ballet.school.lviv@ukr.net ",
		"Address":      "м. Львів, вул. Дорошенка 63",
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "history.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}
func structuralDivisionsHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		// Можливо, додаткові поля, потрібні для admin.html,
		// наприклад, список осіб / інформації про адміністрацію.
		"Phone1":       "+380 (97) 986 49 05",
		"Phone2":       "+380 (63) 309 32 34",
		"FacebookURL":  "https://facebook.com/yourpage",
		"InstagramURL": "https://instagram.com/yourpage",
		"EmailURL":     "ballet.school.lviv@ukr.net ",
		"Address":      "м. Львів, вул. Дорошенка 63",
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "structural-divisions.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func contactHandler(c echo.Context) error {
	data := map[string]interface{}{
		"SchoolName": "Львівська Хореографічна Школа",
		// Можливо, додаткові поля, потрібні для admin.html,
		// наприклад, список осіб / інформації про адміністрацію.
		"Phone1":       "+380 (97) 986 49 05",
		"Phone2":       "+380 (63) 309 32 34",
		"FacebookURL":  "https://facebook.com/yourpage",
		"InstagramURL": "https://instagram.com/yourpage",
		"EmailURL":     "ballet.school.lviv@ukr.net ",
		"Address":      "м. Львів, вул. Дорошенка 63",
	}

	err := tmpl.ExecuteTemplate(c.Response().Writer, "contact.html", data)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}
