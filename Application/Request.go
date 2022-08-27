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
