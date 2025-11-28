package bootstrap

import (
	"saitama/internal/bootstrap/engine"
	_ "saitama/internal/bootstrap/routes"

	"github.com/gin-gonic/gin"
)

type Container struct {
	app *gin.Engine
}

func NewContainer() *Container {
	return &Container{
		app: engine.GetEngine(),
	}
}

func (c *Container) StartServer() error {
	return c.app.Run("0.0.0.0:8080")
}
