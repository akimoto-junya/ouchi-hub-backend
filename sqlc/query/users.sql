-- name: CreateUser :exec
INSERT INTO users (id)
VALUES (?)
;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?
;
