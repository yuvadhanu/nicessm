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

func (d *Daos) SaveMarket(ctx *models.Context, market *models.Market) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONMARKET).InsertOne(ctx.CTX, market)
	if err != nil {
		return err
	}
	market.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleMarket(ctx *models.Context, UniqueID string) (*models.RefMarket, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMARKET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var markets []models.RefMarket
	var market *models.RefMarket
	if err = cursor.All(ctx.CTX, &markets); err != nil {
		return nil, err
	}
	if len(markets) > 0 {
		market = &markets[0]
	}
	return market, nil
}

func (d *Daos) UpdateMarket(ctx *models.Context, market *models.Market) error {

	selector := bson.M{"_id": market.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": market}
	_, err := ctx.DB.Collection(constants.COLLECTIONMARKET).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterMarket(ctx *models.Context, Marketfilter *models.MarketFilter, pagination *models.Pagination) ([]models.RefMarket, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Marketfilter != nil {

		if len(Marketfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activestatus": bson.M{"$in": Marketfilter.ActiveStatus}})
		}
		if len(Marketfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Marketfilter.Status}})
		}
		//Regex
		if Marketfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Marketfilter.Regex.Name, Options: "xi"}})
		}
		if len(Marketfilter.Address.StateCode) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": Marketfilter.Address.StateCode}})
		}
		if len(Marketfilter.Address.DistrictCode) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": Marketfilter.Address.DistrictCode}})
		}
		if len(Marketfilter.Address.VillageCode) > 0 {
			query = append(query, bson.M{"village": bson.M{"$in": Marketfilter.Address.VillageCode}})
		}
		if len(Marketfilter.Address.BlockCode) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": Marketfilter.Address.BlockCode}})
		}
		if len(Marketfilter.Address.GramPanchayathCode) > 0 {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$in": Marketfilter.Address.GramPanchayathCode}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if Marketfilter != nil {
		if Marketfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{Marketfilter.SortBy: Marketfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMARKET).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("language query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMARKET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var markets []models.RefMarket
	if err = cursor.All(context.TODO(), &markets); err != nil {
		return nil, err
	}
	return markets, nil
}

func (d *Daos) EnableMarket(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.MARKETSTATUSTRUE, "status": constants.MARKETSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONMARKET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableMarket(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.MARKETSTATUSFALSE, "status": constants.MARKETSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONMARKET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteMarket(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.MARKETSTATUSFALSE, "status": constants.MARKETSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONMARKET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
