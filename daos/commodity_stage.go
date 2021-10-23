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

//SaveCommodityStage :""
func (d *Daos) SaveCommodityStage(ctx *models.Context, stage *models.CommodityStage) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYSTAGE).InsertOne(ctx.CTX, stage)
	return err
}

//UpdateCommodityStage : ""
func (d *Daos) UpdateCommodityStage(ctx *models.Context, stage *models.CommodityStage) error {

	selector := bson.M{"_id": stage.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": stage}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYSTAGE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableCommodityStage :""
func (d *Daos) EnableCommodityStage(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYSTAGESTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYSTAGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCommodityStage :""
func (d *Daos) DisableCommodityStage(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYSTAGESTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYSTAGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCommodityStage :""
func (d *Daos) DeleteCommodityStage(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYSTAGESTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITYSTAGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleCommodityStage : ""
func (d *Daos) GetSingleCommodityStage(ctx *models.Context, UniqueID string) (*models.RefCommodityStage, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYCATEGORY, "category", "_id", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYSTAGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommodityStages []models.RefCommodityStage
	var CommodityStage *models.RefCommodityStage
	if err = cursor.All(ctx.CTX, &CommodityStages); err != nil {
		return nil, err
	}
	if len(CommodityStages) > 0 {
		CommodityStage = &CommodityStages[0]
	}
	return CommodityStage, nil
}

//FilterCommodityStage : ""
func (d *Daos) FilterCommodityStage(ctx *models.Context, filter *models.CommodityStageFilter, pagination *models.Pagination) ([]models.RefCommodityStage, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.Category) > 0 {
			query = append(query, bson.M{"category": bson.M{"$in": filter.Category}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYSTAGE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("CommodityStage query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITYSTAGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommodityStages []models.RefCommodityStage
	if err = cursor.All(context.TODO(), &CommodityStages); err != nil {
		return nil, err
	}
	return CommodityStages, nil
}
