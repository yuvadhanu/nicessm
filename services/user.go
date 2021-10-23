package services

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUser :""
func (s *Service) SaveUserwithtransaction(ctx *models.Context, user *models.User) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.SaveUserwithouttransaction(ctx, user)
		if err != nil {
			return err
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}

	return nil
}

//SaveUser :""
func (s *Service) SaveUserwithouttransaction(ctx *models.Context, user *models.User) error {
	log.Println("transaction start")
	//Start Transaction

	user.UserName = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSER)
	user.Status = constants.USERSTATUSACTIVE
	user.Password = "#nature32" //Default Password
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 user.created")
	user.Created = created
	log.Println("b4 user.created")
	prod, er2 := s.Daos.GetactiveProductConfig(ctx, true)
	if er2 != nil {
		return er2
	}
	if prod == nil {
		return errors.New("product config is nil")
	}
	if prod.ValidateUserregistration {
		err := s.ValidateUser(ctx, user)
		if err != nil {
			return err
		}

	}
	dberr := s.Daos.SaveUser(ctx, user)
	if dberr != nil {

		return errors.New("Db Error" + dberr.Error())
	}

	return nil
}
func (s *Service) ValidateUser(ctx *models.Context, user *models.User) error {
	switch user.Type {
	case constants.USERTYPECALLCENTERAGENT:
		_, err := user.ValidateCallcenterAgent()
		if err != nil {
			return err
		}
	case constants.USERTYPECONTENTCREATOR:
		_, err := user.ValidateContentCreator()
		if err != nil {
			return err
		}
	case constants.USERTYPECONTENTMANAGER:
		_, err := user.ValidateContentManager()
		if err != nil {
			return err
		}
	case constants.USERTYPECONTENTPROVIDER:
		_, err := user.ValidateContentProvider()
		if err != nil {
			return err
		}
	case constants.USERTYPECONTENTDISSEMINATOR:
		_, err := user.ValidateContentDisseminator()
		if err != nil {
			return err
		}
	case constants.USERTYPEFIELDAGENT:
		_, err := user.ValidateFieldAgent()
		if err != nil {
			return err
		}
	case constants.USERTYPELANGUAGETRANSLATOR:
		_, err := user.ValidateLanguageTranslator()
		if err != nil {
			return err
		}
	case constants.USERTYPELANGUAGEAPPROVER:
		_, err := user.ValidateLanguageApprover()
		if err != nil {
			return err
		}
	case constants.USERTYPEMANAGEMENT:
		_, err := user.ValidateManagement()
		if err != nil {
			return err
		}
	case constants.USERTYPEMODERATOR:
		_, err := user.ValidateModerator()
		if err != nil {
			return err
		}
	case constants.USERTYPESUBJECTMATTEREXPERT:
		_, err := user.ValidateSubjectMatterExpert()
		if err != nil {
			return err
		}
	case constants.USERTYPESYSTEMADMIN:
		_, err := user.ValidateSystemAdmin()
		if err != nil {
			return err
		}
	case constants.USERTYPEVISTORVIEWER:
		_, err := user.ValidateSystemAdmin()
		if err != nil {
			return err
		}
	case constants.USERTYPETRAINER:
		_, err := user.ValidateTrainer()
		if err != nil {
			return err
		}
	case constants.USERTYPEFIELDAGENTLEAD:
		_, err := user.ValidateFieldAgentLead()
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid user type")
	}
	return nil
}

//UpdateUser : ""
func (s *Service) UpdateUser(ctx *models.Context, user *models.User) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUser(ctx, user)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//EnableUser : ""
func (s *Service) EnableUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableUser(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DisableUser : ""
func (s *Service) DisableUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableUser(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DeleteUser : ""
func (s *Service) DeleteUser(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteUser(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//GetSingleUser :""
func (s *Service) GetSingleUser(ctx *models.Context, UniqueID string) (*models.RefUser, error) {
	user, err := s.Daos.GetSingleUser(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//FilterUser :""
func (s *Service) FilterUser(ctx *models.Context, userfilter *models.UserFilter, pagination *models.Pagination) (user []models.RefUser, err error) {
	return s.Daos.FilterUser(ctx, userfilter, pagination)
}

//ResetUserPassword : ""
func (s *Service) ResetUserPassword(ctx *models.Context, userName string) error {
	return s.Daos.ResetUserPassword(ctx, userName, "#nature32")
}

//ChangePassword : ""
func (s *Service) ChangePassword(ctx *models.Context, cp *models.UserChangePassword) (bool, string, error) {
	user, err := s.Daos.GetSingleUser(ctx, cp.UserName)
	if err != nil {
		return false, "", err
	}
	if user.Password != cp.OldPassword {
		return false, "Wrong Password", nil
	}
	err = s.Daos.ResetUserPassword(ctx, cp.UserName, cp.NewPassword)
	if err != nil {
		return false, "", err
	}

	d := make(map[string]interface{})
	d["UserName"] = user.UserName
	d["Name"] = user.Name
	d["URL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.APIBASEURL) + user.Profile
	d["LoginURL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.BASEURL) + s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.LOGINURL)
	d["ContactUsURL"] = s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.BASEURL) + s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV)+"."+constants.CONTACTUSURL)

	templateID := "successfull_registration.html"
	err = s.SendEmailWithTemplate("Kochas Municipality - Password Changed Successfully", []string{"solomon2261993@gmail.com"}, "templates/"+templateID, d)
	if err != nil {
		log.Println("Email not sent - " + err.Error())
		// return errors.New("Unable to send email - " + err.Error())
	}
	return true, "", nil
}

//ForgetPasswordValidateOTP : ""
func (s *Service) ForgetPasswordValidateOTP(ctx *models.Context, UniqueID string, otp string) (string, error) {
	user, err := s.Daos.GetSingleUser(ctx, UniqueID)
	if err != nil {
		return "", err
	}
	err = s.ValidateOTP(constants.OTPSCENARIOPASSWORD, user.Mobile, otp)
	if err != nil {
		return "", err
	}
	token, err := s.GenerateOTP(constants.OTPSCENARIOTOKEN, user.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return "", err
	}
	sEnc := b64.StdEncoding.EncodeToString([]byte(token))

	fmt.Println(sEnc)
	return sEnc, nil
}

//ForgetPasswordGenerateOTP : ""
func (s *Service) ForgetPasswordGenerateOTP(ctx *models.Context, UniqueID string) error {
	user, err := s.Daos.GetSingleUser(ctx, UniqueID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user is nil")
	}
	otp, err := s.GenerateOTP(constants.OTPSCENARIOPASSWORD, user.Mobile, constants.PHONEOTPLENGTH, constants.OTPEXPIRY)
	if err != nil {
		return err
	}
	msg := "Use otp " + otp + " for municipal forget password. municipal doesnt ask otp to be shared with anyone"
	err = s.SendSMS(user.Mobile, msg)
	fmt.Println(err)
	return nil
}

//UpdateUser : ""
func (s *Service) PasswordUpdate(ctx *models.Context, user *models.RefPassword) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.PasswordUpdate(ctx, user)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//UpdateUser : ""
func (s *Service) UserCollectionLimit(ctx *models.Context, UserName string, user *models.CollectionLimit) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UserCollectionLimit(ctx, UserName, user)
		if err != nil {
			return err
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil

	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return err
	}
	return nil
}
