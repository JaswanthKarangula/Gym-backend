-- name: CreateClass :one
INSERT INTO class (
    instructorid,
    reg_status,
    start_time,
    end_time,
    description,
    classtype,
    locationid
)
VALUES
    ($1, $2, $3,$4, $5, $6,$7) RETURNING *;


-- name: GetClass :one
SELECT * FROM class
WHERE id = $1 LIMIT 1;

-- -- name: UpdateUser :one
-- UPDATE users
-- SET
--     hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
--     password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
--     full_name = COALESCE(sqlc.narg(full_name), full_name),
--     email = COALESCE(sqlc.narg(email), email)
-- WHERE
--         username = sqlc.arg(username)
--     RETURNING *;