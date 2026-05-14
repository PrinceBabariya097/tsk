-- name: GetAllTodos :many
SELECT * FROM todos
WHERE (CAST(sqlc.arg(sort_by) AS TEXT) IS NOT NULL OR CAST(sqlc.arg(sort_by) AS TEXT) IS NULL)
ORDER BY 
  CASE WHEN CAST(?1 AS TEXT) = 'priority' THEN priority END DESC,
  CASE WHEN CAST(?1 AS TEXT) = 'created_at' THEN created_at END DESC,
  id DESC
LIMIT sqlc.arg(limit) 
OFFSET sqlc.arg(offset);

-- name: GetPendingTodos :many
SELECT * FROM todos
WHERE
 CASE WHEN sqlc.arg(completed) = 1 THEN completed = 1
      WHEN sqlc.arg(pending) = 1 THEN completed = 0
      ELSE 1=1 END 
 AND (CAST(sqlc.arg(sort_by) AS TEXT) IS NOT NULL OR CAST(sqlc.arg(sort_by) AS TEXT) IS NULL)
ORDER BY 
  CASE WHEN CAST(?3 AS TEXT) = 'priority' THEN priority END DESC,
  CASE WHEN CAST(?3 AS TEXT) = 'created_at' THEN created_at END DESC,
  id DESC
LIMIT sqlc.arg(limit) 
OFFSET sqlc.arg(offset);

-- name: GetTodoByPriority :many
SELECT * FROM todos
WHERE
  CASE WHEN sqlc.arg(priority) IS NOT NULL THEN priority = sqlc.arg(priority)
        ELSE 1=1 END
  AND (CAST(sqlc.arg(sort_by) AS TEXT) IS NOT NULL OR CAST(sqlc.arg(sort_by) AS TEXT) IS NULL)
ORDER BY 
  CASE WHEN CAST(?2 AS TEXT) = 'priority' THEN priority END DESC,
  CASE WHEN CAST(?2 AS TEXT) = 'created_at' THEN created_at END DESC,
  id DESC
LIMIT sqlc.arg(limit) 
OFFSET sqlc.arg(offset);

-- name: GetTodosByTag :many
SELECT * FROM todos
WHERE
  CASE WHEN sqlc.arg(tag) IS NOT NULL THEN tag = sqlc.arg(tag)
        ELSE 1=1 END
  AND (CAST(sqlc.arg(sort_by) AS TEXT) IS NOT NULL OR CAST(sqlc.arg(sort_by) AS TEXT) IS NULL)
ORDER BY 
  CASE WHEN CAST(?2 AS TEXT) = 'priority' THEN priority END DESC,
  CASE WHEN CAST(?2 AS TEXT) = 'created_at' THEN created_at END DESC,
  id DESC
LIMIT sqlc.arg(limit) 
OFFSET sqlc.arg(offset);


-- name: GetTodosByFiltersById :many
WITH config AS (SELECT CAST(sqlc.arg(is_desc) AS INTEGER) AS is_desc)
SELECT * FROM todos 
WHERE 
  -- Filter by completion status
  CASE 
    WHEN sqlc.arg(all) = 1 THEN TRUE
    WHEN sqlc.arg(completed) = 1 THEN completed = 1
    WHEN sqlc.arg(pending) = 1 THEN completed = 0
    ELSE TRUE
  END
  -- Filter by priority
  AND CASE 
    WHEN sqlc.arg(priority) IS NOT NULL THEN priority = sqlc.arg(priority)
    ELSE TRUE
  END
  -- Filter by tag
  AND CASE 
    WHEN sqlc.arg(tag) IS NOT NULL THEN tag = sqlc.arg(tag)
    ELSE TRUE
  END
ORDER BY 
  CASE WHEN (SELECT is_desc FROM config) = 1 THEN id END DESC,
  CASE WHEN (SELECT is_desc FROM config) = 0 THEN id END ASC,
  id DESC
LIMIT sqlc.arg(limit) 
OFFSET sqlc.arg(offset);

-- name: GetTodosByFiltersByPriority :many
WITH config AS (SELECT CAST(sqlc.arg(is_desc) AS INTEGER) AS is_desc)
SELECT * FROM todos 
WHERE 
  -- Filter by completion status
  CASE 
    WHEN sqlc.arg(all) = 1 THEN TRUE
    WHEN sqlc.arg(completed) = 1 THEN completed = 1
    WHEN sqlc.arg(pending) = 1 THEN completed = 0
    ELSE TRUE
  END
  -- Filter by priority
  AND CASE 
    WHEN sqlc.arg(priority) IS NOT NULL THEN priority = sqlc.arg(priority)
    ELSE TRUE
  END
  -- Filter by tag
  AND CASE 
    WHEN sqlc.arg(tag) IS NOT NULL THEN tag = sqlc.arg(tag)
    ELSE TRUE
  END
ORDER BY 
  CASE WHEN (SELECT is_desc FROM config) = 1 THEN priority END DESC,
  CASE WHEN (SELECT is_desc FROM config) = 0 THEN priority END ASC,
  id DESC
LIMIT sqlc.arg(limit) 
OFFSET sqlc.arg(offset);


-- name: GetTodosByFiltersByCreatedAt :many
WITH config AS (SELECT CAST(sqlc.arg(is_desc) AS INTEGER) AS is_desc)
SELECT * FROM todos 
WHERE 
  -- Filter by completion status
  CASE 
    WHEN sqlc.arg(all) = 1 THEN TRUE
    WHEN sqlc.arg(completed) = 1 THEN completed = 1
    WHEN sqlc.arg(pending) = 1 THEN completed = 0
    ELSE TRUE
  END
  -- Filter by priority
  AND CASE 
    WHEN sqlc.arg(priority) IS NOT NULL THEN priority = sqlc.arg(priority)
    ELSE TRUE
  END
  -- Filter by tag
  AND CASE 
    WHEN sqlc.arg(tag) IS NOT NULL THEN tag = sqlc.arg(tag)
    ELSE TRUE
  END
ORDER BY 
  CASE WHEN (SELECT is_desc FROM config) = 1 THEN created_at END DESC,
  CASE WHEN (SELECT is_desc FROM config) = 0 THEN created_at END ASC,
  id DESC
LIMIT sqlc.arg(limit) 
OFFSET sqlc.arg(offset);

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
INSERT INTO todos (name,  completed, priority, tag, completed_at, created_at, updated_at)
VALUES (:name, :completed, :priority, :tag, :completed_at, :created_at, :updated_at)
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
