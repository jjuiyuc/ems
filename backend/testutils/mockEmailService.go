package testutils

import (
	"github.com/gin-gonic/gin"
)

// MockEmailService godoc
type MockEmailService struct{}

// SendResetEmail godoc
func (m *MockEmailService) SendResetEmail(c *gin.Context, name, address, token string) error {
	return nil
}
