package commands

import (
	"context"
	"fmt"
	
	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/service"
)

func Add(ctx context.Context, s *service.TaskService, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: title required")
		return
	}

	task := models.Task{
		Title: args[0],
	}
	
	if len(args) > 1 {
		task.Description = args[1]
	}

	created, err := s.CreateTask(ctx, task)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Created task #%s\n", created.ID)
}