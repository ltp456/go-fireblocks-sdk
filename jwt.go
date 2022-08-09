package go_fireblocks

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type TokenProvider struct {
	apiKey    string
	secretKey []byte // todo memory safe
}

func NewTokenProvider(apiKey string, secretKey []byte) *TokenProvider {
	return &TokenProvider{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func (tp *TokenProvider) GetApiKey() string {
	return tp.apiKey
}

func (tp *TokenProvider) SignJwt(path string, body []byte) (string, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(tp.secretKey)
	if err != nil {
		return "", err
	}
	hash, err := Sha256(body)
	if err != nil {
		return "", err
	}
	iat := time.Now().Unix()
	mapClaims := jwt.MapClaims{
		"uri":      path,
		"nonce":    uuid.New().ID(),
		"iat":      iat,
		"exp":      iat + 50,
		"sub":      tp.apiKey,
		"bodyHash": hash,
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(jwt.SigningMethodRS256.Alg()), mapClaims)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", nil
	}
	return tokenStr, nil
}
