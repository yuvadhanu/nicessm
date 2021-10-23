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

//SaveCommoditySubVariety :""
func (d *Daos) SaveCommoditySubVariety(ctx *models.Context, subVariety *models.CommoditySubVariety) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYSUBVARIETY).InsertOne(ctx.CTX, subVariety)
	return err
}

//UpdateCommoditySubVariety : ""
func (d *Daos) UpdateCommoditySubVariety(ctx *models.Context, subVariety *models.CommoditySubVariety) error {

	selector := bson.M{"_id": subVariety.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": subVariety}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYSUBVARIETY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableCommoditySubVariety :""
func (d *Daos) EnableCommoditySubVariety(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYSUBVARIETYSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYSUBVARIETY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCommoditySubVariety :""
func (d *Daos) DisableCommoditySubVariety(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYSUBVARIETYSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYSUBVARIETY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCommoditySubVariety :""
func (d *Daos) DeleteCommoditySubVariety(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYSUBVARIETYSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYSUBVARIETY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleCommoditySubVariety : ""
func (d *Daos) GetSingleCommoditySubVariety(ctx *models.Context, UniqueID string) (*models.RefCommoditySubVariety, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYVARIETY, "commodityVariety", "_id", "ref.commodityVariety", "ref.commodityVariety")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYSUBVARIETY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommoditySubVarietys []models.RefCommoditySubVariety
	var CommoditySubVariety *models.RefCommoditySubVariety
	if err = cursor.All(ctx.CTX, &CommoditySubVarietys); err != nil {
		return nil, err
	}
	if len(CommoditySubVarietys) > 0 {
		CommoditySubVariety = &CommoditySubVarietys[0]
	}
	return CommoditySubVariety, nil
}

//FilterCommoditySubVariety : ""
func (d *Daos) FilterCommoditySubVariety(ctx *models.Context, filter *models.CommoditySubVarietyFilter, pagination *models.Pagination) ([]models.RefCommoditySubVariety, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.CommodityVariety) > 0 {
			query = append(query, bson.M{"commodityVariety": bson.M{"$in": filter.CommodityVariety}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYSUBVARIETY).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYVARIETY, "commodityVariety", "_id", "ref.commodityVariety", "ref.commodityVariety")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("CommoditySubVariety query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYSUBVARIETY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommoditySubVarietys []models.RefCommoditySubVariety
	if err = cursor.All(context.TODO(), &CommoditySubVarietys); err != nil {
		return nil, err
	}
	return CommoditySubVarietys, nil
}
