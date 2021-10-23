package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveStateLiveStock :""
func (s *Service) SaveStateLiveStock(ctx *models.Context, StateLiveStock *models.StateLiveStock) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//StateLiveStock.Code = s.Daos.GetUniqueID(ctx, constants.COLLECTIONStateLiveStock)
	StateLiveStock.Status = constants.STATELIVESTOCKSTATUSACTIVE
	StateLiveStock.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 StateLiveStock.created")
	StateLiveStock.Created = created
	log.Println("b4 StateLiveStock.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveStateLiveStock(ctx, StateLiveStock)
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

//UpdateStateLiveStock : ""
func (s *Service) UpdateStateLiveStock(ctx *models.Context, StateLiveStock *models.StateLiveStock) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateStateLiveStock(ctx, StateLiveStock)
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

//EnableStateLiveStock : ""
func (s *Service) EnableStateLiveStock(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableStateLiveStock(ctx, UniqueID)
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

//DisableStateLiveStock : ""
func (s *Service) DisableStateLiveStock(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableStateLiveStock(ctx, UniqueID)
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

//DeleteStateLiveStock : ""
func (s *Service) DeleteStateLiveStock(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteStateLiveStock(ctx, UniqueID)
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

//GetSingleStateLiveStock :""
func (s *Service) GetSingleStateLiveStock(ctx *models.Context, UniqueID string) (*models.RefStateLiveStock, error) {
	StateLiveStock, err := s.Daos.GetSingleStateLiveStock(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return StateLiveStock, nil
}

//FilterStateLiveStock :""
func (s *Service) FilterStateLiveStock(ctx *models.Context, StateLiveStockfilter *models.StateLiveStockFilter, pagination *models.Pagination) (StateLiveStock []models.RefStateLiveStock, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.FilterStateLiveStock(ctx, StateLiveStockfilter, pagination)
}
