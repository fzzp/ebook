-- name: CreateBook :execresult
INSERT INTO books 
SET 
    book_title=?,
    book_isbn=?,
    book_cover=?,
    book_author=?,
    book_price=?,
    book_inventory=?,
    ebook_price=?,
    publish_at=?
;


-- name: UpdateBook :exec
UPDATE books 
SET
    book_title=COALESCE(sqlc.narg(book_title), book_title),
    book_isbn=COALESCE(sqlc.narg(book_isbn), book_isbn),
    book_cover=COALESCE(sqlc.narg(book_cover), book_cover),
    book_author=COALESCE(sqlc.narg(book_author), book_author),
    book_price=COALESCE(sqlc.narg(book_price), book_price),
    book_inventory=COALESCE(sqlc.narg(book_inventory), book_inventory),
    ebook_price=COALESCE(sqlc.narg(ebook_price), ebook_price),
    publish_at=COALESCE(sqlc.narg(publish_at), publish_at),
    version=version+1
WHERE 
    id = ? AND version = ?
;

-- name: GetBookByID :one
SELECT * FROM books WHERE id = ?;


-- name: GetBookList :many
SELECT * FROM books ORDER BY id DESC LIMIT ? OFFSET ?;


-- name: GetBooksByCategory :many
SELECT b.* FROM books b JOIN category_book cb ON b.id = cb.book_id WHERE cb.category_id=?;