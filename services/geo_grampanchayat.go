package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveGrampanchayat :""
func (s *Service) SaveGrampanchayat(ctx *models.Context, Grampanchayat *models.GramPanchayat) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//Grampanchayat.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONVILLAGE)
	Grampanchayat.Status = constants.VILLAGESTATUSACTIVE
	Grampanchayat.ActiveStatus = true
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Grampanchayat.created")
	Grampanchayat.Created = created
	log.Println("b4 Grampanchayat.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveGramPanchayat(ctx, Grampanchayat)
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

//UpdateGrampanchayat : ""
func (s *Service) UpdateGrampanchayat(ctx *models.Context, Grampanchayat *models.GramPanchayat) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateGramPanchayat(ctx, Grampanchayat)
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

//EnableGrampanchayat : ""
func (s *Service) EnableGrampanchayat(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableGramPanchayat(ctx, UniqueID)
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

//DisableGrampanchayat : ""
func (s *Service) DisableGrampanchayat(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableGramPanchayat(ctx, UniqueID)
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

//DeleteGrampanchayat : ""
func (s *Service) DeleteGrampanchayat(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteGramPanchayat(ctx, UniqueID)
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

//GetSingleGrampanchayat :""
func (s *Service) GetSingleGrampanchayat(ctx *models.Context, UniqueID string) (*models.RefGramPanchayat, error) {
	Grampanchayat, err := s.Daos.GetSingleGramPanchayat(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Grampanchayat, nil
}

//FilterGrampanchayat :""
func (s *Service) FilterGrampanchayat(ctx *models.Context, Grampanchayatfilter *models.GramPanchayatFilter, pagination *models.Pagination) (Grampanchayat []models.RefGramPanchayat, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterGramPanchayat(ctx, Grampanchayatfilter, pagination)

}
