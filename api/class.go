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
	Instructorid int64     `json:"instructorid" binding:"required"`
	RegStatus    string    `json:"reg_status" binding:"required"`
	StartTime    time.Time `json:"start_time" binding:"required"`
	EndTime      time.Time `json:"end_time" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	// weekly daily or monthly
	Classtype  string `json:"classtype" binding:"required"`
	Locationid int64  `json:"locationid" binding:"required"`
}

type getClassFromIDRequest struct {
	Classid int64 `form:"classid" binding:"required"`
}

type getClassFromNameRequest struct {
	Classname string `form:"classname" binding:"required"`
}

type classResponse struct {
	ID           int64     `json:"id"`
	Instructorid int64     `json:"instructorid"`
	RegStatus    string    `json:"reg_status"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Description  string    `json:"description"`
	// weekly daily or monthly
	Classtype  string `json:"classtype"`
	Locationid int64  `json:"locationid"`
}

func newClassResponse(class db.Class) classResponse {
	return classResponse{
		ID:           class.ID,
		Instructorid: class.Instructorid,
		RegStatus:    class.RegStatus,
		EndTime:      class.EndTime,
		Description:  class.Description,
		Classtype:    class.Classtype,
		Locationid:   class.Locationid,
		StartTime:    class.StartTime,
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	arg := db.CreateClassParams{
		Instructorid: req.Instructorid,
		RegStatus:    req.RegStatus,
		EndTime:      req.EndTime,
		Description:  req.Description,
		Classtype:    req.Classtype,
		Locationid:   req.Locationid,
		StartTime:    req.StartTime,
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
