package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveLiveStockVaccination :""
func (s *Service) SaveLiveStockVaccination(ctx *models.Context, LiveStockVaccination *models.LiveStockVaccination) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//LiveStockVaccination.Code = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLiveStockVaccination)
	LiveStockVaccination.Status = constants.LIVESTOCKVACCINATIONSTATUSACTIVE
	LiveStockVaccination.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 LiveStockVaccination.created")
	LiveStockVaccination.Created = created
	log.Println("b4 LiveStockVaccination.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveLiveStockVaccination(ctx, LiveStockVaccination)
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

//UpdateLiveStockVaccination : ""
func (s *Service) UpdateLiveStockVaccination(ctx *models.Context, LiveStockVaccination *models.LiveStockVaccination) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateLiveStockVaccination(ctx, LiveStockVaccination)
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

//EnableLiveStockVaccination : ""
func (s *Service) EnableLiveStockVaccination(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableLiveStockVaccination(ctx, UniqueID)
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

//DisableLiveStockVaccination : ""
func (s *Service) DisableLiveStockVaccination(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableLiveStockVaccination(ctx, UniqueID)
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

//DeleteLiveStockVaccination : ""
func (s *Service) DeleteLiveStockVaccination(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteLiveStockVaccination(ctx, UniqueID)
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

//GetSingleLiveStockVaccination :""
func (s *Service) GetSingleLiveStockVaccination(ctx *models.Context, UniqueID string) (*models.RefLiveStockVaccination, error) {
	LiveStockVaccination, err := s.Daos.GetSingleLiveStockVaccination(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return LiveStockVaccination, nil
}

//FilterLiveStockVaccination :""
func (s *Service) FilterLiveStockVaccination(ctx *models.Context, LiveStockVaccinationfilter *models.LiveStockVaccinationFilter, pagination *models.Pagination) (LiveStockVaccination []models.RefLiveStockVaccination, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.FilterLiveStockVaccination(ctx, LiveStockVaccinationfilter, pagination)
}
