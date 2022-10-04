package routers

import (
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// Client godoc
type Client struct {
	ID          int64
	Token       string
	GatewayUUID string
	Conn        *websocket.Conn
	Pool        *Pool
}

func (c *Client) run(w *APIWorker) {
	isOpen := true
	var previousLogTime time.Time

	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	go func() {
		for isOpen {
			if _, _, err := c.Conn.ReadMessage(); err != nil {
				isOpen = false
				log.WithFields(log.Fields{
					"caused-by": "c.Conn.ReadMessage",
					"err":       err,
				}).Error()
				break
			}
		}
	}()

	for isOpen {
		logTime, devicesEnergyInfo, err := w.Services.Devices.GetLatestDevicesEnergyInfo(c.GatewayUUID)
		if err != nil {
			isOpen = false
			response := app.Response{
				Code: e.ErrDashboardDataGen,
				Msg:  e.GetMsg(e.ErrDashboardDataGen),
				Data: err.Error(),
			}
			err = c.Conn.WriteJSON(response)
			if err != nil {
				log.WithFields(log.Fields{
					"caused-by": "c.Conn.WriteJSON",
					"err":       err,
				}).Error()
			}
			break
		}
		if previousLogTime == logTime {
			log.Debug("latest log time is the same : ", logTime)
			time.Sleep(time.Minute)
			continue
		}
		previousLogTime = logTime
		response := app.Response{
			Code: e.Success,
			Msg:  e.GetMsg(e.Success),
			Data: devicesEnergyInfo,
		}
		err = c.Conn.WriteJSON(response)
		if err != nil {
			isOpen = false
			log.WithFields(log.Fields{
				"caused-by": "c.Conn.WriteJSON",
				"err":       err,
			}).Error()
			break
		}
		time.Sleep(time.Minute)
	}
}
