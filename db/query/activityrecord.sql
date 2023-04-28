-- name: CreateActivityRecords :one
insert INTO activityrecords
(type,userid,locationid )
VALUES
    ($1,$2,$3) RETURNING *;


-- name: GetLatestActivityRecord :one
SELECT * FROM activityrecords
WHERE userid = $1
ORDER BY time desc
    LIMIT 1;