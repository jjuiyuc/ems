package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"der-ems/internal/app"
	"der-ems/internal/e"
)

// Client godoc
type Client struct {
	ID    int
	Token string
	Conn  *websocket.Conn
	Pool  *Pool
}

// DashboardData godoc
type DashboardData struct {
	GridIsPeakShaving             int            `json:"gridIsPeakShaving"`
	LoadGridAveragePowerAC        float32        `json:"loadGridAveragePowerAC"`
	BatteryGridAveragePowerAC     float32        `json:"batteryGridAveragePowerAC"`
	GridContractPowerAC           float32        `json:"gridContractPowerAC"`
	LoadPvAveragePowerAC          float32        `json:"loadPvAveragePowerAC"`
	LoadBatteryAveragePowerAC     float32        `json:"loadBatteryAveragePowerAC"`
	BatterySoC                    float32        `json:"batterySoC"`
	BatteryProducedAveragePowerAC float32        `json:"batteryProducedAveragePowerAC"`
	BatteryConsumedAveragePowerAC float32        `json:"batteryConsumedAveragePowerAC"`
	BatteryChargingFrom           string         `json:"batteryChargingFrom"`
	BatteryDischargingTo          string         `json:"batteryDischargingTo"`
	PvAveragePowerAC              float32        `json:"pvAveragePowerAC"`
	LoadAveragePowerAC            float32        `json:"loadAveragePowerAC"`
	LoadLinks                     map[string]int `json:"loadLinks"`
	GridLinks                     map[string]int `json:"gridLinks"`
	PvLinks                       map[string]int `json:"pvLinks"`
	BatteryLinks                  map[string]int `json:"batteryLinks"`
	BatteryPvAveragePowerAC       float32        `json:"batteryPvAveragePowerAC"`
	GridPvAveragePowerAC          float32        `json:"gridPvAveragePowerAC"`
	GridProducedAveragePowerAC    float32        `json:"gridProducedAveragePowerAC"`
	GridConsumedAveragePowerAC    float32        `json:"gridConsumedAveragePowerAC"`
}

func (c *Client) run(ctx *gin.Context) {
	isOpen := true

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
		// TODO: implement to get dashboard data, not hardcode string
		loadLinksValue := map[string]int{
			"grid":    0,
			"battery": 0,
			"pv":      0,
		}
		gridLinksValue := map[string]int{
			"load":    1,
			"battery": 0,
			"pv":      0,
		}
		pvLinksValue := map[string]int{
			"load":    1,
			"battery": 1,
			"grid":    0,
		}
		batteryLinksValue := map[string]int{
			"load": 0,
			"pv":   0,
			"grid": 0,
		}
		dashboardData := DashboardData{
			GridIsPeakShaving:             0,
			LoadGridAveragePowerAC:        10,
			BatteryGridAveragePowerAC:     0,
			GridContractPowerAC:           15,
			LoadPvAveragePowerAC:          20,
			LoadBatteryAveragePowerAC:     0,
			BatterySoC:                    80,
			BatteryProducedAveragePowerAC: 20,
			BatteryConsumedAveragePowerAC: 0,
			BatteryChargingFrom:           "Solar",
			BatteryDischargingTo:          "",
			PvAveragePowerAC:              40,
			LoadAveragePowerAC:            30,
			LoadLinks:                     loadLinksValue,
			GridLinks:                     gridLinksValue,
			PvLinks:                       pvLinksValue,
			BatteryLinks:                  batteryLinksValue,
			BatteryPvAveragePowerAC:       20,
			GridPvAveragePowerAC:          0,
			GridProducedAveragePowerAC:    10,
			GridConsumedAveragePowerAC:    0,
		}
		response := app.Response{
			Code: e.Success,
			Msg:  e.GetMsg(e.Success),
			Data: dashboardData,
		}
		err := c.Conn.WriteJSON(response)
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
