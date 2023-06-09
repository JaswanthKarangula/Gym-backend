// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: checkinrecord.sql

package db

import (
	"context"
)

const createCheckinRecords = `-- name: CreateCheckinRecords :one
insert INTO checkinrecords
(type,userid,employeeid,locationid )
VALUES
    ($1,$2,$3,$4) RETURNING id, userid, type, time, employeeid, locationid
`

type CreateCheckinRecordsParams struct {
	Type       int32 `json:"type"`
	Userid     int64 `json:"userid"`
	Employeeid int64 `json:"employeeid"`
	Locationid int64 `json:"locationid"`
}

func (q *Queries) CreateCheckinRecords(ctx context.Context, arg CreateCheckinRecordsParams) (Checkinrecord, error) {
	row := q.db.QueryRowContext(ctx, createCheckinRecords,
		arg.Type,
		arg.Userid,
		arg.Employeeid,
		arg.Locationid,
	)
	var i Checkinrecord
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Type,
		&i.Time,
		&i.Employeeid,
		&i.Locationid,
	)
	return i, err
}

const getLatestCheckinRecord = `-- name: GetLatestCheckinRecord :one
SELECT id, userid, type, time, employeeid, locationid FROM checkinrecords
WHERE userid = $1
ORDER BY time desc
LIMIT 1
`

func (q *Queries) GetLatestCheckinRecord(ctx context.Context, userid int64) (Checkinrecord, error) {
	row := q.db.QueryRowContext(ctx, getLatestCheckinRecord, userid)
	var i Checkinrecord
	err := row.Scan(
		&i.ID,
		&i.Userid,
		&i.Type,
		&i.Time,
		&i.Employeeid,
		&i.Locationid,
	)
	return i, err
}
