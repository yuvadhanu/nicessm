package services

import (
	"errors"
	"log"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveACLUserTypeModuleMultiple : ""
func (s *Service) SaveACLUserTypeModuleMultiple(ctx *models.Context, modules []models.ACLUserTypeModule) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.Daos.SaveACLUserTypeModuleMultiple(ctx, modules)
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

//FilterACLUserTypeModule : ""
func (s *Service) FilterACLUserTypeModule(ctx *models.Context, filter *models.ACLUserTypeModuleFilter, pagination *models.Pagination) ([]models.RefACLUserTypeModule, error) {
	return s.Daos.FilterACLUserTypeModule(ctx, filter, pagination)
}
