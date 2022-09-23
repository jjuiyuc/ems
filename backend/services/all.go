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
	Billing BillingService
}

// NewServices godoc
func NewServices(cfg *viper.Viper, repo *repository.Repository) (services *Services) {
	services = &Services{
		Auth:    NewAuthService(repo),
		Email:   NewEmailService(cfg),
		User:    NewUserService(repo),
		Billing: NewBillingService(repo),
	}
	services.Battery = NewBatteryService(repo, services.Billing)
	services.Devices = NewDevicesService(repo, services.Billing)
	return
}
