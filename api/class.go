package api

import (
	db "Gym-backend/db/sqlc"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type createClassRequest struct {
	Instructorname string    `json:"instructorname" binding:"required"`
	Startdate      time.Time `json:"startdate" binding:"required"`
	Enddate        time.Time `json:"enddate" binding:"required"`
	Starttime      time.Time `json:"starttime" binding:"required"`
	Endtime        time.Time `json:"endtime" binding:"required"`
	Day            string    `json:"day" binding:"required"`
	Name           string    `json:"name" binding:"required"`
	Locationid     int64     `json:"locationid" binding:"required"`
	Cost           int32     `json:"cost" binding:"required"`
}

type getClassFromIDRequest struct {
	Classid int64 `form:"classid" binding:"required"`
}

type getClassFromNameRequest struct {
	Classname string `form:"classname" binding:"required"`
}

type classResponse struct {
	ID             int64     `json:"id"`
	Instructorname string    `json:"instructorname"`
	Regstatus      string    `json:"regstatus"`
	Startdate      time.Time `json:"startdate"`
	Enddate        time.Time `json:"enddate"`
	Starttime      time.Time `json:"starttime"`
	Endtime        time.Time `json:"endtime"`
	Day            string    `json:"day"`
	Name           string    `json:"name"`
	// weekly daily or monthly
	Classtype  string `json:"classtype"`
	Locationid int64  `json:"locationid"`
	Cost       int32  `json:"cost"`
}

func newClassResponse(class db.Class) classResponse {
	return classResponse{
		ID:             class.ID,
		Instructorname: class.Instructorname,
		Regstatus:      class.Regstatus,
		Endtime:        class.Endtime,
		Starttime:      class.Starttime,
		Name:           class.Name,
		Classtype:      class.Classtype,
		Locationid:     class.Locationid,
		Startdate:      class.Startdate,
		Enddate:        class.Enddate,
		Day:            class.Day,
		Cost:           class.Cost,
	}
}

// CreateTags		godoc
// @Summary			Create Class
// @Description 	Create Class data in Db.
// @Param 			class body createClassRequest true "Create class"
// @Produce 		application/json
// @Tags 			class
// @Success 		200 {object} classResponse{}
// @Router			/class [post]
func (server *Server) createClass(ctx *gin.Context) {
	var req createClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Invalidadata")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateClassParams{
		Instructorname: req.Instructorname,
		Endtime:        req.Endtime,
		Starttime:      req.Starttime,
		Name:           req.Name,
		Locationid:     req.Locationid,
		Startdate:      req.Startdate,
		Enddate:        req.Enddate,
		Day:            req.Day,
		Cost:           req.Cost,
	}

	class, err := server.store.CreateClass(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newClassResponse(class)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			get Class
// @Description 	get Class data in Db.
// @Param 			class query getClassFromIDRequest true "get class"
// @Produce 		application/json
// @Tags 			class
// @Success 		200 {object} userResponse{}
// @Router			/class [get]
func (server *Server) getClass(ctx *gin.Context) {
	var req getClassFromIDRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	class, err := server.store.GetClass(ctx, req.Classid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newClassResponse(class)
	ctx.JSON(http.StatusOK, rsp)
}
