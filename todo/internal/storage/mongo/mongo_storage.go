package mongo

import (
	"context"
	"errors"
	"fmt"
	
	"github.com/adsyandex/otus_shool/todo/internal/storage"
	"github.com/adsyandex/otus_shool/todo/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStorage(uri, dbName, collectionName string) (*MongoStorage, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return &MongoStorage{
		client:     client,
		collection: client.Database(dbName).Collection(collectionName),
	}, nil
}

func (s *MongoStorage) SaveTask(ctx context.Context, task models.Task) error {
	_, err := s.collection.InsertOne(ctx, task)
	return err
}

func (s *MongoStorage) GetTask(ctx context.Context, id string) (*models.Task, error) {
	var task models.Task
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, storage.ErrNotFound
	}
	return &task, err
}

func (s *MongoStorage) UpdateTask(ctx context.Context, task models.Task) error {
	res, err := s.collection.UpdateByID(ctx, task.ID, bson.M{"$set": task})
	if res.MatchedCount == 0 {
		return storage.ErrNotFound
	}
	return err
}

func (s *MongoStorage) DeleteTask(ctx context.Context, id string) error {
	res, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {                                        // Сначала проверяем err
        return fmt.Errorf("delete failed: %w", err)        // Оборачиваем исходную ошибку с контекстом (%w)
	}
	if res.DeletedCount == 0 {
		return storage.ErrNotFound
	}
	return nil
}

func (s *MongoStorage) ListTasks(ctx context.Context, completed *bool) ([]models.Task, error) {
	filter := bson.M{}
	if completed != nil {
		filter["done"] = *completed
	}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *MongoStorage) Close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
}