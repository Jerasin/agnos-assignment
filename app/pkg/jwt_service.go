package pkg

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"agnos-assignment/app/config"
	"agnos-assignment/app/constant"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserContext struct {
	Username   string
	HospitalID uint
}

type JWTServiceInterface interface {
	GenerateToken(username string, hospitalID uint) string
	ValidateToken(token string) (*jwt.Token, error)
	GenerateRefreshToken(username string) string
	GetPayloadInToken(c *gin.Context) *UserContext
}

type JWTService struct {
	secretKey string
	issure    string
}

type authCustomClaims struct {
	Username   string `json:"username"`
	HospitalID uint   `json:"hospital_id"`
	jwt.StandardClaims
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func JWTServiceInit() JWTServiceInterface {
	return &JWTService{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func (service *JWTService) GenerateRefreshToken(username string) string {
	config.EnvConfig()
	JWT_EXPIRE_MINUTE := config.GetEnv("JWT_EXPIRE_MINUTE", "15")

	expire_time, err := strconv.Atoi(JWT_EXPIRE_MINUTE)
	if err != nil {
		// ... handle error
		panic(err)
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["username"] = username
	rtClaims["exp"] = time.Now().Add(time.Minute * time.Duration(expire_time)).Unix()
	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		panic(err)
	}

	return rt
}

func (service *JWTService) GenerateToken(username string, hospitalID uint) string {
	config.EnvConfig()
	JWT_EXPIRE_MINUTE := config.GetEnv("JWT_EXPIRE_MINUTE", "15")

	expire_time, err := strconv.Atoi(JWT_EXPIRE_MINUTE)
	if err != nil {
		// ... handle error
		panic(err)
	}

	claims := &authCustomClaims{
		username,
		hospitalID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expire_time)).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token = %s", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}

func (service *JWTService) GetPayloadInToken(c *gin.Context) *UserContext {
	defer PanicHandler(c)

	var claims jwt.MapClaims
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		PanicException(constant.Unauthorized)
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := service.ValidateToken(tokenString)

	if token.Valid {
		claims = token.Claims.(jwt.MapClaims)
		fmt.Println("claims", claims)

	} else {
		fmt.Println("testing")
		fmt.Println(err)
		PanicException(constant.Unauthorized)
	}

	return mapUserContext(claims)
}

func mapUserContext(payload jwt.MapClaims) *UserContext {
	username, ok := payload["username"].(string)

	if !ok {
		PanicException(constant.BadRequest)
	}

	hospitalID, ok := payload["hospital_id"].(float64)
	if !ok {
		PanicException(constant.BadRequest)
	}

	return &UserContext{
		Username:   username,
		HospitalID: uint(hospitalID),
	}
}
