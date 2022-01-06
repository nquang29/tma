package rest

import (
	"TMA/internal/tma_management/service"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetUser godoc
// @Summary User
// @Description Get User
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/get_user [get]
func GetUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		ginR := Response{c: c}
		//id := c.Param("id")
		user, err := service.Find("1234")
		if err != nil{
			ginR.Response(400, "ERROR", err)
			return
		}
		ginR.Response(200, "OK", user)
	}
}
