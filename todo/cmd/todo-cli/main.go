package main

import (
	"fmt"
	"os"
	"context"
	
	"github.com/adsyandex/otus_shool/todo/internal/config"
	"github.com/adsyandex/otus_shool/todo/internal/service"
	postgresstorage "github.com/adsyandex/otus_shool/todo/internal/storage/postgres"
	"github.com/adsyandex/otus_shool/todo/cmd/todo-cli/commands"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	cfg := config.MustLoad()
	storage := postgresstorage.MustNewStorage(cfg.Postgres)
	service := service.NewTaskService(storage)
	ctx := context.Background()

	switch os.Args[1] {
	case "add":
		commands.Add(ctx, service, os.Args[2:])
	case "list":
		commands.List(ctx, service, os.Args[2:])
	case "complete":
		commands.Complete(ctx, service, os.Args[2:])
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Todo CLI Manager")
	fmt.Println("Usage:")
	fmt.Println("  add <title> [description] - Add new task")
	fmt.Println("  list [--completed]       - List tasks")
	fmt.Println("  complete <id>            - Mark task as done")
	os.Exit(0)
}