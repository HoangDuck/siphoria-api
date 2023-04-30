package security

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"hotel-booking-api/model"
	"os"
	"strconv"
	"time"
)

func GenTokenResetPassword(email string) (string, error) {
	jwtExpConfig := os.Getenv("JWT_EXPIRED")
	jwtExpValue, _ := strconv.Atoi(jwtExpConfig)

	jwtExpDuration :=
		time.Hour * time.Duration(jwtExpValue)
	claims := &model.JwtCustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenToken(user *model.User) (string, error) {
	jwtExpConfig := os.Getenv("JWT_EXPIRED")
	jwtExpValue, _ := strconv.Atoi(jwtExpConfig)

	jwtExpDuration := time.Hour * time.Duration(jwtExpValue)

	claims := &model.JwtCustomClaims{
		UserId: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	user.Token = &model.Token{}
	user.Token.AccessToken = tokenString
	return tokenString, nil
}
func GenRefToken(user *model.User) (string, time.Duration, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	var key = os.Getenv("SECRET_REFRESH_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	timeDura := time.Duration(time.Now().Add(time.Hour * 48).Unix())
	if err != nil {
		return "", timeDura, err
	}
	user.Token.RefreshToken = tokenString
	user.Token.ExpiredTime = timeDura
	return tokenString, timeDura, nil
}

func GetClaimsJWT(c echo.Context) *model.JwtCustomClaims {
	token := c.Get("user").(*jwt.Token)
	return token.Claims.(*model.JwtCustomClaims)
}

func CheckRole(claim *model.JwtCustomClaims, RoleCheck model.Role) bool {
	return claim.Role == RoleCheck.String()
}
