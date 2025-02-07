package account

import (
    "golang.org/x/crypto/bcrypt"
)

// User представляет пользователя системы
type User struct {
    ID       int
    Username string
    password string // Приватное поле (не хранится в открытом виде)
}

// SetPassword хеширует и устанавливает пароль
func (u *User) SetPassword(password string) error {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.password = string(hash)
    return nil
}

// CheckPassword сравнивает введенный пароль с хешем
func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
    return err == nil
}

