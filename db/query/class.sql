-- name: CreateClass :one
INSERT INTO class (
    instructorname,
    name,
    cost,
    scheduleid
)
VALUES
    ($1, $2, $3,$4) RETURNING *;


-- name: GetClass :one
SELECT * FROM class
WHERE id = $1 LIMIT 1;

-- name: GetAllClasses :one
SELECT * FROM class;


-- name: GetClasses :many
SELECT
    c.id AS class_id,
    c.name AS class_name,
    c.instructorname,
    c.cost,
    s.startdate,
    s.enddate,
    s.starttime,
    s.endtime,
    CASE
        WHEN cc.id IS NULL THEN 'Not Enrolled'
        ELSE 'Enrolled'
        END AS enrollment_status
FROM
    class c
        JOIN schedule s ON c.scheduleid = s.id
        LEFT JOIN classcatalogue cc ON c.id = cc.courseid AND cc.userid = $1
WHERE
        s.locationid = $2
  AND s.day = $3
ORDER BY
    s.starttime;



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