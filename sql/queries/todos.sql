-- name: GetAllTodos :many
SELECT * FROM todos
ORDER BY created_at DESC;

-- name: GetTodoById :one
SELECT * FROM todos
WHERE id = ?;

-- name: UpdateTodo :exec
UPDATE todos
SET name = ?, completed = ?, completed_at = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;

-- name: MarkTodoComplete :exec
UPDATE todos
SET completed = 1, completed_at = ?, updated_at = ?
WHERE id = ?;

-- name: CreateTodo :one
INSERT INTO todos (name,  completed, completed_at, created_at, updated_at)
VALUES (:name, :completed, :completed_at, :created_at, :updated_at)
RETURNING *;

-- name: GetAllTodo :many
SELECT * FROM todos;

-- name: CompleteTodo :one
UPDATE todos
SET completed = 1, completed_at = ?, updated_at = ?
WHERE id = ?
RETURNING id;

-- name: GetTodoCountById :one
SELECT COUNT(*) AS found
FROM todos 
WHERE id = ?;
