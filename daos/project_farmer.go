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

//SaveProjectFarmer :""
func (d *Daos) SaveProjectFarmer(ctx *models.Context, farmer *models.ProjectFarmer) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTFARMER).InsertOne(ctx.CTX, farmer)
	return err
}

//UpdateProjectFarmer : ""
func (d *Daos) UpdateProjectFarmer(ctx *models.Context, farmer *models.ProjectFarmer) error {

	selector := bson.M{"_id": farmer.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": farmer}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTFARMER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableProjectFarmer :""
func (d *Daos) EnableProjectFarmer(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTFARMERSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTFARMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProjectFarmer :""
func (d *Daos) DisableProjectFarmer(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTFARMERSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTFARMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProjectFarmer :""
func (d *Daos) DeleteProjectFarmer(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTFARMERSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTFARMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleProjectFarmer : ""
func (d *Daos) GetSingleProjectFarmer(ctx *models.Context, UniqueID string) (*models.RefProjectFarmer, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECTFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ProjectFarmers []models.RefProjectFarmer
	var ProjectFarmer *models.RefProjectFarmer
	if err = cursor.All(ctx.CTX, &ProjectFarmers); err != nil {
		return nil, err
	}
	if len(ProjectFarmers) > 0 {
		ProjectFarmer = &ProjectFarmers[0]
	}
	return ProjectFarmer, nil
}

//FilterProjectFarmer : ""
func (d *Daos) FilterProjectFarmer(ctx *models.Context, filter *models.ProjectFarmerFilter, pagination *models.Pagination) ([]models.RefProjectFarmer, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROJECTFARMER).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "project", "_id", "ref.project", "ref.project")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("ProjectFarmer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECTFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ProjectFarmers []models.RefProjectFarmer
	if err = cursor.All(context.TODO(), &ProjectFarmers); err != nil {
		return nil, err
	}
	return ProjectFarmers, nil
}
