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

//SaveProjectPartner :""
func (d *Daos) SaveProjectPartner(ctx *models.Context, partner *models.ProjectPartner) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTPARTNER).InsertOne(ctx.CTX, partner)
	return err
}

//UpdateProjectPartner : ""
func (d *Daos) UpdateProjectPartner(ctx *models.Context, partner *models.ProjectPartner) error {

	selector := bson.M{"_id": partner.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": partner}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTPARTNER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableProjectPartner :""
func (d *Daos) EnableProjectPartner(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTPARTNERSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTPARTNER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProjectPartner :""
func (d *Daos) DisableProjectPartner(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTPARTNERSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTPARTNER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProjectPartner :""
func (d *Daos) DeleteProjectPartner(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTPARTNERSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTPARTNER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleProjectPartner : ""
func (d *Daos) GetSingleProjectPartner(ctx *models.Context, UniqueID string) (*models.RefProjectPartner, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECTPARTNER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ProjectPartners []models.RefProjectPartner
	var ProjectPartner *models.RefProjectPartner
	if err = cursor.All(ctx.CTX, &ProjectPartners); err != nil {
		return nil, err
	}
	if len(ProjectPartners) > 0 {
		ProjectPartner = &ProjectPartners[0]
	}
	return ProjectPartner, nil
}

//FilterProjectPartner : ""
func (d *Daos) FilterProjectPartner(ctx *models.Context, filter *models.ProjectPartnerFilter, pagination *models.Pagination) ([]models.RefProjectPartner, error) {
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROJECTPARTNER).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("ProjectPartner query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECTPARTNER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ProjectPartners []models.RefProjectPartner
	if err = cursor.All(context.TODO(), &ProjectPartners); err != nil {
		return nil, err
	}
	return ProjectPartners, nil
}
