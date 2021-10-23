package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveKnowlegdeDomain :""
func (s *Service) SaveKnowlegdeDomain(ctx *models.Context, KnowlegdeDomain *models.KnowledgeDomain) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//KnowlegdeDomain.Code = s.Daos.GetUniqueID(ctx, constants.COLLECTIONKnowlegdeDomain)

	KnowlegdeDomain.Status = constants.KNOWLEDGEDOMAINSTATUSACTIVE
	KnowlegdeDomain.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 KnowlegdeDomain.created")
	KnowlegdeDomain.Created = created
	log.Println("b4 KnowlegdeDomain.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveKnowlegdeDomain(ctx, KnowlegdeDomain)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//UpdateKnowlegdeDomain : ""
func (s *Service) UpdateKnowlegdeDomain(ctx *models.Context, KnowlegdeDomain *models.KnowledgeDomain) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateKnowlegdeDomain(ctx, KnowlegdeDomain)
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

//EnableKnowlegdeDomain : ""
func (s *Service) EnableKnowlegdeDomain(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableKnowlegdeDomain(ctx, UniqueID)
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

//DisableKnowlegdeDomain : ""
func (s *Service) DisableKnowlegdeDomain(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableKnowlegdeDomain(ctx, UniqueID)
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

//DeleteKnowlegdeDomain : ""
func (s *Service) DeleteKnowlegdeDomain(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteKnowlegdeDomain(ctx, UniqueID)
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

//GetSingleKnowlegdeDomain :""
func (s *Service) GetSingleKnowledgeDomain(ctx *models.Context, UniqueID string) (*models.RefKnowledgeDomain, error) {
	KnowlegdeDomain, err := s.Daos.GetSingleKnowlegdeDomain(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return KnowlegdeDomain, nil
}

//FilterKnowlegdeDomain :""
func (s *Service) FilterKnowlegdeDomain(ctx *models.Context, KnowledgeDomainfilter *models.KnowledgeDomainFilter, pagination *models.Pagination) (KnowlegdeDomain []models.RefKnowledgeDomain, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterKnowledgeDomain(ctx, KnowledgeDomainfilter, pagination)

}
