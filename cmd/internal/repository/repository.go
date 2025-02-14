package repository

import (
	"fmt"
	"sync"

	"github.com/adsyandex/otus_shool/cmd/internal/model"

)

var (
	users    []model.User
	products []model.Product
	mu       sync.Mutex
)

// AddEntity принимает интерфейс и распределяет его в нужный слайс
func AddEntity(e model.Entity) {
	mu.Lock()
	defer mu.Unlock()

	switch v := e.(type) {
	case model.User:
		users = append(users, v)
		fmt.Println("Добавлен пользователь:", v)
	case model.Product:
		products = append(products, v)
		fmt.Println("Добавлен продукт:", v)
	default:
		fmt.Println("Неизвестный тип:", v.GetType())
	}
}

// GetUsers возвращает список пользователей
func GetUsers() []model.User {
	return users
}

// GetProducts возвращает список продуктов
func GetProducts() []model.Product {
	return products
}
