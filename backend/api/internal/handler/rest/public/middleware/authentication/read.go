package authentication

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func AuthRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, err := AUTH_CONFIG.GetTokenFromHeaderAndVerify(w, r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GetUserIdFromToken(r *http.Request) int {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	tokenString := splitToken[1]

	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte{}, nil
	})

	id, err := strconv.Atoi(token.Claims.(jwt.MapClaims)["sub"].(string))

	if err != nil {
		return 0
	}

	return id
}
