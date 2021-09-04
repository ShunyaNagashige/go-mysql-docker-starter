package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type userInfo struct {
	userId   uint
	userName string
	email    string
}

type customClaims struct {
	*jwt.StandardClaims
	userInfo
}

func CreateToken(userId uint, userName, email string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))

	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", xerrors.Errorf("Failed to get a random UUID. : %w", err)
	}

	t.Claims = &customClaims{
		&jwt.StandardClaims{
			// JWTの対象となる受信者
			// Audience:  "",

			// 有効期限
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),

			// JWTのユニーク性を担保するID
			// 同じJWTを使い回すことを抑制する
			Id: uuidObj.String(),

			// JWTが発行された時刻
			IssuedAt: time.Now().Unix(),

			// JWTの発行者
			// Issuer:    "",

			// JWTが有効になる時刻
			NotBefore: time.Now().Add(time.Second * -5).Unix(),

			// JWTの用途
			Subject: "AccessToken",
		},
		userInfo{
			userId:   userId,
			userName: userName,
			email:    email,
		},
	}

	token, err := t.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		return "", xerrors.Errorf("Failed to get the signed token. : %w", err)
	}

	return token, nil
}
