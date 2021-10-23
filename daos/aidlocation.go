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

func (d *Daos) SaveAidlocation(ctx *models.Context, aidlocation *models.Aidlocation) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONAIDLOCATION).InsertOne(ctx.CTX, aidlocation)
	if err != nil {
		return err
	}
	aidlocation.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleAidlocation(ctx *models.Context, UniqueID string) (*models.RefAidlocation, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONAIDCATEGORY, "aidCategory", "_id", "ref.aidCategory", "ref.aidCategory")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAIDLOCATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var aidlocations []models.RefAidlocation
	var aidlocation *models.RefAidlocation
	if err = cursor.All(ctx.CTX, &aidlocations); err != nil {
		return nil, err
	}
	if len(aidlocations) > 0 {
		aidlocation = &aidlocations[0]
	}
	return aidlocation, nil
}

func (d *Daos) UpdateAidlocation(ctx *models.Context, aidlocation *models.Aidlocation) error {

	selector := bson.M{"_id": aidlocation.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": aidlocation}
	_, err := ctx.DB.Collection(constants.COLLECTIONAIDLOCATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterAidlocation(ctx *models.Context, aidlocationfilter *models.AidlocationFilter, pagination *models.Pagination) ([]models.RefAidlocation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if aidlocationfilter != nil {

		if len(aidlocationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": aidlocationfilter.Status}})
		}
		//Regex
		if aidlocationfilter.Searchbox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: aidlocationfilter.Searchbox.Name, Options: "xi"}})
		}
		if aidlocationfilter.Searchbox.Description != "" {
			query = append(query, bson.M{"description": primitive.Regex{Pattern: aidlocationfilter.Searchbox.Description, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if aidlocationfilter != nil {
		if aidlocationfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{aidlocationfilter.SortBy: aidlocationfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONAIDLOCATION).CountDocuments(ctx.CTX, func() bson.M {
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
	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONAIDCATEGORY, "aidCategory", "_id", "ref.aidCategory", "ref.aidCategory")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Aidlocation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAIDLOCATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var aidlocations []models.RefAidlocation
	if err = cursor.All(context.TODO(), &aidlocations); err != nil {
		return nil, err
	}
	return aidlocations, nil
}

func (d *Daos) EnableAidlocation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.AIDLOCATIONSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONAIDLOCATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableAidlocation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.AIDLOCATIONSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONAIDLOCATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteAidlocation(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.AIDLOCATIONSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONAIDLOCATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
