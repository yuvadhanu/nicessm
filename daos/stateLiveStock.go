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

//SaveStateLiveStock :""
func (d *Daos) SaveStateLiveStock(ctx *models.Context, StateLiveStock *models.StateLiveStock) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).InsertOne(ctx.CTX, StateLiveStock)
	if err != nil {
		return err
	}
	StateLiveStock.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleStateLiveStock : ""
func (d *Daos) GetSingleStateLiveStock(ctx *models.Context, code string) (*models.RefStateLiveStock, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)
	// //Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateLiveStocks []models.RefStateLiveStock
	var StateLiveStock *models.RefStateLiveStock
	if err = cursor.All(ctx.CTX, &StateLiveStocks); err != nil {
		return nil, err
	}
	if len(StateLiveStocks) > 0 {
		StateLiveStock = &StateLiveStocks[0]
	}
	return StateLiveStock, nil
}

//UpdateStateLiveStock : ""
func (d *Daos) UpdateStateLiveStock(ctx *models.Context, StateLiveStock *models.StateLiveStock) error {

	selector := bson.M{"_id": StateLiveStock.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": StateLiveStock, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterStateLiveStock : ""
func (d *Daos) FilterStateLiveStock(ctx *models.Context, StateLiveStockfilter *models.StateLiveStockFilter, pagination *models.Pagination) ([]models.RefStateLiveStock, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if StateLiveStockfilter != nil {
		if len(StateLiveStockfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": StateLiveStockfilter.ActiveStatus}})
		}
		if len(StateLiveStockfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": StateLiveStockfilter.State}})
		}
		if len(StateLiveStockfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": StateLiveStockfilter.Status}})
		}
		//Regex
		if StateLiveStockfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: StateLiveStockfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).CountDocuments(ctx.CTX, func() bson.M {
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

	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONAGROECOLOGICALZONE, "agroEcologicalZones", "_id", "ref.agroEcologicalZones", "ref.agroEcologicalZones")...)
	// //Aggregation
	d.Shared.BsonToJSONPrintTag("StateLiveStock query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var StateLiveStocks []models.RefStateLiveStock
	if err = cursor.All(context.TODO(), &StateLiveStocks); err != nil {
		return nil, err
	}
	return StateLiveStocks, nil
}

//EnableStateLiveStock :""
func (d *Daos) EnableStateLiveStock(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATELIVESTOCKSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableStateLiveStock :""
func (d *Daos) DisableStateLiveStock(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATELIVESTOCKSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteStateLiveStock :""
func (d *Daos) DeleteStateLiveStock(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATELIVESTOCKSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATELIVESTOCK).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
