package testutils

import (
	"github.com/gin-gonic/gin"
)

// MockEmailService ...
type MockEmailService struct{}

// SendResetEmail ...
func (m *MockEmailService) SendResetEmail(c *gin.Context, name, address, token string) error {
	return nil
}
