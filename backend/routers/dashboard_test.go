package routers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"der-ems/config"
	"der-ems/internal/e"
	"der-ems/internal/utils"
	"der-ems/models"
	"der-ems/repository"
	"der-ems/services"
	"der-ems/testutils"
	"der-ems/testutils/fixtures"
)

type DashboardSuite struct {
	suite.Suite
	router *gin.Engine
	repo   *repository.Repository
	worker *APIWorker
	token  string
}

func Test_Dashboard(t *testing.T) {
	suite.Run(t, &DashboardSuite{})
}

func (s *DashboardSuite) SetupSuite() {
	config.Init(testutils.GetConfigDir(), "ut.yaml")
	cfg := config.GetConfig()
	models.Init(cfg)
	db := models.GetDB()

	repo := repository.NewRepository(db)
	w := &APIWorker{
		Services: services.NewServices(cfg, repo),
	}
	s.worker = w

	// Truncate & seed data
	err := testutils.SeedUtUser(db)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	err = testutils.SeedUtLocationAndGateway(db)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	token, err := utils.GenerateToken(fixtures.UtUser.ID)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.token = token
	// Mock user_gateway_right table
	_, err = db.Exec("TRUNCATE TABLE user_gateway_right")
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	_, err = db.Exec(`
			INSERT INTO user_gateway_right (id,user_id,gw_id) VALUES
			(1,1,1);
		`)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())

	s.repo = repo
	s.router = InitRouter(cfg.GetBool("server.cors"), cfg.GetString("server.ginMode"), w)
}

func (s *DashboardSuite) TearDownSuite() {
	models.Close()
}

func (s *DashboardSuite) Test_dashboardHandler() {
	type response struct {
		Code int                                 `json:"code"`
		Msg  string                              `json:"msg"`
		Data *services.DevicesEnergyInfoResponse `json:"data"`
	}

	testLoadLinks := map[string]interface{}{
		"grid":    0.0,
		"battery": 0.0,
		"pv":      0.0,
	}
	testGridLinks := map[string]interface{}{
		"load":    1.0,
		"battery": 0.0,
		"pv":      0.0,
	}
	testPvLinks := map[string]interface{}{
		"load":    1.0,
		"battery": 1.0,
		"grid":    0.0,
	}
	testBatteryLinks := map[string]interface{}{
		"load": 0.0,
		"pv":   0.0,
		"grid": 0.0,
	}
	testResponseData := &services.DevicesEnergyInfoResponse{
		GridIsPeakShaving:             0,
		LoadGridAveragePowerAC:        10,
		BatteryGridAveragePowerAC:     0,
		GridContractPowerAC:           15,
		LoadPvAveragePowerAC:          20,
		LoadBatteryAveragePowerAC:     0,
		BatterySoC:                    160,
		BatteryProducedAveragePowerAC: 20,
		BatteryConsumedAveragePowerAC: 0,
		BatteryChargingFrom:           "Solar",
		BatteryDischargingTo:          "",
		PvAveragePowerAC:              40,
		LoadAveragePowerAC:            30,
		LoadLinks:                     testLoadLinks,
		GridLinks:                     testGridLinks,
		PVLinks:                       testPvLinks,
		BatteryLinks:                  testBatteryLinks,
		BatteryPvAveragePowerAC:       20,
		GridPvAveragePowerAC:          0,
		GridProducedAveragePowerAC:    10,
		GridConsumedAveragePowerAC:    0,
	}

	server := httptest.NewServer(http.HandlerFunc(s.dashboardHandler))
	defer server.Close()
	seedUtURLStr := "ws" + strings.TrimPrefix(server.URL, "http") + "/api/" + fixtures.UtGateway.UUID + "/devices/energy-info"
	tt := struct {
		name   string
		token  string
		urlStr string
		wantRv response
	}{
		name:   "devicesEnergyInfo",
		token:  s.token,
		urlStr: seedUtURLStr,
		wantRv: response{
			Code: e.Success,
			Msg:  "ok",
			Data: testResponseData,
		},
	}

	log.Info("test name: ", tt.name)
	ws, _, err := websocket.DefaultDialer.Dial(tt.urlStr, http.Header{"Sec-WebSocket-Protocol": {tt.token}})
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	defer ws.Close()

	_, p, err := ws.ReadMessage()
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())

	var res response
	err = json.Unmarshal([]byte(p), &res)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	s.Equalf(tt.wantRv.Code, res.Code, e.ErrNewMessageNotEqual.Error())
	s.Equalf(tt.wantRv.Msg, res.Msg, e.ErrNewMessageNotEqual.Error())
	s.Equalf(tt.wantRv.Data, res.Data, e.ErrNewMessageNotEqual.Error())
}

func (s *DashboardSuite) dashboardHandler(writer http.ResponseWriter, request *http.Request) {
	s.Equalf(s.token, request.Header["Sec-Websocket-Protocol"][0], e.ErrNewMessageNotEqual.Error())

	pool := newPool()
	go pool.start()
	conn, err := s.worker.upgrade(writer, request)
	s.Require().NoErrorf(err, e.ErrNewMessageReceivedUnexpectedErr.Error())
	client := &Client{
		ID:          fixtures.UtUser.ID,
		Token:       s.token,
		GatewayUUID: fixtures.UtGateway.UUID,
		Conn:        conn,
		Pool:        pool,
	}

	pool.Register <- client
	client.run(s.worker)
}
