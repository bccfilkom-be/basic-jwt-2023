package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bccfilkom-be/basic-jwt-2023/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Fungsi untuk membuat token
// Menambahkan ID dalam tokennya dan Waktu expired token
func GenerateToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  int(id),
		"exp": time.Now().Add(20 * time.Minute).Unix(),
	})
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		//Handle Error
		return "", err
	}
	return signedToken, nil
}

// Fungsi untuk mengvalidasi token
// Middleware
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			c.JSON(http.StatusBadRequest, utils.ResponseWhenFail("token not found"))
			c.Abort()
			return
		}
		// Authorization : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzgwMDQyNjMsImlkIjoxMH0.73CGtJVGusHlchhaEFavFOh7z37JnZDW-bv4Gh6z1-w
		bearerToken = strings.ReplaceAll(bearerToken, "Bearer ", "")
		tokenAfterEkstract, err := jwt.Parse(bearerToken, ekstractToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.ResponseWhenFail(err.Error()))
			c.Abort()
			return
		}
		if claims, ok := tokenAfterEkstract.Claims.(jwt.MapClaims); ok && tokenAfterEkstract.Valid {
			userId := uint(claims["id"].(float64))
			c.Set("id", userId)
			c.Next()
			return
		}
		c.JSON(http.StatusForbidden, utils.ResponseWhenFail("invalid token"))
		c.Abort()
	}
}

// Validasi token
// Gagal => JSON dan keterangan error
// Berhasil => [POST] todo => Json TODO

func ekstractToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}
	return []byte(os.Getenv("SECRET_KEY")), nil
}
