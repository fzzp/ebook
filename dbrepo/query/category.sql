-- name: CreateCategory :execresult
INSERT INTO category SET category_name=?;

-- name: GetAllCategory :many
SELECT * FROM category;
