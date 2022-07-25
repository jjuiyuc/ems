package services

import (
	"github.com/spf13/viper"

	"der-ems/repository"
)

// Services godoc
type Services struct {
	Auth    AuthService
	Email   EmailService
	User    UserService
	Devices DevicesService
	Battery BatteryService
}

// NewServices godoc
func NewServices(cfg *viper.Viper, repo *repository.Repository) *Services {
	return &Services{
		Auth:    NewAuthService(repo),
		Email:   NewEmailService(cfg),
		User:    NewUserService(repo),
		Devices: NewDevicesService(repo),
		Battery: NewBatteryService(repo),
	}
}
