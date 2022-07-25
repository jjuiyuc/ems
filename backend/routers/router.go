package routers

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"der-ems/docs"
	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/services"
)

// APIType godoc
type APIType int

const (
	// REST godoc
	REST APIType = iota
	// WebSocket godoc
	WebSocket
)

// APIWorker godoc
type APIWorker struct {
	Services *services.Services
}

// NewAPIWorker godoc
func NewAPIWorker(cfg *viper.Viper, services *services.Services) {
	w := &APIWorker{Services: services}

	r := InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), w)
	r.Run(cfg.GetString("server.port"))
}

// InitRouter godoc
// @Title DER_EMS
// @BasePath /api
func InitRouter(isCORS bool, ginMode string, w *APIWorker) *gin.Engine {
	r := gin.New()
	if isCORS {
		r.Use(cors.New(cors.Config{
			AllowAllOrigins:        true,
			AllowCredentials:       true,
			AllowBrowserExtensions: true,
			AllowMethods:           []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
			AllowHeaders: []string{
				"Authorization", "Content-Type", "Upgrade", "Origin",
				"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		}))
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.MaxMultipartMemory = 256 << 20

	if ginMode == gin.DebugMode {
		docs.SwaggerInfo.BasePath = "/api"
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	apiGroup := r.Group("/api")

	// Auth
	apiGroup.POST("/auth", w.GetAuth)

	// User
	apiGroup.PUT("/users/password/lost", w.PasswordLost)
	apiGroup.PUT("/users/password/reset-by-token", w.PasswordResetByToken)
	apiGroup.GET("/users/profile", authorize(REST), w.GetProfile)

	// Dashboard
	apiGroup.GET("/:gwid/devices/energy-info", authorize(WebSocket), w.dashboardHandler)

	// Energy Resources - Battery tab
	apiGroup.GET("/:gwid/devices/battery/energy-info", authorize(), w.GetBatteryEnergyInfo)

	return r
}

func authorize(apiType APIType) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{c}

		var authHeader string
		switch apiType {
		case REST:
			authHeader = c.GetHeader("Authorization")
		case WebSocket:
			authHeader = c.GetHeader("Sec-WebSocket-Protocol")
		}
		if authHeader == "" {
			log.WithFields(log.Fields{"caused-by": "no header"}).Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthNoHeader, nil)
			c.Abort()
			return
		}

		var token string
		switch apiType {
		case REST:
			bearers := strings.Split(authHeader, " ")
			if len(bearers) == 2 && bearers[0] == "Bearer" {
				token = bearers[1]
			}
		case WebSocket:
			if len(strings.Split(authHeader, " ")) == 1 {
				token = authHeader
			}
		}
		if token == "" {
			log.WithFields(log.Fields{"caused-by": "invalid header"}).Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthInvalidHeader, nil)
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			// Token timeout included
			log.WithFields(log.Fields{"caused-by": "token parse"}).Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthTokenParse, nil)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		if apiType == WebSocket {
			c.Set("token", token)
		}

		c.Next()
	}
}
