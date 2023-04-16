package api

import (
	db "Gym-backend/db/sqlc"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    int64  `json:"email" binding:"required"` //,emai for email validation
}
type getUserFromIDRequest struct {
	UserId int64 `form:"userid" binding:"required"`
}

type getUserFromNameRequest struct {
	Username string `form:"username" binding:"required"`
}

type userResponse struct {
	Username  string    `json:"username"`
	Email     int64     `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:  user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// CreateTags		godoc
// @Summary			Create User
// @Description 	Create User data in Db.
// @Param 			users body createUserRequest true "Create user"
// @Produce 		application/json
// @Tags 			user
// @Success 		200 {object} userResponse{}
// @Router			/users [post]
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}

	arg := db.CreateUserParams{
		Name:           req.Username,
		Hashedpassword: req.Password,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
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

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

// CreateTags		godoc
// @Summary			Get User From UserName
// @Description 	Get User data from Db.
// @Param 			users query getUserFromNameRequest true "Get user"
// @Produce 		application/json
// @Tags 			user
// @Success 		200 {object} userResponse{}
// @Router			/users [get]
func (server *Server) getUser(ctx *gin.Context) {
	var req getUserFromNameRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		fmt.Println("Failed")
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	//username := ctx.Query("username")
	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

//type loginUserRequest struct {
//	Username string `json:"username" binding:"required,alphanum"`
//	Password string `json:"password" binding:"required,min=6"`
//}
//
//type loginUserResponse struct {
//	AccessToken string       `json:"access_token"`
//	User        userResponse `json:"user"`
//}

//
//func (server *Server) loginUser(ctx *gin.Context) {
//	var req loginUserRequest
//	if err := ctx.ShouldBindJSON(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	user, err := server.store.GetUser(ctx, req.Username)
//	if err != nil {
//		if err == sql.ErrNoRows {
//			ctx.JSON(http.StatusNotFound, errorResponse(err))
//			return
//		}
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	err = util.CheckPassword(req.Password, user.HashedPassword)
//	if err != nil {
//		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
//		return
//	}
//
//	accessToken, err := server.tokenMaker.CreateToken(
//		user.Username,
//		server.config.AccessTokenDuration,
//	)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	rsp := loginUserResponse{
//		AccessToken: accessToken,
//		User:        newUserResponse(user),
//	}
//	ctx.JSON(http.StatusOK, rsp)
//}
