package routes

import (
	"polarisApi/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {

	// DISCORD BOT REQUESTS
	e.POST("/users/:discordId", controllers.RegisterUser)
	e.POST("/projects/:serverId", controllers.RegisterProject)
	e.POST("/orders/", controllers.NewOrder)
	e.PUT("/users/newKey/:discordId/:keyHash", controllers.PutNewKeyHash)
	e.GET("/users/:discordId", controllers.GetUser)
	e.GET("/projects", controllers.GetProjects)
	e.GET("/projects/:serverId", controllers.GetProjectByServerId)
	e.GET("/", controllers.StandardResponse)

	// CHAIN LISTENER, EXECUTOR AND VALIDATOR REQUESTS
	e.PUT("/users/:keyHash/:chain", controllers.LinkWalletToUser)
	e.PUT("/orders/:orderId/:validatorId", controllers.UpdatedOrder)
	e.GET("/orders/:chain", controllers.GetOrders)

}
