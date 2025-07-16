package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"main/response"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := extractUserIDFromJWT(c, jwtSecret)
		if err != nil {
			response.UnauthorizedError(c, err.Error())
			return
		}

		c.Set("UserId", userID)
		c.Next()
	}
}

func extractUserIDFromJWT(c *gin.Context, jwtSecret string) (uuid.UUID, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return uuid.Nil, errors.New("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return uuid.Nil, errors.New("invalid authorization header format")
	}

	tokenStr := parts[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return uuid.Nil, errors.New("missing sub claim")
	}

	userID, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, errors.New("invalid uuid in token")
	}

	return userID, nil
}

func GetUserIDFromContext(c *gin.Context) (uuid.UUID, error) {
	rawUserID, exists := c.Get("UserId")
	if !exists {
		return uuid.Nil, errors.New("user ID not found in context")
	}

	userID, ok := rawUserID.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("invalid user ID type in context")
	}

	return userID, nil
}

func GenerateJWT(userID uuid.UUID, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
