package routers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/casbin/casbin/v2"
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

type (
	// APIType godoc
	APIType int
	// PolicyWebpageObject godoc
	PolicyWebpageObject int
	// PolicyWebpageAction godoc
	PolicyWebpageAction int
)

const (
	// REST godoc
	REST APIType = iota
	// WebSocket godoc
	WebSocket
)

const (
	// Dashboard godoc
	Dashboard PolicyWebpageObject = iota
	// Analysis godoc
	Analysis
	// TimeOfUseEnergy godoc
	TimeOfUseEnergy
	// Economics godoc
	Economics
	// EnergyResources godoc
	EnergyResources
	// DemandCharge godoc
	DemandCharge
	// FieldManagement godoc
	FieldManagement
	// AccountManagementGroup godoc
	AccountManagementGroup
	// AccountManagementUser godoc
	AccountManagementUser
	// Settings godoc
	Settings
	// AdvancedSettings godoc
	AdvancedSettings
)

const (
	// Create godoc
	Create PolicyWebpageAction = iota
	// Read godoc
	Read
	// Update godoc
	Update
	// Delete godoc
	Delete
)

func (o PolicyWebpageObject) String() string {
	switch o {
	case Dashboard:
		return "dashboard"
	case Analysis:
		return "analysis"
	case TimeOfUseEnergy:
		return "timeOfUseEnergy"
	case Economics:
		return "economics"
	case EnergyResources:
		return "energyResources"
	case DemandCharge:
		return "demandCharge"
	case FieldManagement:
		return "fieldManagement"
	case AccountManagementGroup:
		return "accountManagementGroup"
	case AccountManagementUser:
		return "accountManagementUser"
	case Settings:
		return "settings"
	case AdvancedSettings:
		return "advancedSettings"
	}
	return ""
}

func (a PolicyWebpageAction) String() string {
	switch a {
	case Create:
		return "create"
	case Read:
		return "read"
	case Update:
		return "update"
	case Delete:
		return "delete"
	}
	return ""
}

// APIWorker godoc
type APIWorker struct {
	Cfg      *viper.Viper
	Services *services.Services
}

// NewAPIWorker godoc
func NewAPIWorker(dir string, cfg *viper.Viper, services *services.Services) {
	w := &APIWorker{
		Cfg:      cfg,
		Services: services,
	}

	r := InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), initPolicy(dir), w)
	r.Run(cfg.GetString("server.port"))
}

// InitRouter godoc
// @title DER-EMS API
// @BasePath /api
func InitRouter(isCORS bool, ginMode string, enforcer *casbin.Enforcer, w *APIWorker) *gin.Engine {
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
	apiGroup.GET("/users/profile", authorizeJWT(REST), w.GetProfile)
	apiGroup.PUT("/users/name", authorizeJWT(REST), w.UpdateName)
	apiGroup.PUT("/users/password", authorizeJWT(REST), w.UpdatePassword)

	// Analysis
	apiGroup.GET("/:gwid/devices/energy-distribution-info", authorizeJWT(REST), authorizePolicy(Analysis.String(), Read.String(), enforcer), w.GetEnergyDistributionInfo)
	apiGroup.GET("/:gwid/devices/power-state", authorizeJWT(REST), authorizePolicy(Analysis.String(), Read.String(), enforcer), w.GetPowerState)
	apiGroup.GET("/:gwid/devices/accumulated-power-state", authorizeJWT(REST), authorizePolicy(Analysis.String(), Read.String(), enforcer), w.GetAccumulatedPowerState)
	apiGroup.GET("/:gwid/devices/power-self-supply-rate", authorizeJWT(REST), authorizePolicy(Analysis.String(), Read.String(), enforcer), w.GetPowerSelfSupplyRate)

	// Time of Use
	apiGroup.GET("/:gwid/devices/battery/usage-info", authorizeJWT(REST), authorizePolicy(TimeOfUseEnergy.String(), Read.String(), enforcer), w.GetBatteryUsageInfo)
	apiGroup.GET("/:gwid/devices/tou/info", authorizeJWT(REST), authorizePolicy(TimeOfUseEnergy.String(), Read.String(), enforcer), w.GetTimeOfUseInfo)
	apiGroup.GET("/:gwid/devices/solar/energy-usage", authorizeJWT(REST), authorizePolicy(TimeOfUseEnergy.String(), Read.String(), enforcer), w.GetSolarEnergyUsage)

	// Economics
	apiGroup.GET("/:gwid/devices/tou/energy-cost", authorizeJWT(REST), authorizePolicy(Economics.String(), Read.String(), enforcer), w.GetTimeOfUseEnergyCost)

	// Demand Charge
	apiGroup.GET("/:gwid/devices/charge-info", authorizeJWT(REST), authorizePolicy(DemandCharge.String(), Read.String(), enforcer), w.GetChargeInfo)
	apiGroup.GET("/:gwid/devices/demand-state", authorizeJWT(REST), authorizePolicy(DemandCharge.String(), Read.String(), enforcer), w.GetDemandState)

	// Energy Resources - Solar tab
	apiGroup.GET("/:gwid/devices/solar/energy-info", authorizeJWT(REST), authorizePolicy(EnergyResources.String(), Read.String(), enforcer), w.GetSolarEnergyInfo)
	apiGroup.GET("/:gwid/devices/solar/power-state", authorizeJWT(REST), authorizePolicy(EnergyResources.String(), Read.String(), enforcer), w.GetSolarPowerState)
	// Energy Resources - Battery tab
	apiGroup.GET("/:gwid/devices/battery/energy-info", authorizeJWT(REST), authorizePolicy(EnergyResources.String(), Read.String(), enforcer), w.GetBatteryEnergyInfo)
	apiGroup.GET("/:gwid/devices/battery/power-state", authorizeJWT(REST), authorizePolicy(EnergyResources.String(), Read.String(), enforcer), w.GetBatteryPowerState)
	apiGroup.GET("/:gwid/devices/battery/charge-voltage-state", authorizeJWT(REST), authorizePolicy(EnergyResources.String(), Read.String(), enforcer), w.GetBatteryChargeVoltageState)
	// Energy Resources - Grid tab
	apiGroup.GET("/:gwid/devices/grid/energy-info", authorizeJWT(REST), authorizePolicy(EnergyResources.String(), Read.String(), enforcer), w.GetGridEnergyInfo)
	apiGroup.GET("/:gwid/devices/grid/power-state", authorizeJWT(REST), authorizePolicy(EnergyResources.String(), Read.String(), enforcer), w.GetGridPowerState)

	// Dashboard
	wsGroup.GET("/:gwid/devices/energy-info", authorizeJWT(WebSocket), authorizePolicy(Dashboard.String(), Read.String(), enforcer), w.dashboardHandler)

	// Leap - webhook endpoint
	apiGroup.POST("/leap/bidding/dispatch/webhook", leapAuthorize(), w.GetLeapBiddingDispatch)

	return r
}

func authorizeJWT(apiType APIType) gin.HandlerFunc {
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
			log.WithField("caused-by", "no header").Error()
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
			log.WithField("caused-by", "invalid header").Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthInvalidHeader, nil)
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			// Token timeout included
			log.WithField("caused-by", "token parse").Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthTokenParse, nil)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("groupID", claims.GroupID)
		if apiType == WebSocket {
			c.Set("token", token)
		}

		c.Next()
	}
}

func initPolicy(dir string) (enforcer *casbin.Enforcer) {
	var err error
	enforcer, err = casbin.NewEnforcer(dir+"/rbac_model.conf", dir+"/rbac_policy.csv")
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "casbin.NewEnforcer",
			"err":       err,
		}).Panic()
	}
	return
}

func authorizePolicy(webpage string, action string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{c}

		groupID, _ := c.Get("groupID")
		if groupID == nil {
			log.WithField("caused-by", "error token").Error()
			appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
			c.Abort()
			return
		}

		err := enforcer.LoadPolicy()
		if err != nil {
			log.WithField("caused-by", "load policy").Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthPolicyLoad, nil)
			c.Abort()
			return
		}

		sub := strconv.FormatInt(groupID.(int64), 10)
		ok, err := enforcer.Enforce(sub, webpage, action)
		if !ok {
			log.WithField("caused-by", "access denied").Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthPolicyNotAllow, nil)
			c.Abort()
			return
		} else if err != nil {
			log.WithField("caused-by", "authorize policy").Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthPolicyCheck, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

func leapAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// XXX: Hardcode for demo
		const APIKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZXhwIjoxNjcxMTA4NTQzLCJpc3MiOiJkZXJlbXMifQ.NVwtEo5w8xLfxNGtF3bM8jT6OgG-oW-1JZEwj72ILHM"
		appG := app.Gin{c}

		token := c.GetHeader("x-api-key")
		if token != APIKey {
			log.WithFields(log.Fields{"caused-by": "invalid header"}).Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthInvalidHeader, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
