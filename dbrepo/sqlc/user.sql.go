// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package dbrepo

import (
	"context"
	"database/sql"
)

const CreateUser = `-- name: CreateUser :execresult
INSERT INTO users SET email=?, password_hash=?, username=?, avatar=?
`

type CreateUserParams struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Username     string `json:"username"`
	Avatar       string `json:"avatar"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, CreateUser,
		arg.Email,
		arg.PasswordHash,
		arg.Username,
		arg.Avatar,
	)
}

const GetUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password_hash, username, avatar, role, created_at, updated_at FROM users WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRowContext(ctx, GetUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Username,
		&i.Avatar,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const GetUserByID = `-- name: GetUserByID :one
SELECT id, email, password_hash, username, avatar, role, created_at, updated_at FROM users WHERE id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, id uint64) (*User, error) {
	row := q.db.QueryRowContext(ctx, GetUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Username,
		&i.Avatar,
		&i.Role,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const UpdateUser = `-- name: UpdateUser :exec
UPDATE users SET
    password_hash = COALESCE(?, password_hash),
    username = COALESCE(?, username),
    avatar = COALESCE(?, avatar)
WHERE 
    id = ?
`

type UpdateUserParams struct {
	PasswordHash sql.NullString `json:"password_hash"`
	Username     sql.NullString `json:"username"`
	Avatar       sql.NullString `json:"avatar"`
	ID           uint64         `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, UpdateUser,
		arg.PasswordHash,
		arg.Username,
		arg.Avatar,
		arg.ID,
	)
	return err
}