package model

// Entity - общий интерфейс для всех структур
type Entity interface {
	GetType() string
}

// User - структура пользователя
type User struct {
	ID   int
	Name string
}

// Product - структура продукта
type Product struct {
	ID    int
	Title string
	Price float64
}

// Реализация метода GetType для User
func (u User) GetType() string {
	return "User"
}

// Реализация метода GetType для Product
func (p Product) GetType() string {
	return "Product"
}
