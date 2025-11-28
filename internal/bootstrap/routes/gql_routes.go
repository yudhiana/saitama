package routes

import (
	"saitama/internal/bootstrap/engine"
	"saitama/modules/users/delivery/api/gql"
)

func init() {
	application := engine.GetEngine()
	group := application.Group("/gql")
	group.POST("/query", gql.GraphqlHandler())
	group.GET("/playground", gql.PlaygroundHandler())
}
