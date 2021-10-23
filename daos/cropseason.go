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

func (d *Daos) SaveCropseason(ctx *models.Context, cropseason *models.Cropseason) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONCROPSEASON).InsertOne(ctx.CTX, cropseason)
	if err != nil {
		return err
	}
	cropseason.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleCropseason(ctx *models.Context, UniqueID string) (*models.RefCropseason, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCROPSEASON).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var cropseasons []models.RefCropseason
	var cropseason *models.RefCropseason
	if err = cursor.All(ctx.CTX, &cropseasons); err != nil {
		return nil, err
	}
	if len(cropseasons) > 0 {
		cropseason = &cropseasons[0]
	}
	return cropseason, nil
}

func (d *Daos) UpdateCropseason(ctx *models.Context, cropseason *models.Cropseason) error {

	selector := bson.M{"_id": cropseason.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": cropseason}
	_, err := ctx.DB.Collection(constants.COLLECTIONCROPSEASON).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterCropseason(ctx *models.Context, cropseasonfilter *models.CropseasonFilter, pagination *models.Pagination) ([]models.RefCropseason, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if cropseasonfilter != nil {

		if len(cropseasonfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activestatus": bson.M{"$in": cropseasonfilter.ActiveStatus}})
		}
		if len(cropseasonfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": cropseasonfilter.Status}})
		}
		if len(cropseasonfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": cropseasonfilter.State}})
		}
		//Regex
		if cropseasonfilter.Searchbox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: cropseasonfilter.Searchbox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if cropseasonfilter != nil {
		if cropseasonfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{cropseasonfilter.SortBy: cropseasonfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCROPSEASON).CountDocuments(ctx.CTX, func() bson.M {
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
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("language query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCROPSEASON).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var cropseason []models.RefCropseason
	if err = cursor.All(context.TODO(), &cropseason); err != nil {
		return nil, err
	}
	return cropseason, nil
}

func (d *Daos) EnableCropseason(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.CROPSEASONSTATUSTRUE, "status": constants.CROPSEASONSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCROPSEASON).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableCropseason(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.CROPSEASONSTATUSFALSE, "status": constants.CROPSEASONSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCROPSEASON).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteCropseason(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.CROPSEASONSTATUSFALSE, "status": constants.CROPSEASONSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCROPSEASON).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
