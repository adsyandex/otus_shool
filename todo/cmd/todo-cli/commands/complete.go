package commands

import (
	"context"
	"fmt"
	
	"github.com/adsyandex/otus_shool/todo/internal/service"
)

func Complete(ctx context.Context, s *service.TaskService, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: task ID required")
		fmt.Println("Usage: complete <task-id>")
		return
	}

	taskID := args[0]
	
	// Получаем текущую задачу
	task, err := s.GetTask(ctx, taskID)
	if err != nil {
		fmt.Printf("Error finding task: %v\n", err)
		return
	}

	// Помечаем как выполненную
	task.Completed = true
	
	// Обновляем в хранилище
	_, err = s.UpdateTask(ctx, task)
	if err != nil {
		fmt.Printf("Error completing task: %v\n", err)
		return
	}

	fmt.Printf("Task #%s marked as completed\n", taskID)
}