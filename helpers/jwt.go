package helpers

import (
	"candra/backend-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Nilai secret diambil dari environment variable JWT_SECRET
// var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

// GenerateToken membuat token JWT untuk pengguna terautentikasi
func GenerateToken(username string) (string, error) {
	jwtKey := []byte(config.GetEnv("JWT_SECRET", "secret_key"))
	// fmt.Println("JWT_SECRET:", config.GetEnv("JWT_SECRET", "secret_key"))
	// Membuat klaim (claims) JWT
	// Subject berisi username, dan ExpiresAt menentukan waktu expired token
	claims := &jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	// Membuat token baru dengan klaim yang telah dibuat
	// Menggunakan algoritma HS256 untuk menandatangani token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
