package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	timeOutInSeconds = 10000000
	timeOutInMinues  = 60
	sessionSecret    = "yourSe$$ion$ecret"
)

// //CreateToken : ""
// func CreateToken(authentication *models.Authentication) (string, error) {
// 	var claims = jws.Claims{
// 		"UserID":   authentication.UserID,
// 		"UserName": authentication.UserName,
// 		"Status":   authentication.Status,
// 		"Role":     authentication.Role,
// 	}
// 	// claims.SetIssuedAt(time.Now())
// 	// claims.SetExpiration(time.Now().Add(time.Duration(timeOutInMinues) * time.Minute))

// 	jwt := jws.NewJWT(claims, crypto.SigningMethodHS256)
// 	jwtToken, err := jwt.Serialize([]byte(sessionSecret))
// 	log.Println("TOKKKKKKKEN ", string(jwtToken))
// 	return string(jwtToken), err
// }

//CreateTokenV2
func CreateTokenV2(authentication *models.Authentication) (string, error) {
	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt.MapClaims{
		"data": authentication,
		"exp":  time.Now().Add(time.Second * 50).Unix(),
	}
	// Sign and get the complete encoded token as a string
	return token.SignedString([]byte(sessionSecret))
}

// //ValidateToken : ""
// func ValidateToken(token string) (*models.Authentication, error) {
// 	// TODO: Check if the access token is available on redis store
// 	// if not? then simply return unauthorized

// 	parsedToken, err := jws.ParseJWT([]byte(string(token)))
// 	if err != nil {
// 		return nil, err
// 	}
// 	// err = (parsedToken.Validate([]byte(sessionSecret), crypto.SigningMethodHS256))
// 	// if err != nil {
// 	// 	if err.Error() == "token is expired" {
// 	// 		log.Println("token is expired")
// 	// 		return nil, err
// 	// 	}
// 	// }
// 	cbytes, err1 := json.Marshal(parsedToken.Claims())
// 	if err1 != nil {
// 		return nil, err1
// 	}
// 	authenticationData := new(models.Authentication)
// 	json.Unmarshal(cbytes, &authenticationData)
// 	return authenticationData, nil
// }

//ValidateTokenV2
func ValidateTokenV2(token string) (*models.Authentication, bool, error) {
	var claims jwt.MapClaims
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(sessionSecret), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("Err 1")
			return nil, false, nil
		}
		fmt.Println("Err 2")

		return nil, false, err
	}
	if !tkn.Valid {
		fmt.Println("Err 4")

		return nil, false, nil
	}
	fmt.Println("Err 5")

	cbytes, err1 := json.Marshal(claims["data"])
	if err1 != nil {
		return nil, false, err1
	}
	authenticationData := new(models.Authentication)
	json.Unmarshal(cbytes, &authenticationData)

	return authenticationData, true, nil
}

//Login :
func (s *Service) Login(ctx *models.Context, login *models.Login) (string, bool, error) {
	data, err := s.Daos.GetSingleUser(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return "dal err", false, err
	}
	if ok := data.Password == login.PassWord; !ok {
		log.Println("Data password ==>", data.Password)
		log.Println("login password ==>", login.PassWord)
		return "Passs false", false, nil
	}
	if data.Status == constants.USERSTATUSINIT {
		return "", false, errors.New("Awaiting Activation")
	}
	if data.Status != constants.USERSTATUSACTIVE {
		return "", false, errors.New("Please Contact Administrator")
	}
	// var auth models.Authentication
	// auth.UserID = data.ID
	// auth.UserName = data.UserName

	// auth.Status = data.Status
	// auth.Role = data.Role
	// fmt.Println("auth user ==>", auth, data)
	// token, err := CreateToken(&auth)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return "", false, errors.New("Error in Generating Token - " + err.Error())
	// }
	// data.Token = token
	// data.CurrentLocation = login.Location
	// err = s.Daos.UpdateUserWithUniqueID(data.UserName, data)
	// if err != nil {
	// 	log.Println("Error in saving token - " + err.Error())
	// 	return "", false, errors.New(constants.INTERNALSERVERERROR)
	// }
	return "", true, nil
}

//OTPLoginGenerateOTP :
func (s *Service) OTPLoginGenerateOTP(ctx *models.Context, login *models.Login) error {
	data, err := s.Daos.GetSingleUserWithMobileNo(ctx, login.UserName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if data == nil {
		return errors.New("User Not Available")
	}
	if data.Status == constants.USERSTATUSINIT {
		return errors.New("Awaiting Activation")
	}
	if data.Status != constants.USERSTATUSACTIVE {
		return errors.New("Please Contact Administrator")
	}

	otp, err := s.GenerateOTP(constants.USERLOGIN, login.UserName, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return errors.New("Otp Generate Error - " + err.Error())
	}
	text := fmt.Sprintf("Hi %v, /n Otp For Logikoof Reporting App Login is %v .", data.Name, otp)
	err = s.SendSMS(login.UserName, text)
	if err != nil {
		return errors.New("Sms Sending Error - " + err.Error())
	}

	return nil
}

//OTPLoginValidateOTP :
func (s *Service) OTPLoginValidateOTP(ctx *models.Context, login *models.OTPLogin) (*models.RefUser, bool, error) {

	data, err := s.Daos.GetSingleUserWithMobileNo(ctx, login.Mobile)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}
	if data == nil {
		return nil, false, errors.New("User Not Available")
	}
	if data.Status == constants.USERSTATUSINIT {
		return nil, false, errors.New("Awaiting Activation")
	}
	if data.Status != constants.USERSTATUSACTIVE {
		return nil, false, errors.New("Please Contact Administrator")
	}

	err = s.ValidateOTP(constants.USERLOGIN, login.Mobile, login.OTP)
	if err != nil {
		fmt.Println(err)
		return nil, false, err
	}

	var auth models.Authentication
	auth.UserID = data.ID
	auth.UserName = data.UserName
	auth.Type = data.Type
	auth.Status = data.Status
	auth.Role = data.Role

	token, err := CreateTokenV2(&auth)
	if err != nil {
		log.Println(err.Error())
		return nil, false, errors.New("Error in Generating Token - " + err.Error())
	}
	//data.User.Token = token
	data.Token = token

	return data, true, nil

}
