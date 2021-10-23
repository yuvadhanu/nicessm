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

//SaveProjectState :""
func (d *Daos) SaveProjectState(ctx *models.Context, domain *models.ProjectState) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONPROJECTSTATE).InsertOne(ctx.CTX, domain)
	if err != nil {
		return err
	}
	domain.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//SaveProjectState :""
func (d *Daos) SaveProjectMultipleState(ctx *models.Context, projectState []models.ProjectState) error {
	var temProjectState []interface{}
	for _, v := range projectState {
		temProjectState = append(temProjectState, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTSTATE).InsertMany(ctx.CTX, temProjectState)
	return err
}

//UpdateProjectState : ""
func (d *Daos) UpdateProjectState(ctx *models.Context, domain *models.ProjectState) error {

	selector := bson.M{"_id": domain.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": domain}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTSTATE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableProjectState :""
func (d *Daos) EnableProjectState(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATESTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTSTATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProjectState :""
func (d *Daos) DisableProjectState(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATESTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTSTATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProjectState :""
func (d *Daos) DeleteProjectState(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTSTATESTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTSTATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleProjectState : ""
func (d *Daos) GetSingleProjectState(ctx *models.Context, UniqueID string) (*models.RefProjectState, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "project", "_id", "ref.project", "ref.project")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECTSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ProjectStates []models.RefProjectState
	var ProjectState *models.RefProjectState
	if err = cursor.All(ctx.CTX, &ProjectStates); err != nil {
		return nil, err
	}
	if len(ProjectStates) > 0 {
		ProjectState = &ProjectStates[0]
	}
	return ProjectState, nil
}

//FilterProjectState : ""
func (d *Daos) FilterProjectState(ctx *models.Context, filter *models.ProjectStateFilter, pagination *models.Pagination) ([]models.RefProjectState, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.Project) > 0 {
			query = append(query, bson.M{"project": bson.M{"$in": filter.Project}})
		}
		if len(filter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": filter.State}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROJECTSTATE).CountDocuments(ctx.CTX, func() bson.M {
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
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "project", "_id", "ref.project", "ref.project")...)

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("ProjectState query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECTSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ProjectStates []models.RefProjectState
	if err = cursor.All(context.TODO(), &ProjectStates); err != nil {
		return nil, err
	}
	return ProjectStates, nil
}
