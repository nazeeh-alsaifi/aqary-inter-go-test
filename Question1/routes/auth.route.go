package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nazeeh-alsaifi/aqary-inter-go-test/controllers"
)

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController}
}

func (rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {

	router := rg.Group("/users")
	router.POST("/", rc.authController.SignUpUser)
	router.POST("/generateotp", rc.authController.GenerateOtp)
	router.POST("/verifyotp", rc.authController.VerifyOtp)

}
