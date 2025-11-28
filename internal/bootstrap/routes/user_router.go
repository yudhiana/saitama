package routes

import (
	"saitama/internal/bootstrap/engine"
	"saitama/modules/users/delivery/api/rest"
)

func init() {
	application := engine.GetEngine()
	restApplication := rest.NewRestApplication()
	group := application.Group("/users")
	group.POST("/", restApplication.CreateUser())
}
