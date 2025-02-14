# Otus_shool
## 14/02/25 Создание интерфейсов
### Структура проекта для интерфейсов
-----------------------------------------------------------

```
📂 project-root
├── 📂 cmd
│ ├── 📝 main.go
├── 📂 internal
│ ├── 📂 model
│ │ ├── 📝 models.go (определение структур и интерфейсов)
│ ├── 📂 repository
│ │ ├── 📝 repository.go (добавление структур в слайсы)
│ ├── 📂 service
│ │ ├── 📝 service.go (генерация структур и вызов репозитория)
```
 **Разбор логики:**
  1. **models.go**
   - Определяет интерфейс Entity с методом GetType().
   - Создает структуры User и Product, реализующие Entity.

  2. **repository.go**
   - Определяет глобальные слайсы users и products.
   - Функция AddEntity принимает Entity, определяет его тип и добавляет в соответствующий слайс.

  3. **service.go**
   - Функция StartDataGeneration каждую секунду создает 
     новый User и Product и передает в репозиторий.

  4. **main.go**
   - Запускает StartDataGeneration в горутине.
   - Через 10 секунд выводит содержимое слайсов.
##### Код демонстрирует работу с слайсами, интерфейсами и динамическим определением типа переданного объекта. 🚀


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

