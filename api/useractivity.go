package api

import (
	db "Gym-backend/db/sqlc"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"time"
)

type createUserActivityRequest struct {
	Start    time.Time `json:"start" binding:"required"`
	End      time.Time `json:"end" binding:"required"`
	Userid   int64     `json:"userid" binding:"required"`
	Deviceid int64     `json:"deviceid" binding:"required"`
}

type getUserActivityRequest struct {
	Userid int64 `form:"userid" binding:"required"`
}

type userActivityResponse struct {
	ID       int64     `json:"id"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Userid   int64     `json:"userid"`
	Deviceid int64     `json:"deviceid"`
}

func newUserActivityResponse(activity db.Useractivity) userActivityResponse {
	return userActivityResponse{
		ID:       activity.ID,
		Start:    activity.Start,
		End:      activity.End,
		Userid:   activity.Userid,
		Deviceid: activity.Deviceid,
	}
}

// CreateTags		godoc
// @Summary			Create UserActivity
// @Description 	Create UserActivity data in Db.
// @Param 			device body createUserActivityRequest true "Create Device"
// @Produce 		application/json
// @Tags 			userActivity
// @Success 		200 {object} userActivityResponse{}
// @Router			/userActivity [post]
func (server *Server) createUserActivity(ctx *gin.Context) {
	var req createUserActivityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	arg := db.CreateUserActivityParams{
		Start:    req.Start,
		End:      req.End,
		Userid:   req.Userid,
		Deviceid: req.Deviceid,
	}

	activity, err := server.store.CreateUserActivity(ctx, arg)
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

	rsp := newUserActivityResponse(activity)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Get User Activity From ID
// @Description 	Get User Activity data from Db.
// @Param 			users query getUserActivityRequest true "Get user"
// @Produce 		application/json
// @Tags 			userActivity
// @Success 		200 {object} []userActivityResponse{}
// @Router			/userActivity [get]
func (server *Server) getUserActivity(ctx *gin.Context) {
	var req getUserActivityRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	activity, err := server.store.GetUserActivity(ctx, req.Userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//rsp := newUserActivityResponse(activity)
	ctx.JSON(http.StatusOK, activity)
}
