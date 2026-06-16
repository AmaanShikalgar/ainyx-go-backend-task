package repository

import (
	"context"

	db "github.com/AmaanShikalgar/ainyx-users-api/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetAllUsers(ctx context.Context) ([]db.User, error)
	GetUserByID(ctx context.Context, id int32) (db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

type userRepository struct {
	queries *db.Queries
}

func NewUserRepository(queries *db.Queries) UserRepository {
	return &userRepository{queries: queries}
}

func (r *userRepository) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.queries.CreateUser(ctx, arg)
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]db.User, error) {
	return r.queries.GetAllUsers(ctx)
}

func (r *userRepository) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

func (r *userRepository) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return r.queries.UpdateUser(ctx, arg)
}

func (r *userRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}
