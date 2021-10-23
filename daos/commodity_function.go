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

//SaveCommodityFunction :""
func (d *Daos) SaveCommodityFunction(ctx *models.Context, function *models.CommodityFunction) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYFUNCTION).InsertOne(ctx.CTX, function)
	return err
}

//UpdateCommodityFunction : ""
func (d *Daos) UpdateCommodityFunction(ctx *models.Context, function *models.CommodityFunction) error {

	selector := bson.M{"_id": function.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": function}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYFUNCTION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableCommodityFunction :""
func (d *Daos) EnableCommodityFunction(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYFUNCTIONSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYFUNCTION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCommodityFunction :""
func (d *Daos) DisableCommodityFunction(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYFUNCTIONSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYFUNCTION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCommodityFunction :""
func (d *Daos) DeleteCommodityFunction(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYFUNCTIONSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYFUNCTION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleCommodityFunction : ""
func (d *Daos) GetSingleCommodityFunction(ctx *models.Context, UniqueID string) (*models.RefCommodityFunction, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYCATEGORY, "category", "_id", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYFUNCTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommodityFunctions []models.RefCommodityFunction
	var CommodityFunction *models.RefCommodityFunction
	if err = cursor.All(ctx.CTX, &CommodityFunctions); err != nil {
		return nil, err
	}
	if len(CommodityFunctions) > 0 {
		CommodityFunction = &CommodityFunctions[0]
	}
	return CommodityFunction, nil
}

//FilterCommodityFunction : ""
func (d *Daos) FilterCommodityFunction(ctx *models.Context, filter *models.CommodityFunctionFilter, pagination *models.Pagination) ([]models.RefCommodityFunction, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.Category) > 0 {
			query = append(query, bson.M{"category": bson.M{"$in": filter.Category}})
		}
		if filter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: filter.SearchBox.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYFUNCTION).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYCATEGORY, "category", "_id", "ref.category", "ref.category")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("CommodityFunction query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYFUNCTION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommodityFunctions []models.RefCommodityFunction
	if err = cursor.All(context.TODO(), &CommodityFunctions); err != nil {
		return nil, err
	}
	return CommodityFunctions, nil
}
