package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"

	db "github.com/AmaanShikalgar/ainyx-users-api/db/sqlc"
	"github.com/AmaanShikalgar/ainyx-users-api/internal/logger"
	"github.com/AmaanShikalgar/ainyx-users-api/internal/models"
	"github.com/AmaanShikalgar/ainyx-users-api/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error)
	GetAllUsers(ctx context.Context) ([]models.UserResponse, error)
	GetUserByID(ctx context.Context, id int32) (models.UserResponse, error)
	UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error)
	DeleteUser(ctx context.Context, id int32) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("invalid date format, use YYYY-MM-DD: %w", err)
	}

	user, err := s.repo.CreateUser(ctx, db.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		logger.Error("failed to create user",
			zap.String("name", req.Name),
			zap.Error(err),
		)
		return models.UserResponse{}, fmt.Errorf("failed to create user: %w", err)
	}

	logger.Info("user created",
		zap.Int32("id", user.ID),
		zap.String("name", user.Name),
	)

	return toUserResponse(user), nil
}

func (s *userService) GetAllUsers(ctx context.Context) ([]models.UserResponse, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		logger.Error("failed to get users", zap.Error(err))
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	logger.Info("fetched all users", zap.Int("count", len(users)))

	responses := make([]models.UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, toUserResponse(user))
	}

	return responses, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Info("user not found", zap.Int32("id", id))
			return models.UserResponse{}, fmt.Errorf("user with id %d not found", id)
		}
		logger.Error("failed to get user",
			zap.Int32("id", id),
			zap.Error(err),
		)
		return models.UserResponse{}, fmt.Errorf("failed to get user: %w", err)
	}

	return toUserResponse(user), nil
}

func (s *userService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("invalid date format, use YYYY-MM-DD: %w", err)
	}

	user, err := s.repo.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Info("user not found for update", zap.Int32("id", id))
			return models.UserResponse{}, fmt.Errorf("user with id %d not found", id)
		}
		logger.Error("failed to update user",
			zap.Int32("id", id),
			zap.Error(err),
		)
		return models.UserResponse{}, fmt.Errorf("failed to update user: %w", err)
	}

	logger.Info("user updated",
		zap.Int32("id", user.ID),
		zap.String("name", user.Name),
	)

	return toUserResponse(user), nil
}

func (s *userService) DeleteUser(ctx context.Context, id int32) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		logger.Error("failed to delete user",
			zap.Int32("id", id),
			zap.Error(err),
		)
		return fmt.Errorf("failed to delete user: %w", err)
	}

	logger.Info("user deleted", zap.Int32("id", id))
	return nil
}

func toUserResponse(user db.User) models.UserResponse {
	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Format("2006-01-02"),
		Age:  calculateAge(user.Dob),
	}
}

func calculateAge(dob time.Time) int {
	today := time.Now()
	age := today.Year() - dob.Year()

	birthdayThisYear := time.Date(
		today.Year(),
		dob.Month(),
		dob.Day(),
		0, 0, 0, 0,
		today.Location(),
	)

	if today.Before(birthdayThisYear) {
		age--
	}

	if age < 0 {
		return 0
	}

	return age
}
