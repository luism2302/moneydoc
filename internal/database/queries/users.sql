-- name: CreateNewUser :one
INSERT INTO users(username, email, created_at, updated_at) VALUES (
    ?,
    ?,
    DATETIME('now'),
    DATETIME('now')
)
RETURNING *;
