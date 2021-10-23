package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveProjectKnowledgeDomain :""
func (s *Service) SaveProjectKnowledgeDomain(ctx *models.Context, domain *models.ProjectKnowledgeDomain) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	domain.Status = constants.KNOWLEDGEDOMAINSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	domain.Created = &created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveProjectKnowledgeDomain(ctx, domain)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
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

//UpdateProjectKnowledgeDomain : ""
func (s *Service) UpdateProjectKnowledgeDomain(ctx *models.Context, project *models.ProjectKnowledgeDomain) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateProjectKnowledgeDomain(ctx, project)
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

//EnableProjectKnowledgeDomain : ""
func (s *Service) EnableProjectKnowledgeDomain(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableProjectKnowledgeDomain(ctx, UniqueID)
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

//DisableProjectKnowledgeDomain : ""
func (s *Service) DisableProjectKnowledgeDomain(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableProjectKnowledgeDomain(ctx, UniqueID)
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

//DeleteProjectKnowledgeDomain : ""
func (s *Service) DeleteProjectKnowledgeDomain(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteProjectKnowledgeDomain(ctx, UniqueID)
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

//GetSingleProjectKnowledgeDomain :""
func (s *Service) GetSingleProjectKnowledgeDomain(ctx *models.Context, UniqueID string) (*models.RefProjectKnowledgeDomain, error) {
	ProjectKnowledgeDomain, err := s.Daos.GetSingleProjectKnowledgeDomain(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ProjectKnowledgeDomain, nil
}

//FilterProjectKnowledgeDomain :""
func (s *Service) FilterProjectKnowledgeDomain(ctx *models.Context, ProjectKnowledgeDomainfilter *models.ProjectKnowledgeDomainFilter, pagination *models.Pagination) (user []models.RefProjectKnowledgeDomain, err error) {
	return s.Daos.FilterProjectKnowledgeDomain(ctx, ProjectKnowledgeDomainfilter, pagination)
}
