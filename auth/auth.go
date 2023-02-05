package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v4"
)

// Auth ...
type Auth struct {
	jwk               *JWK
	jwkURL            string
	cognitoRegion     string
	cognitoUserPoolID string
}

// Config ...
type Config struct {
	CognitoRegion     string
	CognitoUserPoolID string
}

// JWK ...
type JWK struct {
	Keys []struct {
		Alg string `json:"alg"`
		E   string `json:"e"`
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		N   string `json:"n"`
	} `json:"keys"`
}

// NewAuth ...
func NewAuth(config *Config) *Auth {
	a := &Auth{
		cognitoRegion:     config.CognitoRegion,
		cognitoUserPoolID: config.CognitoUserPoolID,
	}

	a.jwkURL = fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", a.cognitoRegion, a.cognitoUserPoolID)
	err := a.CacheJWK()
	if err != nil {
		log.Fatal(err)
	}

	return a
}

// CacheJWK ...
func (a *Auth) CacheJWK() error {
	req, err := http.NewRequest("GET", a.jwkURL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	jwk := new(JWK)
	err = json.Unmarshal(body, jwk)
	if err != nil {
		return err
	}

	a.jwk = jwk
	return nil
}

// ParseJWT ...
func (a *Auth) ParseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		key := convertKey(a.jwk.Keys[1].E, a.jwk.Keys[1].N)
		return key, nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}

// JWK ...
func (a *Auth) JWK() *JWK {
	return a.jwk
}

// JWKURL ...
func (a *Auth) JWKURL() string {
	return a.jwkURL
}

// https://gist.github.com/MathieuMailhos/361f24316d2de29e8d41e808e0071b13
func convertKey(rawE, rawN string) *rsa.PublicKey {
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		panic(err)
	}
	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(rawN)
	if err != nil {
		panic(err)
	}
	pubKey.N.SetBytes(decodedN)
	return pubKey
}

func CheckJWT(CognitoUserPool string, CognitoRegion string, tokenStr string) (bool, error, string) {

	auth := NewAuth(&Config{
		CognitoRegion:     CognitoRegion,
		CognitoUserPoolID: CognitoUserPool,
	})

	err := auth.CacheJWK()
	if err != nil {
		return false, err, string("")
	}

	token, err := auth.ParseJWT(tokenStr)
	if err != nil {
		return false, err, string("")
	}

	if !token.Valid {
		return false, nil, token.Claims.Valid().Error()
	}

	return true, nil, string("")
}
