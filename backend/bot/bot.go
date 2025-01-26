package bot

import (
	"log"
	"os"
	"site-horeography/backend/services"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// StartBot запускає long-polling бота
func StartBot() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Println("TELEGRAM_BOT_TOKEN not set, skipping bot start.")
		return
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Error creating Telegram bot: %v", err)
	}

	// Якщо потрібно «приховати» логи — вимикаємо debug
	bot.Debug = false

	log.Printf("Бот авторизувався як %s", bot.Self.UserName)

	// Зчитуємо список дозволених chat_id
	// Наприклад, TELEGRAM_ADMIN_CHAT_IDS=12345,67890
	adminIDsStr := os.Getenv("TELEGRAM_ADMIN_CHAT_IDS")
	adminIDs := make(map[int64]bool)
	if adminIDsStr != "" {
		parts := strings.Split(adminIDsStr, ",")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			id, err := strconv.ParseInt(p, 10, 64)
			if err == nil {
				adminIDs[id] = true
			}
		}
	}

	// Налаштування для отримання оновлень
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // Якщо користувач надіслав текстове повідомлення
			chatID := update.Message.Chat.ID

			if !adminIDs[chatID] {
				// Користувач не є у списку дозволених — відмовляємо або ігноруємо
				msg := tgbotapi.NewMessage(chatID, "У вас немає прав для використання цього бота.")
				bot.Send(msg)
				continue
			}

			// Якщо користувач (адмін) є у списку, обробляємо команди
			text := update.Message.Text

			switch text {
			case "/start":
				reply := tgbotapi.NewMessage(chatID, "Вітаю, адміне! Ви авторизовані.")
				bot.Send(reply)

			case "/help":
				helpText := "/start - початок\n/help - довідка\n/test - тестове повідомлення"
				bot.Send(tgbotapi.NewMessage(chatID, helpText))

			case "/test":
				bot.Send(tgbotapi.NewMessage(chatID, "Це тестова відповідь"))

			default:
				// Якщо адмін, обробляємо його повідомлення
				if strings.HasPrefix(text, "/addnews") {
					// Команда вигляду /addnews <тут-текст-новини>
					// Відокремлюємо "/addnews " (8 символів + пробіл = 9)
					// Або розділити за пробілом, якщо треба.
					parts := strings.SplitN(text, " ", 2)
					if len(parts) < 2 {
						// Немає тексту після команди
						reply := tgbotapi.NewMessage(chatID, "Використання: /addnews Текст новини")
						bot.Send(reply)
						continue
					}

					newsText := parts[1] // Текст після /addnews
					// Зберігаємо в сховище
					services.AddNewsItem(newsText)

					// Відповідаємо адміну
					reply := tgbotapi.NewMessage(chatID, "Новину додано!")
					bot.Send(reply)

				}
				// Відповідаємо, що не розуміємо команду
				bot.Send(tgbotapi.NewMessage(chatID, "Невідома команда. Спробуйте /help"))
			}
		}
	}
}
