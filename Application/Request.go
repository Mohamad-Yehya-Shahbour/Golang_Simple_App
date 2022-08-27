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

func (req *Request) Share() {}

// handle request data
func Req() func(c *gin.Context) *Request {
	return func(c *gin.Context) *Request {
		var request Request
		request.Context = c
		connectToDataBase(&request)
		SetLang(&request)
		return &request
	}
}

// init new request
func NewRequest(c *gin.Context) *Request {
	request := Req()
	return request(c)
}

func SetLang(req *Request){
	lang := gotrans.DetectLanguage(req.Context.GetHeader("Accept-Language"))
	gotrans.SetDefaultLocale(lang)
	req.Lang = lang
}

func NewRequestWithAuth(c *gin.Context) *Request {
	return NewRequest(c).Auth()
}
