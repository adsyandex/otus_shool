package main

import (
	"log"
	"os"
	

	"github.com/adsyandex/otus_shool/todo/internal/api/grpc/client"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	conn, err := grpc.Dial(
		os.Getenv("GRPC_SERVER_ADDR"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	todoClient := client.NewTodoClient(conn)

	// Пример использования
	task, err := todoClient.CreateTask("gRPC Task", "Created via gRPC client")
	if err != nil {
		log.Fatalf("CreateTask failed: %v", err)
	}
	log.Printf("Created task: %v", task)
}