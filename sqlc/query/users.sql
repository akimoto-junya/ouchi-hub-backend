-- name: CreateUser :exec
INSERT INTO users
  (id)
VALUES
  ($1)
;

