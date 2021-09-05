package auth

import (
	"os"
	"time"

	"github.com/ShunyaNagashige/golang-jwt-sample/domain/model"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type UserInfo struct {
	UserId   uint
	UserName string
	Email    string
}

type CustomClaims struct {
	UserInfo
	jwt.StandardClaims
}

func CreateToken(userId uint, userName, email string) (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", xerrors.Errorf("Failed to get a random UUID. : %w", err)
	}

	claims := CustomClaims{
		UserInfo{
			UserId:   userId,
			UserName: userName,
			Email:    email,
		},
		jwt.StandardClaims{
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
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		return "", xerrors.Errorf("Failed to get the signed token. : %w", err)
	}

	return ss, nil
}

func ValidateToken(tokenStr string) (*model.User, error) {
	cc := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, cc, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})

	// token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	// 	// validate the alg is what you expect
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, xerrors.Errorf("The singing method of jwt must be HMAC")
	// 	}

	// 	return []byte(os.Getenv("SIGNINGKEY")), nil
	// })
	if err != nil {
		return nil, xerrors.Errorf("Failed to parse and validate a token : %w", err)
	}

	if !token.Valid {
		return nil, xerrors.Errorf("JWT is invalid")
	}

	if claims, ok := token.Claims.(*CustomClaims); !ok {
		return nil, xerrors.Errorf("Can't cast variant value %#v to *customClaims", token.Claims)
	} else {
		return &model.User{
			UserId:   claims.UserInfo.UserId,
			UserName: claims.UserInfo.UserName,
			Email:    claims.UserInfo.Email,
		}, nil
	}
}
