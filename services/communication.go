package services

import (
	"bytes"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"text/template"

	"errors"
	"fmt"
	"log"
	"net/smtp"
	"net/url"
)

//SendSMS : ""
func (s *Service) SendSMS(mobileNo string, msg string) error {
	smsConfig := s.GetSMSConfig()
	log.Println(smsConfig)
	var URL *url.URL
	URL, err := url.Parse(smsConfig.URL)
	if err != nil {
		return errors.New("url const err - " + err.Error())
	}
	parameters := url.Values{}
	parameters.Add("user", smsConfig.User)
	parameters.Add("password", smsConfig.Password)
	parameters.Add("senderid", smsConfig.Senderid)
	parameters.Add("channel", smsConfig.Channel)
	parameters.Add("DCS", fmt.Sprintf("%v", smsConfig.DCS))
	parameters.Add("flashsms", fmt.Sprintf("%v", smsConfig.Flashsms))
	parameters.Add("number", mobileNo)
	parameters.Add("text", msg)
	parameters.Add("route", smsConfig.Route)
	URL.RawQuery = parameters.Encode()
	fmt.Println("URL : ", URL.String())
	resp, err := s.Shared.Get(URL.String(), nil)
	if err != nil {
		return errors.New("api err - " + err.Error())
	}
	log.Println(resp)
	return nil
}

//GetSMSConfig : ""
func (s *Service) GetSMSConfig() *models.SMSConfig {
	smsConfig := new(models.SMSConfig)
	smsConfig.URL = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSURL)
	smsConfig.User = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSUSERNAME)
	smsConfig.Password = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSPASSWORD)
	smsConfig.Senderid = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSSENDERID)
	smsConfig.Channel = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSCHANNEL)
	smsConfig.DCS = s.ConfigReader.GetInt(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSDCS)
	smsConfig.Flashsms = s.ConfigReader.GetInt(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSFLASH)
	smsConfig.Route = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMSROUTE)
	return smsConfig
}

//SendEmailWithTemplate : ""
func (s *Service) SendEmailWithTemplate(subject string, to []string, templateURL string, data interface{}) error {
	fromEmail := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FROMEMAIL)
	fromPassword := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FROMEMAILPASSWORD)
	smtpHost := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMTPHOST)
	smtpPort := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMTPPORT)
	// Authentication.
	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)

	t, err := template.ParseFiles(templateURL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: "+subject+" \n%s\n\n", mimeHeaders)))
	t.Execute(&body, data)
	fmt.Println(fromEmail, fromPassword, smtpHost, smtpPort, to)
	// return nil
	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Email Sent!")
	return nil
}

//SendEmail : ""
func (s *Service) SendEmail(subject string, to []string, message string) error {
	fromEmail := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FROMEMAIL)
	fromPassword := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.FROMEMAILPASSWORD)
	smtpHost := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMTPHOST)
	smtpPort := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.SMTPPORT)
	// Authentication.
	auth := smtp.PlainAuth("", fromEmail, fromPassword, smtpHost)

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: "+subject+" \n%s\n\n", mimeHeaders)))
	fmt.Println(fromEmail, fromPassword, smtpHost, smtpPort, to)
	// return nil
	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, to, []byte(message))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Email Sent!")
	return nil
}
