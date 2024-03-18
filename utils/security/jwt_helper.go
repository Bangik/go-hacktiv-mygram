package security

import (
	"fmt"
	"hacktiv-assignment-final/config"
	"hacktiv-assignment-final/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user model.User) (string, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return "", fmt.Errorf("failed to create access token: %s", err.Error())
	}

	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenLifeTime)

	claims := &TokenMyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Id: user.ID,
	}

	token := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(cfg.JwtSignatureKey)
	if err != nil {
		return "", fmt.Errorf("failed to create access token : %s", err.Error())
	}
	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to verify access token : %s", err.Error())
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != cfg.JwtSigningMethod {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid parse token sdf : %s", err.Error())
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token MapClaims")
	}
	return claims, nil
}

func GetIdFromToken(c *gin.Context) (string, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return "", fmt.Errorf("failed to get id from token")
	}

	claimsMap := claims.(jwt.MapClaims)
	id := claimsMap["id"].(string)

	return id, nil
}
