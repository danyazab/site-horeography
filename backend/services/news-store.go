package services

import "sync"

// Для прикладу зберігаємо новини просто як рядки
// (Якщо потрібно додати заголовок, дату тощо - робимо структуру NewsItem)

var newsData = struct {
	sync.RWMutex
	items []string
}{}

// AddNewsItem додає новину
func AddNewsItem(text string) {
	newsData.Lock()
	defer newsData.Unlock()
	newsData.items = append(newsData.items, text)
}

// GetAllNewsItems повертає копію списку новин
func GetAllNewsItems() []string {
	newsData.RLock()
	defer newsData.RUnlock()
	// Повертаємо новий слайс, щоб випадково не змінювати "items" зовні
	return append([]string(nil), newsData.items...)
}
