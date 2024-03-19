-- name: FindRootItemID :one
SELECT id
FROM items
WHERE work_id = $1 AND lft = 1
;

-- name: FindItemByID :one
SELECT id, lft, rgt
FROM items
WHERE id = $1
;

-- name: FindDirectChildren :many
SELECT c.id
FROM items AS c
JOIN items AS p
ON p.id = $1 AND c.work_id = p.work_id AND p.lft < c.lft AND c.rgt < p.rgt
WHERE NOT EXISTS (
  SELECT 1
  FROM items AS m
  WHERE m.work_id = p.work_id
    AND m.lft BETWEEN p.lft AND p.rgt -- 親に含まれている
    AND c.lft BETWEEN m.lft AND m.rgt -- 子供がいる
    AND m.id NOT IN (c.id, p.id) -- 同じitemでない
)
;

-- name: DeleteItemTree :exec
DELETE FROM items AS i
WHERE i.work_id = $1
  AND lft BETWEEN (SELECT l.lft FROM items AS l WHERE l.id = $2) AND (SELECT r.rgt FROM items AS r WHERE r.id = $2)
;

-- name: ShiftEmptyItemRange :exec
UPDATE items
SET lft = (CASE WHEN lft > sqlc.arg(lft) THEN lft - (sqlc.arg(rgt) - sqlc.arg(lft) + 1) ELSE lft END),
    rgt = (CASE WHEN rgt > sqlc.arg(rgt) THEN rgt - (sqlc.arg(rgt) - sqlc.arg(lft) + 1) ELSE rgt END)
WHERE work_id = @work_id AND (lft > sqlc.arg(lft) OR rgt > sqlc.arg(rgt))
;

-- name: SpreadItemRange :exec
UPDATE items
SET lft = (CASE WHEN lft > sqlc.arg(rgt) THEN lft + 2 * sqlc.arg(num) ELSE lft END),
    rgt = (CASE WHEN rgt > sqlc.arg(rgt) THEN rgt + 2 * sqlc.arg(num) ELSE rgt END)
WHERE work_id = sqlc.arg(work_id) AND rgt > sqlc.arg(rgt)
;

-- name: InsertItems :exec
INSERT INTO items
  (id, name, type_id, work_id,  lft, rgt)
VALUES
  (
    UNNEST(@ids::VARCHAR(64)[]),
    UNNEST(@names::VARCHAR(255)[]),
    UNNEST(@type_ids::VARCHAR(64)[]),
    UNNEST(@work_ids::varchar(64)[]),
    UNNEST(@lfts::BIGINT[]),
    UNNEST(@rgts::BIGINT[])
  )
;
