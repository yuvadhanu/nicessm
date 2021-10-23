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

//SaveCommodityCategory :""
func (d *Daos) SaveCommodityCategory(ctx *models.Context, category *models.CommodityCategory) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYCATEGORY).InsertOne(ctx.CTX, category)
	return err
}

//UpdateCommodityCategory : ""
func (d *Daos) UpdateCommodityCategory(ctx *models.Context, category *models.CommodityCategory) error {

	selector := bson.M{"_id": category.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": category}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYCATEGORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableCommodityCategory :""
func (d *Daos) EnableCommodityCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYCATEGORYSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCommodityCategory :""
func (d *Daos) DisableCommodityCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYCATEGORYSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCommodityCategory :""
func (d *Daos) DeleteCommodityCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYCATEGORYSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleCommodityCategory : ""
func (d *Daos) GetSingleCommodityCategory(ctx *models.Context, UniqueID string) (*models.RefCommodityCategory, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommodityCategorys []models.RefCommodityCategory
	var CommodityCategory *models.RefCommodityCategory
	if err = cursor.All(ctx.CTX, &CommodityCategorys); err != nil {
		return nil, err
	}
	if len(CommodityCategorys) > 0 {
		CommodityCategory = &CommodityCategorys[0]
	}
	return CommodityCategory, nil
}

//FilterCommodityCategory : ""
func (d *Daos) FilterCommodityCategory(ctx *models.Context, filter *models.CommodityCategoryFilter, pagination *models.Pagination) ([]models.RefCommodityCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.Classification) > 0 {
			query = append(query, bson.M{"classification": bson.M{"$in": filter.Classification}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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

	//Aggregation
	d.Shared.BsonToJSONPrintTag("CommodityCategory query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommodityCategorys []models.RefCommodityCategory
	if err = cursor.All(context.TODO(), &CommodityCategorys); err != nil {
		return nil, err
	}
	return CommodityCategorys, nil
}
