package services

import (
	"errors"
	"log"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveACLUserTypeTabMultiple : ""
func (s *Service) SaveACLUserTypeTabMultiple(ctx *models.Context, modules []models.ACLUserTypeTab) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.SaveACLUserTypeTabMultiple(ctx, modules)
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

//GetSingleUserTypeTabAccess : ""
func (s *Service) GetSingleUserTypeTabAccess(ctx *models.Context, userTypeID, moduleID string) (*models.UserTypeTabAccess, error) {
	data, err := s.Daos.GetSingleUserTypeTabAccess(ctx, userTypeID, moduleID)
	if err != nil {
		return nil, err
	}
	if len(data.Module.Tabs) > 0 {
		for k := range data.Module.Tabs {
			if data.Module.Tabs[k].Access == nil {
				data.Module.Tabs[k].Access = new(models.ACLUserTypeTab)
				data.Module.Tabs[k].Access.Check = "No"
			}
		}
	}

	return data, nil
}
