package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveModule :""
func (s *Service) SaveModule(ctx *models.Context, module *models.Module) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	module.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMODULE)
	module.Status = constants.ACLMASTERMODULESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 module.created")
	module.Created = created
	log.Println("b4 module.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveModule(ctx, module)
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

//UpdateModule : ""
func (s *Service) UpdateModule(ctx *models.Context, module *models.Module) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateModule(ctx, module)
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

//EnableModule : ""
func (s *Service) EnableModule(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableModule(ctx, UniqueID)
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

//DisableModule : ""
func (s *Service) DisableModule(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableModule(ctx, UniqueID)
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

//DeleteModule : ""
func (s *Service) DeleteModule(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteModule(ctx, UniqueID)
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

//GetSingleModule :""
func (s *Service) GetSingleModule(ctx *models.Context, UniqueID string) (*models.RefModule, error) {
	module, err := s.Daos.GetSingleModule(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return module, nil
}

//FilterModule :""
func (s *Service) FilterModule(ctx *models.Context, modulefilter *models.ModuleFilter, pagination *models.Pagination) (module []models.RefModule, err error) {
	return s.Daos.FilterModule(ctx, modulefilter, pagination)
}

//GetSingleModuleUserType : ""
func (s *Service) GetSingleModuleUserType(ctx *models.Context, userTypeID string) (*models.UserTypeModuleAccess, error) {
	data, err := s.Daos.GetSingleModuleUserType(ctx, userTypeID)
	if err != nil {
		return nil, err
	}
	if len(data.Modules) > 0 {
		for k := range data.Modules {
			if data.Modules[k].Access == nil {
				data.Modules[k].Access = new(models.ACLUserTypeModule)
				data.Modules[k].Access.Check = "No"
			}
		}
	}

	return data, nil
}
