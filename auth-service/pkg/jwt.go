package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func SignToken(userId uuid.UUID, username, role string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpiresIn := os.Getenv("JWT_EXPIRES_IN")

	claims := jwt.MapClaims{
		"uid":  userId,
		"user": username,
		"role": role,
	}

	if jwtExpiresIn != "" {
		duration, err := time.ParseDuration(jwtExpiresIn)
		if err != nil {
			return "", ErrorHandler(err, "internal error")
		}
		claims["exp"] = jwt.NewNumericDate(time.Now().Add(duration))

	} else {
		claims["exp"] = jwt.NewNumericDate(time.Now().Add(15 * time.Minute))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", ErrorHandler(err, "Internal error")
	}
	return signedToken, nil
}

func VerifyToken(tokenString string) (uuid.UUID, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid token claims")
	}

	uidStr, ok := claims["uid"].(string)
	if !ok {
		return uuid.Nil, errors.New("invalid uid in token")
	}

	userID, err := uuid.Parse(uidStr)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

