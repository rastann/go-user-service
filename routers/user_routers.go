package routers

import (
	"net/http"
	"tutorial/user-service/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var userRoute UserRoute

type UserRoute struct {
	Db gorm.DB
}

func GetUserRoute(db *gorm.DB) UserRoute {
	userRoute = UserRoute{Db: *db}
	return userRoute
}

func (this *UserRoute) GetUsers(c *gin.Context) {
	var users []entity.User
	this.Db.Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}
