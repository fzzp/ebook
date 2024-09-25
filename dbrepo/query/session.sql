-- name: CreateSession :execresult
INSERT INTO sessions SET user_id=?, token_hash=?;

-- name: GetUserByToken :one
SELECT u.* FROM users u JOIN sessions s on u.id = s.user_id WHERE s.token_hash = ?;