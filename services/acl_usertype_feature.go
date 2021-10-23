package services

import (
	"errors"
	"log"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveACLUserTypeFeatureMultiple : ""
func (s *Service) SaveACLUserTypeFeatureMultiple(ctx *models.Context, modules []models.ACLUserTypeFeature) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.SaveACLUserTypeFeatureMultiple(ctx, modules)
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

//GetSingleUserTypeFeatureAccess : ""
func (s *Service) GetSingleUserTypeFeatureAccess(ctx *models.Context, userTypeID, moduleID string) (*models.UserTypeFeatureAccess, error) {
	data, err := s.Daos.GetSingleUserTypeFeatureAccess(ctx, userTypeID, moduleID)
	if err != nil {
		return nil, err
	}
	if len(data.Module.Features) > 0 {
		for k := range data.Module.Features {
			if data.Module.Features[k].Access == nil {
				data.Module.Features[k].Access = new(models.ACLUserTypeFeature)
				data.Module.Features[k].Access.Check = "No"
			}
		}
	}

	return data, nil
}
