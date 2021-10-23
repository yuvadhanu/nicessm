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

//SaveProjectUser :""
func (d *Daos) SaveProjectUser(ctx *models.Context, user *models.ProjectUser) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTUSER).InsertOne(ctx.CTX, user)
	return err
}

//UpdateProjectUser : ""
func (d *Daos) UpdateProjectUser(ctx *models.Context, user *models.ProjectUser) error {

	selector := bson.M{"_id": user.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": user}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTUSER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableProjectUser :""
func (d *Daos) EnableProjectUser(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTUSERSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProjectUser :""
func (d *Daos) DisableProjectUser(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTUSERSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProjectUser :""
func (d *Daos) DeleteProjectUser(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTUSERSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTUSER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleProjectUser : ""
func (d *Daos) GetSingleProjectUser(ctx *models.Context, UniqueID string) (*models.RefProjectUser, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECTUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ProjectUsers []models.RefProjectUser
	var ProjectUser *models.RefProjectUser
	if err = cursor.All(ctx.CTX, &ProjectUsers); err != nil {
		return nil, err
	}
	if len(ProjectUsers) > 0 {
		ProjectUser = &ProjectUsers[0]
	}
	return ProjectUser, nil
}

//FilterProjectUser : ""
func (d *Daos) FilterProjectUser(ctx *models.Context, filter *models.ProjectUserFilter, pagination *models.Pagination) ([]models.RefProjectUser, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.Project) > 0 {
			query = append(query, bson.M{"project": bson.M{"$in": filter.Project}})
		}
		if len(filter.User) > 0 {
			query = append(query, bson.M{"user": bson.M{"$in": filter.User}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROJECTUSER).CountDocuments(ctx.CTX, func() bson.M {
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
	// mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONUSER, "user", "_id", "ref.user", "ref.user")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("ProjectUser query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECTUSER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ProjectUsers []models.RefProjectUser
	if err = cursor.All(context.TODO(), &ProjectUsers); err != nil {
		return nil, err
	}
	return ProjectUsers, nil
}
