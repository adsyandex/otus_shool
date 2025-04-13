package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage"
	pb "github.com/adsyandex/otus_shool/todo/internal/api/grpc/pb"
)

type TodoServer struct {
	pb.UnimplementedTodoServiceServer
	storage storage.Storage
}

func NewTodoServer(storage storage.Storage) *TodoServer {
	return &TodoServer{storage: storage}
}

func (s *TodoServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate ID: %v", err)
	}

	task := &models.Task{
		ID:          id.String(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.storage.CreateTask(ctx, task); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create task: %v", err)
	}

	return &pb.TaskResponse{Task: task.ToProto()}, nil
}

func (s *TodoServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	task, err := s.storage.GetTask(ctx, req.GetId())
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "task not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get task: %v", err)
	}

	return &pb.TaskResponse{Task: task.ToProto()}, nil
}

func (s *TodoServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	// Получаем текущую версию задачи
	existingTask, err := s.storage.GetTask(ctx, req.GetId())
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "task not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get task: %v", err)
	}

	// Обновляем поля
	existingTask.Title = req.GetTitle()
	existingTask.Description = req.GetDescription()
	existingTask.Completed = req.GetCompleted()
	existingTask.UpdatedAt = time.Now()

	if err := s.storage.UpdateTask(ctx, existingTask); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update task: %v", err)
	}

	return &pb.TaskResponse{Task: existingTask.ToProto()}, nil
}

func (s *TodoServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	if err := s.storage.DeleteTask(ctx, req.GetId()); err != nil {
		if err == storage.ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "task not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete task: %v", err)
	}

	return &pb.DeleteTaskResponse{Success: true}, nil
}

func (s *TodoServer) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	tasks, err := s.storage.ListTasks(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list tasks: %v", err)
	}

	pbTasks := make([]*pb.Task, 0, len(tasks))
	for _, task := range tasks {
		pbTasks = append(pbTasks, task.ToProto())
	}

	return &pb.ListTasksResponse{
		Tasks: pbTasks,
		Total: int32(len(pbTasks)),
	}, nil
}