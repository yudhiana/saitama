package rest

import (
	"saitama/modules/users/models/input"

	"github.com/gin-gonic/gin"
)

func (r *module) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user input.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := r.uc.CreateUser(c.Request.Context(), user.Name, user.Email)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "User created successfully",
			"user_id": result,
		})
	}
}
