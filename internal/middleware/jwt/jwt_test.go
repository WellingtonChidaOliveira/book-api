package jwt

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	secret := "mysecret"
	jwtMiddleware := NewJWT(secret)
	userID := "12345"
	role := "admin"

	token, err := jwtMiddleware.GenerateToken(userID, role)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwtMiddleware.ValidateToken(token)
	assert.NoError(t, err)
	assert.NotNil(t, parsedToken)
	assert.Equal(t, jwt.SigningMethodHS256.Alg(), parsedToken.Method.Alg())
	assert.Equal(t, token, parsedToken.Raw)
	claims, err := jwtMiddleware.ExtractClaims(parsedToken)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims["user_id"])
	assert.Equal(t, role, claims["role"])
	assert.NotEmpty(t, claims["exp"])
}
