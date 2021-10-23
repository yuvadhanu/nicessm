package services

import (
	"errors"
	"log"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveACLUserTypeMenuMultiple : ""
func (s *Service) SaveACLUserTypeMenuMultiple(ctx *models.Context, modules []models.ACLUserTypeMenu) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.SaveACLUserTypeMenuMultiple(ctx, modules)
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

//GetSingleUserTypeMenuAccess : ""
func (s *Service) GetSingleUserTypeMenuAccess(ctx *models.Context, userTypeID, moduleID string) (*models.UserTypeMenuAccess, error) {
	data, err := s.Daos.GetSingleUserTypeMenuAccess(ctx, userTypeID, moduleID)
	if err != nil {
		return nil, err
	}
	if len(data.Module.Menus) > 0 {
		for k := range data.Module.Menus {
			if data.Module.Menus[k].Access == nil {
				data.Module.Menus[k].Access = new(models.ACLUserTypeMenu)
				data.Module.Menus[k].Access.Check = "No"
			}
		}
	}

	return data, nil
}
