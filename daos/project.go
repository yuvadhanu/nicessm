package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveProject :""
func (d *Daos) SaveProject(ctx *models.Context, project *models.Project) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).InsertOne(ctx.CTX, project)
	if err != nil {
		return err
	}
	project.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateProject : ""
func (d *Daos) UpdateProject(ctx *models.Context, project *models.Project) error {

	selector := bson.M{"_id": project.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": project}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableProject :""
func (d *Daos) EnableProject(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATUSACTIVE, "activeStatus": constants.PROJECTSTATUSTRUE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProject :""
func (d *Daos) DisableProject(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATUSDISABLED, "activeStatus": constants.PROJECTSTATUSFALSE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProject :""
func (d *Daos) DeleteProject(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleProject : ""
func (d *Daos) GetSingleProject(ctx *models.Context, UniqueID string) (*models.RefProject, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisation", "_id", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECTSTATE, "stateId", "_id", "ref.states", "ref.states")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Projects []models.RefProject
	var Project *models.RefProject
	if err = cursor.All(ctx.CTX, &Projects); err != nil {
		return nil, err
	}
	if len(Projects) > 0 {
		Project = &Projects[0]
	}
	return Project, nil
}

//FilterProject : ""
func (d *Daos) FilterProject(ctx *models.Context, filter *models.ProjectFilter, pagination *models.Pagination) ([]models.RefProject, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.NationalLevel) > 0 {
			query = append(query, bson.M{"nationalLevel": bson.M{"$in": filter.NationalLevel}})
		}
		if len(filter.Organisation) > 0 {
			query = append(query, bson.M{"organisation": bson.M{"$in": filter.Organisation}})
		}

		if filter.BudgetRange != nil {
			if filter.BudgetRange.From != 0 {
				query = append(query, bson.M{"budget": bson.M{"$gte": filter.BudgetRange.From}})

				if filter.BudgetRange.To != 0 {
					query = append(query, bson.M{"budget": bson.M{"$gte": filter.BudgetRange.From, "$lte": filter.BudgetRange.To}})
				}
			}
		}

		if filter.StartDateRange != nil {
			//var sd,ed time.Time
			if filter.StartDateRange.From != nil {
				sd := time.Date(filter.StartDateRange.From.Year(), filter.StartDateRange.From.Month(), filter.StartDateRange.From.Day(), 0, 0, 0, 0, filter.StartDateRange.From.Location())
				ed := time.Date(filter.StartDateRange.From.Year(), filter.StartDateRange.From.Month(), filter.StartDateRange.From.Day(), 23, 59, 59, 0, filter.StartDateRange.From.Location())
				if filter.StartDateRange.To != nil {
					ed = time.Date(filter.StartDateRange.To.Year(), filter.StartDateRange.To.Month(), filter.StartDateRange.To.Day(), 23, 59, 59, 0, filter.StartDateRange.To.Location())
				}
				query = append(query, bson.M{"startDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.EndDateRange != nil {
			//var sd,ed time.Time
			if filter.EndDateRange.From != nil {
				sd := time.Date(filter.EndDateRange.From.Year(), filter.EndDateRange.From.Month(), filter.EndDateRange.From.Day(), 0, 0, 0, 0, filter.EndDateRange.From.Location())
				ed := time.Date(filter.EndDateRange.From.Year(), filter.EndDateRange.From.Month(), filter.EndDateRange.From.Day(), 23, 59, 59, 0, filter.EndDateRange.From.Location())
				if filter.EndDateRange.To != nil {
					ed = time.Date(filter.EndDateRange.To.Year(), filter.EndDateRange.To.Month(), filter.EndDateRange.To.Day(), 23, 59, 59, 0, filter.EndDateRange.To.Location())
				}
				query = append(query, bson.M{"endDate": bson.M{"$gte": sd, "$lte": ed}})

			}
		}
		if filter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.Regex.Name, Options: "xi"}})
		}
		if filter.Regex.Mail != "" {
			query = append(query, bson.M{"mail": primitive.Regex{Pattern: filter.Regex.Mail, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).CountDocuments(ctx.CTX, func() bson.M {
			if query != nil {
				if len(query) > 0 {
					return bson.M{"$and": query}
				}
			}
			return bson.M{}
		}())
		if err != nil {
			log.Println("Error in geting pagination")
		}
		fmt.Println("count", totalCount)
		pagination.Count = int(totalCount)
		d.Shared.PaginationData(pagination)
	}
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisation", "_id", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECTSTATE, "stateId", "_id", "ref.states", "ref.states")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("Project query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Projects []models.RefProject
	if err = cursor.All(context.TODO(), &Projects); err != nil {
		return nil, err
	}
	return Projects, nil
}
