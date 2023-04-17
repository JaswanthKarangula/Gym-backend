-- name: CreateUserActivity :one
INSERT INTO useractivity
("start" ,"end" ,userid,deviceid)
VALUES
    ($1,$2,$3,$4) RETURNING *;


-- name: GetUserActivity :many
SELECT * FROM useractivity
WHERE userid = $1;

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