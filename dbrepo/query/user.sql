-- name: CreateUser :execresult
INSERT INTO users SET email=?, password_hash=?, username=?, avatar=?;

-- name: UpdateUser :exec
UPDATE users SET
    password_hash = COALESCE(sqlc.narg(password_hash), password_hash),
    username = COALESCE(sqlc.narg(username), username),
    avatar = COALESCE(sqlc.narg(avatar), avatar)
WHERE 
    id = ?;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;