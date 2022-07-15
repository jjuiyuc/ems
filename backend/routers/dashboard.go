package routers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
	"der-ems/internal/utils"
)

func (w *APIWorker) websocketAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{c}

		authHeader := c.GetHeader("Sec-WebSocket-Protocol")
		if authHeader == "" {
			log.WithFields(log.Fields{"caused-by": "no header"}).Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthNoHeader, nil)
			c.Abort()
			return
		}

		if len(strings.Split(authHeader, " ")) != 1 {
			log.WithFields(log.Fields{"caused-by": "invalid header"}).Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthInvalidHeader, nil)
			c.Abort()
			return
		}

		token := authHeader
		claims, err := utils.ParseToken(token)
		if err != nil {
			// Token timeout included
			log.WithFields(log.Fields{"caused-by": "token parse"}).Error()
			appG.Response(http.StatusUnauthorized, e.ErrAuthTokenParse, nil)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("token", token)

		c.Next()
	}
}

func (w *APIWorker) dashboardHandler(c *gin.Context) {
	appG := app.Gin{c}

	userID, _ := c.Get("userID")
	if userID == nil {
		log.WithFields(log.Fields{"caused-by": "error token"}).Error()
		appG.Response(http.StatusUnauthorized, e.ErrToken, nil)
		return
	}
	token, _ := c.Get("token")

	gatewayUUID := c.Param("gwid")
	log.Debug("gatewayUUID: ", gatewayUUID)

	pool := newPool()
	go pool.start()

	conn, err := w.upgrade(c)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	client := &Client{
		ID:    userID.(int),
		Token: token.(string),
		Conn:  conn,
		Pool:  pool,
	}

	pool.Register <- client
	client.run(c)
}

func (w *APIWorker) upgrade(c *gin.Context) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
		Subprotocols:    []string{c.Request.Header.Get("Sec-WebSocket-Protocol")},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "upGrader.Upgrade",
			"err":       err,
		}).Error()
		return nil, err
	}
	return conn, nil
}
