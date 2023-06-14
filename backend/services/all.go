package services

import (
	"github.com/spf13/viper"

	"der-ems/repository"
)

// Services godoc
type Services struct {
	Auth              AuthService
	Email             EmailService
	User              UserService
	Devices           DevicesService
	Billing           BillingService
	AccountManagement AccountManagementService
	FieldManagement   FieldManagementService
}

// NewServices godoc
func NewServices(cfg *viper.Viper, repo *repository.Repository) (services *Services) {
	services = &Services{
		Auth:              NewAuthService(repo),
		Email:             NewEmailService(cfg),
		User:              NewUserService(repo),
		Billing:           NewBillingService(repo),
		AccountManagement: NewAccountManagementService(repo),
	}
	services.Devices = NewDevicesService(repo, services.Billing)
	services.FieldManagement = NewFieldManagementService(repo, services.AccountManagement)
	return
}
