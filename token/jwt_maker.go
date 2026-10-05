package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//we will use this to compare it with the secret key in the app.env

const minSecretKeySize=32


type JWTMaker struct{
	secretKey string
}

//constructor to be included in main file and pass to it the secret key also you will use viper to read env file you can see my viper repo if u need
func NewJwtMaker(secertKey string)(Maker,error){
	if len(secertKey)<minSecretKeySize{
		return nil,fmt.Errorf("invalid secretKeySize must be at least:%d",minSecretKeySize)
	}
	return &JWTMaker{secretKey: secertKey},nil
}


func(maker *JWTMaker)CreateToken(email string,role string,duration time.Duration)(string,error){
	Paylod,err:=NewPayload(email,role,duration)
	if err!=nil{
		return "",err
	}
	jwtToken:=jwt.NewWithClaims(jwt.SigningMethodES256,jwt.MapClaims{
		"email":Paylod.Email,
		"role":Paylod.Role,
		"iat":Paylod.IssuedAt.Unix(),
		"exp":Paylod.ExpiredAt.Unix(),
	})
	return jwtToken.SignedString([]byte(maker.secretKey))
}








