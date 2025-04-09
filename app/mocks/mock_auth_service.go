package mocks

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// MockJWTService mocks the JWTService interface
type MockJWTService struct {
	mock.Mock
}

// GenerateToken is a mocked method
func (m *MockJWTService) GenerateToken(username string, hospitalID uint) string {
	args := m.Called(username, hospitalID)
	return args.String(0)
}

// ValidateToken is a mocked method
func (m *MockJWTService) ValidateToken(token string) (*jwt.Token, error) {
	args := m.Called(token)
	return args.Get(0).(*jwt.Token), args.Error(1)
}

// GenerateRefreshToken is a mocked method
func (m *MockJWTService) GenerateRefreshToken(username string) string {
	args := m.Called(username)
	return args.String(0)
}

// GetPayloadInToken is a mocked method
func (m *MockJWTService) GetPayloadInToken(c *gin.Context) jwt.MapClaims {
	args := m.Called(c)
	return args.Get(0).(jwt.MapClaims)
}
