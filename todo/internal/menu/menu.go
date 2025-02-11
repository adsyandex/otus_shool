package menu

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "github.com/adsyandex/otus_shool/internal/task"
)

type Menu struct {
    taskManager *task.TaskManager
}

func NewMenu(taskManager *task.TaskManager) *Menu {
    return &Menu{taskManager: taskManager}
}

func (m *Menu) Run() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("\nВыберите действие:")
        fmt.Println("1. Добавить задачу")
        fmt.Println("2. Просмотреть задачи")
        fmt.Println("3. Обновить задачу")
        fmt.Println("4. Удалить задачу")
        fmt.Println("5. Выйти")
        fmt.Print("> ")

        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            m.addTask(scanner)
        case "2":
            m.viewTasks()
        case "3":
            m.updateTask(scanner)
        case "4":
            m.deleteTask(scanner)
        case "5":
            fmt.Println("Выход из программы.")
            return
        default:
            fmt.Println("Неверный выбор. Пожалуйста, выберите снова.")
        }
    }
}

func (m *Menu) addTask(scanner *bufio.Scanner) {
    fmt.Print("Введите заголовок задачи: ")
    scanner.Scan()
    title := scanner.Text()

    fmt.Print("Введите описание задачи: ")
    scanner.Scan()
    description := scanner.Text()

    err := m.taskManager.AddTask(title, description)
    if err != nil {
        fmt.Println("Ошибка при добавлении задачи:", err)
    } else {
        fmt.Println("Задача успешно добавлена.")
    }
}

func (m *Menu) viewTasks() {
    tasks, err := m.taskManager.GetTasks()
    if err != nil {
        fmt.Println("Ошибка при получении задач:", err)
        return
    }
    if len(tasks) == 0 {
        fmt.Println("Задачи отсутствуют.")
        return
    }
    for _, t := range tasks {
        fmt.Printf("ID: %d, Заголовок: %s, Описание: %s, Статус: %s, Создано: %s\n",
            t.ID, t.Title, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04:05"))
    }
}

func (m *Menu) updateTask(scanner *bufio.Scanner) {
    fmt.Print("Введите ID задачи для обновления: ")
    scanner.Scan()
    id, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Println("Неверный ID задачи.")
        return
    }

    fmt.Print("Введите новый заголовок задачи (оставьте пустым, чтобы не изменять): ")
    scanner.Scan()
    title := scanner.Text()

    fmt.Print("Введите новое описание задачи (оставьте пустым, чтобы не изменять): ")
    scanner.Scan()
    description := scanner.Text()

    fmt.Print("Введите новый статус задачи (оставьте пустым, чтобы не изменять): ")
    scanner.Scan()
    status := scanner.Text()

    err = m.taskManager.UpdateTask(id, title, description, status)
    if err != nil {
        fmt.Println("Ошибка при обновлении задачи:", err)
    } else {
        fmt.Println("Задача успешно обновлена.")
    }
}

func (m *Menu) deleteTask(scanner *bufio.Scanner) {
    fmt.Print("Введите ID задачи для удаления: ")
    scanner.Scan()
    id, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Println("Неверный ID задачи.")
        return
    }

    err = m.taskManager.DeleteTask(id)
    if err != nil {
        fmt.Println("Ошибка при удалении задачи:", err)
    } else {
        fmt.Println("Задача успешно удалена.")
    }
}