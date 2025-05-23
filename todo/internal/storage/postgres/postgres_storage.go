package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/adsyandex/otus_shool/todo/internal/models"
	"github.com/adsyandex/otus_shool/todo/internal/storage/contracts"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
	tx *sql.Tx // используется для транзакций
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	// Устанавливаем максимальное количество открытых соединений
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) CreateTask(ctx context.Context, task models.Task) (models.Task, error) {
	task.ID = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	query := `INSERT INTO tasks (id, title, description, completed, created_at, updated_at, due_date, priority) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := s.execContext(ctx, query,
		task.ID,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.UpdatedAt,
		task.DueDate,
		task.Priority,
	)

	if err != nil {
		return models.Task{}, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

func (s *PostgresStorage) GetTask(ctx context.Context, id string) (models.Task, error) {
	var task models.Task

	query := `SELECT id, title, description, completed, created_at, updated_at, due_date, priority 
	          FROM tasks WHERE id = $1`

	err := s.queryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.DueDate,
		&task.Priority,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Task{}, contracts.ErrTaskNotFound
		}
		return models.Task{}, fmt.Errorf("failed to get task: %w", err)
	}

	return task, nil
}

func (s *PostgresStorage) UpdateTask(ctx context.Context, task models.Task) (models.Task, error) {
	task.UpdatedAt = time.Now()

	query := `UPDATE tasks SET 
	          title = $1, 
	          description = $2, 
	          completed = $3, 
	          updated_at = $4, 
	          due_date = $5, 
	          priority = $6 
	          WHERE id = $7`

	result, err := s.execContext(ctx, query,
		task.Title,
		task.Description,
		task.Completed,
		task.UpdatedAt,
		task.DueDate,
		task.Priority,
		task.ID,
	)

	if err != nil {
		return models.Task{}, fmt.Errorf("failed to update task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return models.Task{}, contracts.ErrTaskNotFound
	}

	return task, nil
}

func (s *PostgresStorage) DeleteTask(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := s.execContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return contracts.ErrTaskNotFound
	}

	return nil
}

func (s *PostgresStorage) ListTasks(ctx context.Context, filter models.TaskFilter) ([]models.Task, error) {
	query := `SELECT id, title, description, completed, created_at, updated_at, due_date, priority 
	          FROM tasks WHERE 1=1`
	args := []interface{}{}

	if filter.Completed != nil {
		query += " AND completed = $1"
		args = append(args, *filter.Completed)
	}

	if filter.Priority != nil {
		query += fmt.Sprintf(" AND priority = $%d", len(args)+1)
		args = append(args, *filter.Priority)
	}

	query += " ORDER BY created_at DESC"

	rows, err := s.queryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.DueDate,
			&task.Priority,
		); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return tasks, nil
}

func (s *PostgresStorage) WithTransaction(ctx context.Context, fn func(tx contracts.Storage) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	txStorage := &PostgresStorage{db: s.db, tx: tx}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(txStorage); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %w", err, rbErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// Вспомогательные методы для работы с транзакциями
func (s *PostgresStorage) execContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if s.tx != nil {
		return s.tx.ExecContext(ctx, query, args...)
	}
	return s.db.ExecContext(ctx, query, args...)
}

func (s *PostgresStorage) queryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if s.tx != nil {
		return s.tx.QueryContext(ctx, query, args...)
	}
	return s.db.QueryContext(ctx, query, args...)
}

func (s *PostgresStorage) queryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if s.tx != nil {
		return s.tx.QueryRowContext(ctx, query, args...)
	}
	return s.db.QueryRowContext(ctx, query, args...)
}

func (s *PostgresStorage) Close(ctx context.Context) error {
	if s.db == nil {
		return nil
	}
	return s.db.Close()
}
