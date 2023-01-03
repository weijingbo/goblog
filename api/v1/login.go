package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/middeware"
	"go-blog/model"
	"go-blog/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var token string
	var code int
	_ = c.ShouldBindJSON(&data)

	data, code = model.CheckLogin(data.Username, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middeware.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
