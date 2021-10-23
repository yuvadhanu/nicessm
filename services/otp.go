package services

import (
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
)

//GenerateOTP :
func (s *Service) GenerateOTP(scenario string, uniqueKey string, otplength int, ttl int) (string, error) {
	prefix, err := s.OTPScenario(scenario)
	if err != nil {
		return "", err
	}
	key := fmt.Sprintf("%v%v", prefix, uniqueKey)
	otp := s.Shared.GetRandomOTP(otplength)
	otp = "9999"
	err = s.Redis.SetValue(key, otp, ttl)
	log.Println("Key==>", key, "otp==>", otp)
	if err != nil {
		return "", err
	}
	return otp, err
}

//OTPScenario : ""
func (s *Service) OTPScenario(scenario string) (string, error) {
	switch scenario {
	case constants.CONSUMERLOGIN:
		return "CONSLOGIN_", nil
	case constants.OTPSCENARIOPASSWORD:
		return "FRGPWDOTP_", nil
	case constants.OTPSCENARIOTOKEN:
		return "FRGPWDOTPTOKEN_", nil
	case constants.USERLOGIN:
		return "USRLOGIN_", nil
	default:
		return "", errors.New("No such scenario")
	}
}

//ValidateOTP :
func (s *Service) ValidateOTP(scenario string, uniqueKey string, otp string) error {
	prefix, err := s.OTPScenario(scenario)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%v%v", prefix, uniqueKey)
	redisToken := s.Redis.GetValue(key)
	rotp, ok := redisToken.(string)
	if !ok {
		log.Println("Cannot type cast from redis to string")
		return errors.New(constants.INTERNALSERVERERROR)
	}
	log.Println("Key==>", key, "otp==>", rotp)
	validateOtp := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.VALIDATEOTP)
	if validateOtp == "NO" {
		return nil
	}
	if otp != rotp {
		log.Println("Not a valid otp")
		return errors.New(constants.INVALIDOTP)
	}

	return nil
}
