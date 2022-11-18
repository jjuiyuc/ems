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
// @title DER-EMS API
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
	wsGroup := r.Group("/ws")

	// Auth
	apiGroup.POST("/auth", w.GetAuth)

	// User
	apiGroup.PUT("/users/password/lost", w.PasswordLost)
	apiGroup.PUT("/users/password/reset-by-token", w.PasswordResetByToken)
	apiGroup.GET("/users/profile", authorize(REST), w.GetProfile)

	// Analysis
	apiGroup.GET("/:gwid/devices/energy-distribution-info", authorize(REST), w.GetEnergyDistributionInfo)
	apiGroup.GET("/:gwid/devices/power-state", authorize(REST), w.GetPowerState)
	apiGroup.GET("/:gwid/devices/accumulated-power-state", authorize(REST), w.GetAccumulatedPowerState)
	apiGroup.GET("/:gwid/devices/power-self-supply-rate", authorize(REST), w.GetPowerSelfSupplyRate)

	// Time of Use
	apiGroup.GET("/:gwid/devices/battery/usage-info", authorize(REST), w.GetBatteryUsageInfo)
	apiGroup.GET("/:gwid/devices/time-of-use-info", authorize(REST), w.GetTimeOfUseInfo)
	apiGroup.GET("/:gwid/devices/solar/energy-usage", authorize(REST), w.GetSolarEnergyUsage)

	// Demand Charge
	apiGroup.GET("/:gwid/devices/charge-info", authorize(REST), w.GetChargeInfo)
	apiGroup.GET("/:gwid/devices/demand-state", authorize(REST), w.GetDemandState)

	// Energy Resources - Solar tab
	apiGroup.GET("/:gwid/devices/solar/energy-info", authorize(REST), w.GetSolarEnergyInfo)
	apiGroup.GET("/:gwid/devices/solar/power-state", authorize(REST), w.GetSolarPowerState)
	// Energy Resources - Battery tab
	apiGroup.GET("/:gwid/devices/battery/energy-info", authorize(REST), w.GetBatteryEnergyInfo)
	apiGroup.GET("/:gwid/devices/battery/power-state", authorize(REST), w.GetBatteryPowerState)
	apiGroup.GET("/:gwid/devices/battery/charge-voltage-state", authorize(REST), w.GetBatteryChargeVoltageState)
	// Energy Resources - Grid tab
	apiGroup.GET("/:gwid/devices/grid/energy-info", authorize(REST), w.GetGridEnergyInfo)
	apiGroup.GET("/:gwid/devices/grid/power-state", authorize(REST), w.GetGridPowerState)

	// Dashboard
	wsGroup.GET("/:gwid/devices/energy-info", authorize(WebSocket), w.dashboardHandler)

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
