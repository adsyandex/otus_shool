package service

import (
	"context"
	"time"

	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
	pb "github.com/adsyandex/otus_shool/todo/internal/api/grpc/pb"
	"github.com/google/uuid"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
	storage storage.Storage
}

func NewTodoServer(storage storage.Storage) *TodoServer {
	return &TodoServer{storage: storage}
}

func (s *TodoServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	task := &models.Task{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.storage.CreateTask(ctx, task); err != nil {
		return nil, err
	}

	return &pb.TaskResponse{Task: task.ToProto()}, nil
}

func (s *TodoServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	task, err := s.storage.GetTask(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.TaskResponse{Task: task.ToProto()}, nil
}

// Аналогично реализуйте остальные методы (UpdateTask, DeleteTask, ListTasks)