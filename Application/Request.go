package Application

import (
	"database/sql"
	"projects-template/Models"

	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type shareResources interface {
	Share()
}

type Request struct {
	Context    *gin.Context
	DB         *gorm.DB
	Connection *sql.DB
	User       *Models.User
	IsAuth     bool
	IsAdmin	   bool
	Lang 	   string
}
