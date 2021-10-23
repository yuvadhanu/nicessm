package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveProject :""
func (s *Service) SaveProject(ctx *models.Context, Project *models.Project) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	Project.Status = constants.PROJECTSTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	Project.Created = &created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveProject(ctx, Project)
		if dberr != nil {

			return errors.New("Db Error" + dberr.Error())
		}
		var projectStates []models.ProjectState
		for _, v := range Project.StateID {
			projectState := models.ProjectState{
				State:   v,
				Project: Project.ID,
				Created: &created,
				Status:  constants.PROJECTSTATESTATUSACTIVE,
			}
			projectStates = append(projectStates, projectState)
		}
		dberr = s.Daos.SaveProjectMultipleState(ctx, projectStates)
		if dberr != nil {
			return errors.New("Error in Saving Project States" + dberr.Error())
		}
		var projectKDs []models.ProjectKnowledgeDomain
		for _, v := range Project.KnowledgeDomainID {
			projectKD := models.ProjectKnowledgeDomain{
				KnowledgeDomain: v,
				Project:         Project.ID,
				Created:         &created,
				Status:          constants.PROJECTKNOWLEDGEDOMAINSTATUSACTIVE,
			}
			projectKDs = append(projectKDs, projectKD)
		}
		dberr = s.Daos.SaveMultipleProjectKnowledgeDomain(ctx, projectKDs)
		if dberr != nil {
			return errors.New("Error in Saving Project knowledge Domain" + dberr.Error())
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

//UpdateProject : ""
func (s *Service) UpdateProject(ctx *models.Context, project *models.Project) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateProject(ctx, project)
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

//EnableProject : ""
func (s *Service) EnableProject(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableProject(ctx, UniqueID)
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

//DisableProject : ""
func (s *Service) DisableProject(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableProject(ctx, UniqueID)
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

//DeleteProject : ""
func (s *Service) DeleteProject(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteProject(ctx, UniqueID)
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

//GetSingleProject :""
func (s *Service) GetSingleProject(ctx *models.Context, UniqueID string) (*models.RefProject, error) {
	Project, err := s.Daos.GetSingleProject(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Project, nil
}

//FilterProject :""
func (s *Service) FilterProject(ctx *models.Context, Projectfilter *models.ProjectFilter, pagination *models.Pagination) (user []models.RefProject, err error) {
	return s.Daos.FilterProject(ctx, Projectfilter, pagination)
}
