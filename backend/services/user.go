package services

import (
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"

	"der-ems/internal/e"
	"der-ems/internal/utils"
	deremsmodels "der-ems/models/der-ems"
	"der-ems/repository"
)

// UserService godoc
type UserService interface {
	CreatePasswordToken(username string) (name, token string, err error)
	PasswordResetByPasswordToken(token, newPassword string) (err error)
	GetProfile(userID int64) (profile *ProfileResponse, err error)
	UpdateName(userID int64, name string) (err error)
	UpdatePassword(userID int64, currentPassword, newPassword string) (errCode int, err error)
}

type defaultUserService struct {
	repo *repository.Repository
}

// ProfileResponse godoc
type ProfileResponse struct {
	*deremsmodels.User
	Group GroupInfo `json:"group"`
}

// GroupInfo godoc
type GroupInfo struct {
	*deremsmodels.Group
	Gateways []GatewayInfo `json:"gateways"`
	Webpages []WebpageInfo `json:"webpages"`
}

// GatewayInfo godoc
type GatewayInfo struct {
	GatewayID   string                  `json:"gatewayID"`
	Permissions []GatewayPermissionInfo `json:"permissions"`
}

// WebpageInfo godoc
type WebpageInfo struct {
	ID          int64                  `json:"id"`
	Name        string                 `json:"name"`
	Permissions WebpagePermissionsInfo `json:"permissions"`
}

// GatewayPermissionInfo godoc
type GatewayPermissionInfo struct {
	EnabledAt  time.Time    `json:"enabledAt"`
	EnabledBy  null.Int64   `json:"enabledBy"`
	DisabledAt null.Time    `json:"disabledAt"`
	DisabledBy null.Int64   `json:"disabledBy"`
	Location   LocationInfo `json:"location"`
}

// LocationInfo godoc
type LocationInfo struct {
	Name    string      `json:"name"`
	Address null.String `json:"address"`
}

// WebpagePermissionsInfo godoc
type WebpagePermissionsInfo struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

// NewUserService godoc
func NewUserService(repo *repository.Repository) UserService {
	return &defaultUserService{repo}
}

// CreatePasswordToken godoc
func (s defaultUserService) CreatePasswordToken(username string) (name, token string, err error) {
	user, err := s.repo.User.GetUserByUsername(username)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetUserByUsername",
			"err":       err,
		}).Error()
		return
	}

	token = uuid.New().String()
	user.ResetPWDToken = null.StringFrom(token)
	user.PWDTokenExpiry = null.TimeFrom(time.Now().UTC().Add(1 * time.Hour))
	err = s.repo.User.UpdateUser(nil, user)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.UpdateUser",
			"err":       err,
		}).Error()
		return
	}

	name = user.Name.String
	return
}

// PasswordResetByPasswordToken godoc
func (s defaultUserService) PasswordResetByPasswordToken(token, newPassword string) (err error) {
	user, err := s.repo.User.GetUserByPasswordToken(token)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetUserByPasswordToken",
			"err":       err,
		}).Error()
		return
	}

	hashedPassword, err := utils.CreateHashedPassword(newPassword)
	if err != nil {
		return
	}
	user.Password = hashedPassword
	user.PasswordLastChanged = null.TimeFrom(time.Now().UTC())
	user.ResetPWDToken = null.StringFrom("")
	err = s.repo.User.UpdateUser(nil, user)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.UpdateUser",
			"err":       err,
		}).Error()
	}
	return
}

// GetProfile godoc
func (s defaultUserService) GetProfile(userID int64) (profile *ProfileResponse, err error) {
	// Get user information
	user, err := s.repo.User.GetUserByUserID(nil, userID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}

	// Get group information
	group, err := s.repo.User.GetGroupByGroupID(nil, user.GroupID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetGroupByGroupID",
			"err":       err,
		}).Error()
		return
	}
	groupInfo := GroupInfo{
		Group:    group,
		Gateways: s.getGroupGatewayInfo(group.ID),
	}
	groupInfo.Webpages, err = s.getGroupWebpageInfo(group.TypeID)
	if err != nil {
		return
	}

	profile = &ProfileResponse{
		User:  user,
		Group: groupInfo,
	}
	return
}

func (s defaultUserService) getGroupGatewayInfo(groupID int64) (gatewayInfos []GatewayInfo) {
	gatewaysPermission, getErr := s.repo.User.GetGatewaysPermissionByGroupID(nil, groupID, true)
	if getErr != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetGatewaysPermissionByGroupID",
			"err":       getErr,
		}).Warn()
	}
	for _, gatewayPermission := range gatewaysPermission {
		var (
			gatewayInfo GatewayInfo
			permissions []GatewayPermissionInfo
		)
		// Ignore gateway without location
		if gatewayPermission.LocationID.IsZero() {
			continue
		}
		gateway, getErr := s.repo.Gateway.GetGatewayByGatewayID(gatewayPermission.GWID)
		if getErr != nil {
			log.WithFields(log.Fields{
				"caused-by": "s.repo.Gateway.GetGatewayByGatewayID",
				"err":       getErr,
			}).Warn()
			continue
		}
		// Check if gateway UUID exists
		existedGatewayIndex := -1
		for i, existedGatewayInfo := range gatewayInfos {
			if existedGatewayInfo.GatewayID == gateway.UUID {
				existedGatewayIndex = i
				gatewayInfo = existedGatewayInfo
				permissions = gatewayInfo.Permissions
				break
			}
		}
		if existedGatewayIndex >= 0 {
			gatewayInfos = s.removeGatewayInfo(gatewayInfos, existedGatewayIndex)
		}
		gatewayInfo.GatewayID = gateway.UUID
		permissionInfo := GatewayPermissionInfo{
			EnabledAt:  gatewayPermission.EnabledAt,
			EnabledBy:  gatewayPermission.EnabledBy,
			DisabledAt: gatewayPermission.DisabledAt,
			DisabledBy: gatewayPermission.DisabledBy,
		}

		location, getErr := s.repo.Location.GetLocationByLocationID(gatewayPermission.LocationID.Int64)
		if getErr == nil {
			permissionInfo.Location = LocationInfo{
				location.Name,
				location.Address,
			}
		} else {
			log.WithFields(log.Fields{
				"caused-by": "s.repo.Location.GetLocationByLocationID",
				"err":       getErr,
			}).Warn()
		}

		permissions = append(permissions, permissionInfo)
		gatewayInfo.Permissions = permissions
		gatewayInfos = append(gatewayInfos, gatewayInfo)
	}
	return
}

func (s defaultUserService) getGroupWebpageInfo(groupTypeID int64) (webpageInfos []WebpageInfo, err error) {
	webpagesPermission, err := s.repo.User.GetWebpagesPermissionByGroupTypeID(groupTypeID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetWebpagesPermissionByGroupTypeID",
			"err":       err,
		}).Error()
		return
	}
	for _, webpagePermission := range webpagesPermission {
		webpage, err := s.repo.User.GetWebpageByWebpageID(webpagePermission.WebpageID)
		if err != nil {
			log.WithFields(log.Fields{
				"caused-by": "s.repo.User.GetWebpageByWebpageID",
				"err":       err,
			}).Error()
			break
		}

		webpageInfo := WebpageInfo{
			webpage.ID,
			webpage.Name,
			WebpagePermissionsInfo{
				webpagePermission.CreateData.Bool,
				webpagePermission.ReadData.Bool,
				webpagePermission.UpdateData.Bool,
				webpagePermission.DeleteData.Bool,
			},
		}
		webpageInfos = append(webpageInfos, webpageInfo)
	}
	return
}

func (s defaultUserService) UpdateName(userID int64, name string) (err error) {
	user, err := s.repo.User.GetUserByUserID(nil, userID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}

	user.Name = null.StringFrom(name)
	if err = s.repo.User.UpdateUser(nil, user); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.UpdateUser",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultUserService) UpdatePassword(userID int64, currentPassword, newPassword string) (errCode int, err error) {
	user, err := s.repo.User.GetUserByUserID(nil, userID)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.GetUserByUserID",
			"err":       err,
		}).Error()
		return
	}

	if err = utils.ComparePassword(currentPassword, user.Password); err != nil {
		errCode = e.ErrAuthPasswordNotMatch
		return
	}

	hashPassword, err := utils.CreateHashedPassword(newPassword)
	if err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.createHashPassword",
			"err":       err,
		}).Error()
		return
	}
	user.Password = hashPassword
	user.PasswordLastChanged = null.TimeFrom(time.Now().UTC())
	if err = s.repo.User.UpdateUser(nil, user); err != nil {
		log.WithFields(log.Fields{
			"caused-by": "s.repo.User.UpdateUser",
			"err":       err,
		}).Error()
	}
	return
}

func (s defaultUserService) removeGatewayInfo(slice []GatewayInfo, i int) []GatewayInfo {
	return append(slice[:i], slice[i+1:]...)
}
