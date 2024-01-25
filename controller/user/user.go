package user

import (
	"FlutterBackend/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {
	var users []orm.Tbl_User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": users})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user orm.Tbl_User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "user": user})
}
