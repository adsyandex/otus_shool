package models

import (
	"time"
	pb "github.com/adsyandex/otus_shool/todo/internal/api/grpc/pb"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) ToProto() *pb.Task {
	return &pb.Task{
		Id:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		CreatedAt:   t.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   t.UpdatedAt.Format(time.RFC3339),
	}
}

func TaskFromProto(p *pb.Task) (*Task, error) {
	createdAt, err := time.Parse(time.RFC3339, p.CreatedAt)
	if err != nil {
		return nil, err
	}
	
	updatedAt, err := time.Parse(time.RFC3339, p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	
	return &Task{
		ID:          p.Id,
		Title:       p.Title,
		Description: p.Description,
		Completed:   p.Completed,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}