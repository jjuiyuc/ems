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
	PolicyWebpageObject string
	// PolicyWebpageAction godoc
	PolicyWebpageAction string
)

const (
	// REST godoc
	REST APIType = iota
	// WebSocket godoc
	WebSocket
)

const (
	// Dashboard godoc
	Dashboard PolicyWebpageObject = "dashboard"
	// Analysis godoc
	Analysis PolicyWebpageObject = "analysis"
	// TimeOfUseEnergy godoc
	TimeOfUseEnergy PolicyWebpageObject = "timeOfUseEnergy"
	// Economics godoc
	Economics PolicyWebpageObject = "economics"
	// DemandCharge godoc
	DemandCharge PolicyWebpageObject = "demandCharge"
	// EnergyResources godoc
	EnergyResources PolicyWebpageObject = "energyResources"
	// FieldManagement godoc
	FieldManagement PolicyWebpageObject = "fieldManagement"
	// AccountManagementGroup godoc
	AccountManagementGroup PolicyWebpageObject = "accountManagementGroup"
	// AccountManagementUser godoc
	AccountManagementUser PolicyWebpageObject = "accountManagementUser"
	// Settings godoc
	Settings PolicyWebpageObject = "settings"
	// AdvancedSettings godoc
	AdvancedSettings PolicyWebpageObject = "advancedSettings"
)

const (
	// Create godoc
	Create PolicyWebpageAction = "create"
	// Read godoc
	Read PolicyWebpageAction = "read"
	// Update godoc
	Update PolicyWebpageAction = "update"
	// Delete godoc
	Delete PolicyWebpageAction = "delete"
)

// EndpointMapping godoc
var EndpointMapping = map[PolicyWebpageObject][]string{
	Dashboard: {
		"/ws/:gwid/devices/energy-info",
	},
	Analysis: {
		"/api/:gwid/devices/energy-distribution-info",
		"/api/:gwid/devices/power-state",
		"/api/:gwid/devices/accumulated-power-state",
		"/api/:gwid/devices/power-self-supply-rate",
	},
	TimeOfUseEnergy: {
		"/api/:gwid/devices/battery/usage-info",
		"/api/:gwid/devices/tou/info",
		"/api/:gwid/devices/solar/energy-usage",
	},
	Economics: {
		"/api/:gwid/devices/tou/energy-cost",
	},
	DemandCharge: {
		"/api/:gwid/devices/charge-info",
		"/api/:gwid/devices/demand-state",
	},
	EnergyResources: {
		"/api/:gwid/devices/solar/energy-info",
		"/api/:gwid/devices/solar/power-state",
		"/api/:gwid/devices/battery/energy-info",
		"/api/:gwid/devices/battery/power-state",
		"/api/:gwid/devices/battery/charge-voltage-state",
		"/api/:gwid/devices/grid/energy-info",
		"/api/:gwid/devices/grid/power-state",
	},
	AccountManagementGroup: {
		"/api/account-management/groups",
		"/api/account-management/groups/:groupid",
	},
}

// MethodMapping godoc
var MethodMapping = map[PolicyWebpageAction]string{
	Create: "POST",
	Read:   "GET",
	Update: "PUT",
	Delete: "DELETE",
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

	// Auth
	apiGroup.POST("/auth", w.GetAuth)

	// User
	apiGroup.PUT("/users/password/lost", w.PasswordLost)
	apiGroup.PUT("/users/password/reset-by-token", w.PasswordResetByToken)
	apiGroup.GET("/users/profile", authorizeJWT(REST), w.GetProfile)
	apiGroup.PUT("/users/name", authorizeJWT(REST), w.UpdateName)
	apiGroup.PUT("/users/password", authorizeJWT(REST), w.UpdatePassword)

	// Dashboard
	r.GET(EndpointMapping[Dashboard][0], authorizeJWT(WebSocket), authorizePolicy(enforcer), w.dashboardHandler)

	// Analysis
	r.GET(EndpointMapping[Analysis][0], authorizeJWT(REST), authorizePolicy(enforcer), w.GetEnergyDistributionInfo)
	r.GET(EndpointMapping[Analysis][1], authorizeJWT(REST), authorizePolicy(enforcer), w.GetPowerState)
	r.GET(EndpointMapping[Analysis][2], authorizeJWT(REST), authorizePolicy(enforcer), w.GetAccumulatedPowerState)
	r.GET(EndpointMapping[Analysis][3], authorizeJWT(REST), authorizePolicy(enforcer), w.GetPowerSelfSupplyRate)

	// Time of Use
	r.GET(EndpointMapping[TimeOfUseEnergy][0], authorizeJWT(REST), authorizePolicy(enforcer), w.GetBatteryUsageInfo)
	r.GET(EndpointMapping[TimeOfUseEnergy][1], authorizeJWT(REST), authorizePolicy(enforcer), w.GetTimeOfUseInfo)
	r.GET(EndpointMapping[TimeOfUseEnergy][2], authorizeJWT(REST), authorizePolicy(enforcer), w.GetSolarEnergyUsage)

	// Economics
	r.GET(EndpointMapping[Economics][0], authorizeJWT(REST), authorizePolicy(enforcer), w.GetTimeOfUseEnergyCost)

	// Demand Charge
	r.GET(EndpointMapping[DemandCharge][0], authorizeJWT(REST), authorizePolicy(enforcer), w.GetChargeInfo)
	r.GET(EndpointMapping[DemandCharge][1], authorizeJWT(REST), authorizePolicy(enforcer), w.GetDemandState)

	// Energy Resources - Solar tab
	r.GET(EndpointMapping[EnergyResources][0], authorizeJWT(REST), authorizePolicy(enforcer), w.GetSolarEnergyInfo)
	r.GET(EndpointMapping[EnergyResources][1], authorizeJWT(REST), authorizePolicy(enforcer), w.GetSolarPowerState)
	// Energy Resources - Battery tab
	r.GET(EndpointMapping[EnergyResources][2], authorizeJWT(REST), authorizePolicy(enforcer), w.GetBatteryEnergyInfo)
	r.GET(EndpointMapping[EnergyResources][3], authorizeJWT(REST), authorizePolicy(enforcer), w.GetBatteryPowerState)
	r.GET(EndpointMapping[EnergyResources][4], authorizeJWT(REST), authorizePolicy(enforcer), w.GetBatteryChargeVoltageState)
	// Energy Resources - Grid tab
	r.GET(EndpointMapping[EnergyResources][5], authorizeJWT(REST), authorizePolicy(enforcer), w.GetGridEnergyInfo)
	r.GET(EndpointMapping[EnergyResources][6], authorizeJWT(REST), authorizePolicy(enforcer), w.GetGridPowerState)

	// Account Management Group
	r.GET(EndpointMapping[AccountManagementGroup][0], authorizeJWT(REST), authorizePolicy(enforcer), w.GetGroups)
	r.POST(EndpointMapping[AccountManagementGroup][0], authorizeJWT(REST), authorizePolicy(enforcer), w.CreateGroup)
	r.GET(EndpointMapping[AccountManagementGroup][1], authorizeJWT(REST), authorizePolicy(enforcer), w.GetGroup)
	r.PUT(EndpointMapping[AccountManagementGroup][1], authorizeJWT(REST), authorizePolicy(enforcer), w.UpdateGroup)
	r.DELETE(EndpointMapping[AccountManagementGroup][1], authorizeJWT(REST), authorizePolicy(enforcer), w.DeleteGroup)

	// Casbin route
	apiGroup.GET("/casbin", w.getFrontendPermission(enforcer))

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
		c.Set("groupType", claims.GroupType)
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

func authorizePolicy(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{c}

		groupType, _ := c.Get("groupType")
		if groupType == nil {
			log.WithField("caused-by", "error token").Error()
			appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
			c.Abort()
			return
		}

		err := enforcer.LoadPolicy()
		if err != nil {
			log.WithField("caused-by", "load policy").Error()
			appG.Response(http.StatusForbidden, e.ErrAuthPolicyLoad, nil)
			c.Abort()
			return
		}

		sub := strconv.FormatInt(groupType.(int64), 10)
		webpage := getWebpage(c.FullPath())
		action := getAction(c.Request.Method)
		ok, err := enforcer.Enforce(sub, string(webpage), string(action))
		if !ok {
			log.WithField("caused-by", "permission denied").Error()
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionNotAllow, nil)
			c.Abort()
			return
		} else if err != nil {
			log.WithField("caused-by", "check permission").Error()
			appG.Response(http.StatusForbidden, e.ErrAuthPermissionCheck, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

func getWebpage(path string) (webpage PolicyWebpageObject) {
	for key, endpoints := range EndpointMapping {
		for _, e := range endpoints {
			if e == path {
				webpage = key
				break
			}
		}
	}
	return
}

func getAction(method string) (action PolicyWebpageAction) {
	for key, m := range MethodMapping {
		if m == method {
			action = key
			break
		}
	}
	return
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
