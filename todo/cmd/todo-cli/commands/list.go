package commands

import (
	"context"
	"fmt"
	"strings"
	
	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/service"
)

func List(ctx context.Context, s *service.TaskService, args []string) {
	filter := models.TaskFilter{}
	
	for _, arg := range args {
		switch arg {
		case "--completed":
			completed := true
			filter.Completed = &completed
		}
	}

	tasks, err := s.ListTasks(ctx, filter)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "✓"
		}
		fmt.Printf("[%s] %s: %s\n", status, task.ID, task.Title)
		if task.Description != "" {
			fmt.Println("   ", strings.TrimSpace(task.Description))
		}
	}
}