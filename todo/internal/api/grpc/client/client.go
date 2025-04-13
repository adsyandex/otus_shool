package client

import (
	"context"
	"time"

	pb "github.com/adsyandex/otus_shool/todo/internal/api/grpc/pb"
	"google.golang.org/grpc"
)

type TodoClient struct {
	client pb.TodoServiceClient
}

func NewTodoClient(conn *grpc.ClientConn) *TodoClient {
	return &TodoClient{
		client: pb.NewTodoServiceClient(conn),
	}
}

func (c *TodoClient) CreateTask(title, description string) (*pb.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.CreateTask(ctx, &pb.CreateTaskRequest{
		Title:       title,
		Description: description,
	})
	if err != nil {
		return nil, err
	}
	return resp.Task, nil
}

// Аналогично реализуйте остальные методы клиента