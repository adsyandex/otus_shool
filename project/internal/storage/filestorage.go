package storage

import (
	"encoding/json"
	"os"
	"example.com/otus_shool/project/internal/model/task"
)

const storageFile = "tasks.json"

type FileStorage struct {
	filePath string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{filePath: path}
}

func (fs *FileStorage) SaveTasks(tasks []task.Task) error {
	file, err := os.Create(fs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func (fs *FileStorage) LoadTasks() ([]task.Task, error) {
	file, err := os.Open(fs.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []task.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
