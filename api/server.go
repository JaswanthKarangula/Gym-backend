package api

import (
	db "Gym-backend/db/sqlc"
	"Gym-backend/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	store  db.Store
	//tokenMaker token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	//tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	//if err != nil {
	//	return nil, fmt.Errorf("cannot create token maker: %w", err)
	//}

	server := &Server{
		config: config,
		store:  store,
		//tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	// https://santoshk.dev/posts/2022/how-to-integrate-swagger-ui-in-go-backend-gin-edition/

	c.JSON(http.StatusOK, res)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/h", HealthCheck)

	router.GET("/users", server.getUser)
	router.POST("/users", server.createUser)

	router.GET("/employee", server.getEmployee)
	router.POST("/employee", server.createEmployee)

	router.GET("/location", server.getLocation)
	router.POST("/location", server.createLocation)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
