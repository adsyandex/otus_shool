# Otus_shool
## Работа над проектом «Программа для управления задачами»
### 8.02.2025 Описание проекта:
#### TODO-App 📋
#### TODO-App — это консольное приложение для управления задачами (TODO-list). Оно позволяет пользователю создавать, просматривать, обновлять и удалять задачи, а также фильтровать их по статусу и  дате создания.

🔹 Возможности
#### Добавление, просмотр, обновление и удаление задач
#### Поддержка статусов задач: в процессе, завершено и т. д.
#### Сохранение задач между запусками (файл или база данных)
#### Удобное текстовое меню для работы с задачами
#### Валидация пользовательского ввода и обработка ошибок
#### Этот проект помогает организовать задачи и улучшить управление временем, предлагая удобный инструмент для личного или командного использования.
#### Структура проекта
```

   ## Структура проекта (PostgreSQL версия 1.0.0)
todo/
├── cmd/
│ └── todo-app/
│ └── main.go # Точка входа приложения
├── config.yaml # Конфигурация приложения
├── docker-compose.yml # Docker-конфигурация (PostgreSQL + Redis)
├── migrations/ # Миграции базы данных
│ ├── 000001_init.up.sql # SQL-миграция (создание таблиц)
│ └── 000001_init.down.sql # Откат миграции
├── internal/
│ ├── api/ # HTTP-обработчики
│ │ ├── handlers.go # Реализация обработчиков
│ │ └── routes.go # Маршрутизация
│ ├── config/ # Конфигурация
│ │ └── config.go # Загрузка конфигурации
│ ├── logger/ # Логирование
│ │ └── logger.go # Инициализация логгера
│ ├── models/ # Модели данных
│ │ └── task.go # Модель задачи
│ ├── service/ # Бизнес-логика
│ │ └── task.go # Сервис задач
│ └── storage/ # Хранилища данных
│ ├── contracts/ # Интерфейсы
│ │ └── storage.go # Интерфейс Storage
│ ├── mongo/ #  Реализация для MongoDB
│ │ └── mongo_storage.go
│ └── postgres/ # Основное хранилище
│ └── postgres_storage.go
└── README.md # Документация            

```
#### 3.03.25 Добавил в структуру  проекта Logger, канал и горутину.

## 27.12.24 Обучение в Otus. Начинаю новый проект
### 10.01.2025 Создание новых веток Git
#### Создал базовую ветку проекта: main;
#### Создал новые ветки от базовой: second && dev;
#### Выполнил комит в новой ветке;
#### Запушил в удаленный репозиторий новую ветку;
#### Создал pull request и смержил изменения в базовую ветку на github;
#### Обновил локальную базовую ветку fetch, pull;
#### Проверил что все изменения комита в основной ветке.
#### 29.01.25 createChessBoard создает шахматную доску проверил пайплайны
#### 30/01/25 после сдачи дом.задания было рекомендовано для оптимизации таких операций использовать strings.Builder
