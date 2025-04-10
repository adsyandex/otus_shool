package main

import (
	"log"
	"net"
	"os"

	"github.com/adsyandex/otus_shool/todo/internal/api/grpc/service"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	pb "github.com/adsyandex/otus_shool/todo/internal/api/grpc/pb"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Создаем хранилище
	store := storage.NewCSVStorage("data/tasks.csv")
	
	// Проверяем соответствие интерфейсу
	var _ storage.Storage = (*storage.CSVStorage)(nil)

	// Создаем gRPC сервер
	server := service.NewTodoServer(store)

	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, server)

	log.Printf("gRPC server started on :%s", os.Getenv("GRPC_PORT"))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}