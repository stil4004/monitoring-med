package secure

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CalcSignature(secret string, message string) string {
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func CalcInternalId(data ...any) string {
	hasher := sha512.New()
	hasher.Write([]byte(fmt.Sprintf("%v_%s", data, time.Now().Format(time.RFC3339))))
	return hex.EncodeToString(hasher.Sum(nil))
}

func DecodeToken(token string) string {
	claims := jwt.MapClaims{}
	_, _ = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	var nickname string
	nickname = fmt.Sprint(claims["nickname"])
	return nickname
}
