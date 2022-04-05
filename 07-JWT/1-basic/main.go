package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	utils "github.com/golinuxcloudnative/go-web-authentication/99-utils"
)

type UserClaims struct {
	// composition.  embedded jwt.StandardClaims struct and methods
	jwt.StandardClaims
	SessionID int64  `json:"session-id"`
	Username  string `json:"username"`
}

// function to valid the token. You custom it as you want.
func (uc *UserClaims) Valid() error {
	if !uc.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}

	if uc.SessionID == 0 {
		return fmt.Errorf("Invalid session id")
	}
	return nil
}

func main() {
	// expired token
	claim := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(-5 * time.Minute).Unix()),
		},
		SessionID: 0,
		Username:  "turnes",
	}
	utils.PrettyPrint(claim)
	fmt.Printf("%v - %v\n", claim.Valid(), time.Now().Unix())

	//  invalid session
	claim = UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		SessionID: 0,
		Username:  "turnes",
	}
	utils.PrettyPrint(claim)
	fmt.Printf("%v - session cannot be 0\n", claim.Valid())

	// valid token
	claim = UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		SessionID: 814804382098,
		Username:  "turnes",
	}
	utils.PrettyPrint(claim)
	fmt.Printf("Session valid - %v\n", time.Now().Unix())

}
