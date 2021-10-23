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

//SaveBannedItem :""
func (d *Daos) SaveBannedItem(ctx *models.Context, bannedItem *models.BannedItem) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONBANNEDITEM).InsertOne(ctx.CTX, bannedItem)
	if err != nil {
		return err
	}
	bannedItem.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleBannedItem : ""
func (d *Daos) GetSingleBannedItem(ctx *models.Context, UniqueID string) (*models.RefBannedItem, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBANNEDITEM).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var bannedItems []models.RefBannedItem
	var bannedItem *models.RefBannedItem
	if err = cursor.All(ctx.CTX, &bannedItems); err != nil {
		return nil, err
	}
	if len(bannedItems) > 0 {
		bannedItem = &bannedItems[0]
	}
	return bannedItem, nil
}

//UpdateBannedItem : ""
func (d *Daos) UpdateBannedItem(ctx *models.Context, bannedItem *models.BannedItem) error {

	selector := bson.M{"_id": bannedItem.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": bannedItem}
	_, err := ctx.DB.Collection(constants.COLLECTIONBANNEDITEM).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterBannedItem : ""
func (d *Daos) FilterBannedItem(ctx *models.Context, bannedItemfilter *models.BannedItemFilter, pagination *models.Pagination) ([]models.RefBannedItem, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if bannedItemfilter != nil {

		if len(bannedItemfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": bannedItemfilter.Status}})
		}
		//Regex
		if bannedItemfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: bannedItemfilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if bannedItemfilter != nil {
		if bannedItemfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{bannedItemfilter.SortBy: bannedItemfilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDISEASE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("bannedItem query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBANNEDITEM).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var bannedItems []models.RefBannedItem
	if err = cursor.All(context.TODO(), &bannedItems); err != nil {
		return nil, err
	}
	return bannedItems, nil
}

//EnableBannedItem :""
func (d *Daos) EnableBannedItem(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.BANNEDITEMSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBANNEDITEM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableBannedItem :""
func (d *Daos) DisableBannedItem(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.BANNEDITEMSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBANNEDITEM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteBannedItem :""
func (d *Daos) DeleteBannedItem(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.BANNEDITEMSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBANNEDITEM).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
