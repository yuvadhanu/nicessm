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

//SaveCommodityVariety :""
func (d *Daos) SaveCommodityVariety(ctx *models.Context, variety *models.CommodityVariety) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYVARIETY).InsertOne(ctx.CTX, variety)
	return err
}

//UpdateCommodityVariety : ""
func (d *Daos) UpdateCommodityVariety(ctx *models.Context, variety *models.CommodityVariety) error {

	selector := bson.M{"_id": variety.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": variety}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYVARIETY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableCommodityVariety :""
func (d *Daos) EnableCommodityVariety(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYVARIETYSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYVARIETY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCommodityVariety :""
func (d *Daos) DisableCommodityVariety(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYVARIETYSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYVARIETY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCommodityVariety :""
func (d *Daos) DeleteCommodityVariety(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYVARIETYSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYVARIETY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleCommodityVariety : ""
func (d *Daos) GetSingleCommodityVariety(ctx *models.Context, UniqueID string) (*models.RefCommodityVariety, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYVARIETY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommodityVarietys []models.RefCommodityVariety
	var CommodityVariety *models.RefCommodityVariety
	if err = cursor.All(ctx.CTX, &CommodityVarietys); err != nil {
		return nil, err
	}
	if len(CommodityVarietys) > 0 {
		CommodityVariety = &CommodityVarietys[0]
	}
	return CommodityVariety, nil
}

//FilterCommodityVariety : ""
func (d *Daos) FilterCommodityVariety(ctx *models.Context, filter *models.CommodityVarietyFilter, pagination *models.Pagination) ([]models.RefCommodityVariety, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.Commodity) > 0 {
			query = append(query, bson.M{"commodity": bson.M{"$in": filter.Commodity}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYVARIETY).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)
	mainPipeline = append(mainPipeline, d.CommonLookupAdvancedArray(constants.COLLECTIONCOMMODITYSUBVARIETY, bson.M{"commodityVariety": "$_id"},
		[]bson.M{
			{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				{"$eq": []string{"$status", constants.COMMODITYSUBVARIETYSTATUSACTIVE}},
				{"$eq": []string{"$commodityVariety", "$$commodityVariety"}},
			}}}},
		},
		"ref.subVariety",
		"ref.subVariety")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("CommodityVariety query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYVARIETY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommodityVarietys []models.RefCommodityVariety
	if err = cursor.All(context.TODO(), &CommodityVarietys); err != nil {
		return nil, err
	}
	return CommodityVarietys, nil
}
