package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
	"github.com/zignd/parsimonious-users-api/users"
)

type PageParams struct {
	Page int `form:"page,default=1" binding:"required,numeric,min=1"`
	Size int `form:"pageSize,default=15" binding:"required,numeric,min=1,max=15"`
}

type GetUsersByNameUriParams struct {
	Name string `uri:"name" binding:"required,max=120"`
}

type GetUsersByUsernameUriParams struct {
	Username string `uri:"username" binding:"required,max=60"`
}

func GetUsersByName(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pageParams PageParams
		if err := c.ShouldBindQuery(&pageParams); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid page query parameters",
			})
			return
		}
		var uriParams GetUsersByNameUriParams
		if err := c.ShouldBindUri(&uriParams); err != nil {
			c.JSON(400, gin.H{
				"error": "Name exceeds the maximum length of 120 characters",
			})
			return
		}

		usrs, err := users.GetUserByName(
			db,
			uriParams.Name,
			pageParams.Page,
			pageParams.Size,
		)
		if err != nil {
			log.Error(fmt.Errorf("failed to get users by name: %w", err))
			c.Status(500)
			c.Done()
			return
		}

		c.JSON(200, gin.H{
			"users": usrs,
		})
	}
}

func GetUsersByUsername(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pageParams PageParams
		if err := c.ShouldBindQuery(&pageParams); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid page query parameters",
			})
			return
		}
		var uriParams GetUsersByUsernameUriParams
		if err := c.ShouldBindUri(&uriParams); err != nil {
			c.JSON(400, gin.H{
				"error": "Username exceeds the maximum length of 60 characters",
			})
			return
		}

		usrs, err := users.GetUserByUsername(
			db,
			uriParams.Username,
			pageParams.Page,
			pageParams.Size,
		)
		if err != nil {
			log.Error(fmt.Errorf("failed to get users by username: %w", err))
			c.Status(500)
			c.Done()
			return
		}

		c.JSON(200, gin.H{
			"users": usrs,
		})
	}
}
