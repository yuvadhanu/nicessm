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

//SaveAsset :""
func (d *Daos) SaveAsset(ctx *models.Context, asset *models.Asset) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONASSET).InsertOne(ctx.CTX, asset)
	if err != nil {
		return err
	}
	asset.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleAsset : ""
func (d *Daos) GetSingleAsset(ctx *models.Context, UniqueID string) (*models.RefAsset, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assets []models.RefAsset
	var asset *models.RefAsset
	if err = cursor.All(ctx.CTX, &assets); err != nil {
		return nil, err
	}
	if len(assets) > 0 {
		asset = &assets[0]
	}
	return asset, nil
}

//UpdateAsset : ""
func (d *Daos) UpdateAsset(ctx *models.Context, asset *models.Asset) error {

	selector := bson.M{"_id": asset.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": asset}
	_, err := ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterAsset : ""
func (d *Daos) FilterAsset(ctx *models.Context, assetFilter *models.AssetFilter, pagination *models.Pagination) ([]models.RefAsset, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if assetFilter != nil {

		if len(assetFilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activestatus": bson.M{"$in": assetFilter.ActiveStatus}})
		}
		if len(assetFilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": assetFilter.Status}})
		}
		//Regex
		if assetFilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: assetFilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if assetFilter != nil {
		if assetFilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{assetFilter.SortBy: assetFilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONASSET).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("asset query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONASSET).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var assets []models.RefAsset
	if err = cursor.All(context.TODO(), &assets); err != nil {
		return nil, err
	}
	return assets, nil
}

//EnableAsset :""
func (d *Daos) EnableAsset(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ASSETSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableAsset :""
func (d *Daos) DisableAsset(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ASSETSTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteAsset :""
func (d *Daos) DeleteAsset(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.ASSETSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONASSET).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
