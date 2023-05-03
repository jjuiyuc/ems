package routers

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// CasbinQuery godoc
type CasbinQuery struct {
	CasbinSubject string `form:"casbin_subject" binding:"required"`
}

func (w *APIWorker) getFrontendPermission(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{c}
		var query CasbinQuery
		if err := c.BindQuery(&query); err != nil {
			appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
			return
		}

		responseData, err := casbin.CasbinJsGetPermissionForUser(enforcer, query.CasbinSubject)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ErrAuthFrontendPermissionGen, nil)
			return
		}
		appG.Response(http.StatusOK, e.Success, responseData)
	}
}
