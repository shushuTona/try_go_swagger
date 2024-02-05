-- name: GetAllTaskList :many
SELECT * FROM tasks;

-- name: CreateTask :execresult
INSERT INTO tasks (`status`, title, `desc`) VALUES (?, ?, ?);

-- name: DeleteAuthor :exec
-- DELETE FROM authors
-- WHERE id = ?;
