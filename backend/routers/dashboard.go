package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

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

	conn, err := w.upgrade(c.Writer, c.Request)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}
	client := &Client{
		ID:          userID.(int),
		Token:       token.(string),
		GatewayUUID: gatewayUUID,
		Conn:        conn,
		Pool:        pool,
	}

	pool.Register <- client
	client.run(w)
}

func (w *APIWorker) upgrade(writer http.ResponseWriter, request *http.Request) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
		Subprotocols:    []string{request.Header.Get("Sec-WebSocket-Protocol")},
	}

	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "upGrader.Upgrade",
			"err":       err,
		}).Error()
		return nil, err
	}
	return conn, nil
}
