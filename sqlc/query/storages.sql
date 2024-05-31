/* storages */
-- name: FindStorageByOwnerID :one
SELECT id, owner_id FROM storages
WHERE owner_id = $1
;

-- name: CreateStorage :exec
INSERT INTO storages
  (id, owner_id)
VALUES
  ($1, $2)
;

-- name: DeleteStorage :exec
DELETE FROM storages
WHERE owner_id = $1
;

/* buckets */
-- name: FindBuckets :many
SELECT
  id,
  name,
  storage_id,
  (
    SELECT
        i.id
    FROM items AS i
    WHERE i.bucket_id = buckets.id
      AND lft = (SELECT MIN(lft) FROM items WHERE bucket_id = buckets.id)
  ) AS root_directory_id
FROM buckets
WHERE storage_id = $1
;

-- name: FindBucketByID :one
SELECT 
  id,
  name,
  storage_id,
  (
    SELECT
        i.id
    FROM items AS i
    WHERE i.bucket_id = buckets.id
      AND lft = (SELECT MIN(lft) FROM items WHERE bucket_id = buckets.id)
  ) AS root_directory_id
FROM buckets
WHERE buckets.id = $1
;

-- name: CreateBucket :exec
INSERT INTO buckets
  (id, name, storage_id)
VALUES
  ($1, $2, $3)
;

-- name: UpdateBucket :exec
UPDATE buckets
SET name = $1
WHERE id = $2
;

-- name: DeleteBucket :exec
DELETE FROM buckets
WHERE id = $1
;

/* items */
-- name: FindRootItemID :one
SELECT id
FROM items
WHERE bucket_id = $1 AND lft = 1
;

-- name: FindItemByID :one
SELECT id, name, is_file, bucket_id, lft, rgt
FROM items
WHERE id = $1
;

-- name: FindDirectChildren :many
SELECT c.id, c.name, c.is_file, c.bucket_id, c.lft, c.rgt
FROM items AS c
JOIN items AS p
ON p.id = $1 AND c.bucket_id = p.bucket_id AND p.lft < c.lft AND c.rgt < p.rgt
WHERE NOT EXISTS (
  SELECT 1
  FROM items AS m
  WHERE m.bucket_id = p.bucket_id
    AND m.lft BETWEEN p.lft AND p.rgt -- 親に含まれている
    AND c.lft BETWEEN m.lft AND m.rgt -- 子供がいる
    AND m.id NOT IN (c.id, p.id) -- 同じitemでない
)
;

-- name: FindItemsWithDepth :many
WITH dir_items AS (
  SELECT id, name, is_file, bucket_id, lft, rgt
  FROM items AS i
  WHERE EXISTS (
      SELECT 1 FROM items AS e
      WHERE e.id = @item_id
        AND e.bucket_id = i.bucket_id
        AND e.is_file = FALSE
        AND i.lft BETWEEN e.lft AND e.rgt
    )
)
SELECT
  c.id,
  c.name,
  c.is_file,
  c.bucket_id,
  c.lft, c.rgt,
  COALESCE((
    SELECT id
    FROM dir_items
    WHERE lft = MAX(CASE WHEN p.id = c.id THEN 0 ELSE p.lft END)
    ), ''
  )::VARCHAR(36) AS parent_id
FROM dir_items AS p
JOIN dir_items AS c ON c.lft BETWEEN p.lft AND p.rgt
GROUP BY c.id, c.name, c.is_file, c.bucket_id, c.lft, c.rgt
HAVING (COUNT(*) - 1) <= @depth
;

-- name: DeleteItemTree :exec
DELETE FROM items AS i
WHERE i.bucket_id = $1
  AND lft BETWEEN (SELECT l.lft FROM items AS l WHERE l.id = $2) AND (SELECT r.rgt FROM items AS r WHERE r.id = $2)
;

-- name: ShiftEmptyItemRange :exec
UPDATE items
SET lft = (CASE WHEN lft > sqlc.arg(lft) THEN lft - (sqlc.arg(rgt) - sqlc.arg(lft) + 1) ELSE lft END),
    rgt = (CASE WHEN rgt > sqlc.arg(rgt) THEN rgt - (sqlc.arg(rgt) - sqlc.arg(lft) + 1) ELSE rgt END)
WHERE bucket_id = @bucket_id AND (lft > sqlc.arg(lft) OR rgt > sqlc.arg(rgt))
;

-- name: SpreadItemRange :exec
UPDATE items
SET lft = (CASE WHEN lft > sqlc.arg(rgt) THEN lft + 2 * sqlc.arg(num) ELSE lft END),
    rgt = (CASE WHEN rgt > sqlc.arg(rgt) THEN rgt + 2 * sqlc.arg(num) ELSE rgt END)
WHERE bucket_id = sqlc.arg(bucket_id) AND rgt > sqlc.arg(rgt)
;

-- name: InsertItems :exec
INSERT INTO items
  (id, name, is_file, bucket_id, lft, rgt)
VALUES
  (
    UNNEST(@ids::VARCHAR(36)[]),
    UNNEST(@names::VARCHAR(256)[]),
    UNNEST(@is_files::BOOL[]),
    UNNEST(@bucket_ids::varchar(36)[]),
    UNNEST(@lfts::BIGINT[]),
    UNNEST(@rgts::BIGINT[])
  )
;

-- name: DeleteByBucketID :exec
DELETE FROM items
WHERE bucket_id = $1
;

/* contents */
-- name: FindLatestContent :one
SELECT id, item_id, path, file_type, created_at
FROM contents
ORDER BY created_at DESC
LIMIT 1
;

-- name: FindByItemID :many
SELECT id, item_id, path, file_type, created_at
FROM contents
WHERE item_id = $1
ORDER BY created_at DESC
;

-- name: FindLatestContentsByItemIDs :many
SELECT c.id, c.item_id, c.path, c.file_type, c.created_at
FROM contents AS c
WHERE c.item_id = ANY (@item_ids::VARCHAR(36)[])
  AND NOT EXISTS (
    SELECT 1
    FROM contents AS c2
    WHERE c.item_id = c2.item_id
      AND c.created_at < c2.created_at
  )
;

-- name: CreateContent :exec
INSERT INTO contents
  (id, item_id, path, file_type)
VALUES
  ($1, $2, $3, $4)
;

-- name: CreateContents :exec
INSERT INTO contents
  (id, item_id, path, file_type)
VALUES
  (
    UNNEST(@ids::VARCHAR(36)[]),
    UNNEST(@item_ids::VARCHAR(36)[]),
    UNNEST(@paths::VARCHAR(2100)[]),
    UNNEST(@file_types::VARCHAR(64)[])
  )
;
