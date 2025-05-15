package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro-services-roadmap/atom-kit/kg"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Swag(r *gin.Engine) {
	// docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix

	// Swagger documentation route
	r.GET(kg.C.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Redirect /swagger to /swagger/index.html
	swagRedirectHandler := func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, kg.C.System.RouterPrefix+"/swagger/index.html")
	}
	r.GET(kg.C.System.RouterPrefix+"/swagger", swagRedirectHandler)
	r.GET(kg.C.System.RouterPrefix+"/api", swagRedirectHandler)
	r.GET(kg.C.System.RouterPrefix+"/", swagRedirectHandler)
}
