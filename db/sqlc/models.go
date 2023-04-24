// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type Checkinactivity struct {
	ID         int64     `json:"id"`
	Checkin    time.Time `json:"checkin"`
	Checkout   time.Time `json:"checkout"`
	Userid     int64     `json:"userid"`
	Employeeid int64     `json:"employeeid"`
	Locationid int64     `json:"locationid"`
}

type Class struct {
	ID             int64          `json:"id"`
	Instructorname string         `json:"instructorname"`
	Regstatus      string `json:"regstatus"`
	Startdate      time.Time      `json:"startdate"`
	Enddate        time.Time      `json:"enddate"`
	Starttime      time.Time      `json:"starttime"`
	Endtime        time.Time      `json:"endtime"`
	Day            string         `json:"day"`
	Name           string `json:"name"`
	// weekly daily or monthly
	Classtype  string `json:"classtype"`
	Locationid int64          `json:"locationid"`
	Cost       int32          `json:"cost"`
}

type Classcatalogue struct {
	ID       int64 `json:"id"`
	Userid   int64 `json:"userid"`
	Courseid int64 `json:"courseid"`
}

type Device struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	// Free,busy,not working
	Status string `json:"status"`
}

type Employee struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Hashedpassword string    `json:"hashedpassword"`
	Locationid     int64     `json:"locationid"`
	CreatedAt      time.Time `json:"created_at"`
}

type Location struct {
	ID      int64  `json:"id"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
}

type Membership struct {
	ID     int64 `json:"id"`
	Userid int64 `json:"userid"`
	// 0 is admin 1 is member 2 is non member
	MemberType int32        `json:"member_type"`
	ExpiryDate sql.NullTime `json:"expiry_date"`
}

type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Hashedpassword string    `json:"hashedpassword"`
	CreatedAt      time.Time `json:"created_at"`
}

type Useractivity struct {
	ID         int64     `json:"id"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	Userid     int64     `json:"userid"`
	Deviceid   int64     `json:"deviceid"`
	Locationid int64     `json:"locationid"`
}
