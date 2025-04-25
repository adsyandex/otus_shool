package storage

import (
    "context"
    "encoding/csv"
    "os"
    "strconv"
    
    "github.com/adsyandex/otus_shool/todo/internal/models"
)

type CSVStorage struct {
    filePath string
}

func NewCSVStorage(filePath string) *CSVStorage {
    return &CSVStorage{filePath: filePath}
}

func (s *CSVStorage) SaveTask(ctx context.Context, task models.Task) (models.Task, error) {
    file, err := os.OpenFile(s.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        return models.Task{}, err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    if task.ID == 0 {
        tasks, _ := s.GetTasks(ctx)
        task.ID = len(tasks) + 1
    }

    err = writer.Write([]string{
        strconv.Itoa(task.ID),
        task.Title,
        strconv.FormatBool(task.Completed),
    })
    return task, err
}

func (s *CSVStorage) GetTasks(ctx context.Context) ([]models.Task, error) {
    file, err := os.Open(s.filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var tasks []models.Task
    for _, record := range records {
        id, _ := strconv.Atoi(record[0])
        completed, _ := strconv.ParseBool(record[2])
        tasks = append(tasks, models.Task{
            ID:        id,
            Title:     record[1],
            Completed: completed,
        })
    }
    return tasks, nil
}

func (s *CSVStorage) GetTaskByID(ctx context.Context, id int) (models.Task, error) {
    tasks, err := s.GetTasks(ctx)
    if err != nil {
        return models.Task{}, err
    }

    for _, task := range tasks {
        if task.ID == id {
            return task, nil
        }
    }
    return models.Task{}, ErrNotFound
}

func (s *CSVStorage) UpdateTask(ctx context.Context, task models.Task) (models.Task, error) {
    tasks, err := s.GetTasks(ctx)
    if err != nil {
        return models.Task{}, err
    }

    found := false
    for i, t := range tasks {
        if t.ID == task.ID {
            tasks[i] = task
            found = true
            break
        }
    }

    if !found {
        return models.Task{}, ErrNotFound
    }

    if err := s.saveAllTasks(tasks); err != nil {
        return models.Task{}, err
    }
    return task, nil
}

func (s *CSVStorage) DeleteTask(ctx context.Context, id int) error {
    tasks, err := s.GetTasks(ctx)
    if err != nil {
        return err
    }

    newTasks := make([]models.Task, 0)
    for _, task := range tasks {
        if task.ID != id {
            newTasks = append(newTasks, task)
        }
    }

    if len(newTasks) == len(tasks) {
        return ErrNotFound
    }

    return s.saveAllTasks(newTasks)
}

func (s *CSVStorage) saveAllTasks(tasks []models.Task) error {
    file, err := os.Create(s.filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, task := range tasks {
        err := writer.Write([]string{
            strconv.Itoa(task.ID),
            task.Title,
            strconv.FormatBool(task.Completed),
        })
        if err != nil {
            return err
        }
    }
    return nil
}