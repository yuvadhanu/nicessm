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
)

//SaveFeature :""
func (d *Daos) SaveFeature(ctx *models.Context, Feature *models.Feature) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONFEATURE).InsertOne(ctx.CTX, Feature)
	return err
}

//GetSingleFeature : ""
func (d *Daos) GetSingleFeature(ctx *models.Context, UniqueID string) (*models.RefFeature, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFEATURE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var features []models.RefFeature
	var Feature *models.RefFeature
	if err = cursor.All(ctx.CTX, &features); err != nil {
		return nil, err
	}
	if len(features) > 0 {
		Feature = &features[0]
	}
	return Feature, nil
}

//UpdateFeature : ""
func (d *Daos) UpdateFeature(ctx *models.Context, Feature *models.Feature) error {
	selector := bson.M{"uniqueId": Feature.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Feature, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFEATURE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterFeature : ""
func (d *Daos) FilterFeature(ctx *models.Context, featurefilter *models.FeatureFilter, pagination *models.Pagination) ([]models.RefFeature, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if featurefilter != nil {

		if len(featurefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": featurefilter.Status}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFEATURE).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFEATURE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var features []models.RefFeature
	if err = cursor.All(context.TODO(), &features); err != nil {
		return nil, err
	}
	return features, nil
}

//EnableFeature :""
func (d *Daos) EnableFeature(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERFEATURESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFEATURE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFeature :""
func (d *Daos) DisableFeature(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERFEATURESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFEATURE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFeature :""
func (d *Daos) DeleteFeature(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERFEATURESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFEATURE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
