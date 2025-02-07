package account

type User struct {
    ID       int
    Username string
    password string // приватное поле
}

// Метод для установки пароля
func (u *User) SetPassword(p string) {
    u.password = p
}

// Метод для проверки пароля
func (u *User) CheckPassword(p string) bool {
    return u.password == p
}
