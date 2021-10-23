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

//SaveCommonLand :""
func (d *Daos) SaveCommonLand(ctx *models.Context, CommonLand *models.CommonLand) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLAND).InsertOne(ctx.CTX, CommonLand)
	if err != nil {
		return err
	}
	CommonLand.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleCommonLand : ""
func (d *Daos) GetSingleCommonLand(ctx *models.Context, code string) (*models.RefCommonLand, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLAND).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommonLands []models.RefCommonLand
	var CommonLand *models.RefCommonLand
	if err = cursor.All(ctx.CTX, &CommonLands); err != nil {
		return nil, err
	}
	if len(CommonLands) > 0 {
		CommonLand = &CommonLands[0]
	}
	return CommonLand, nil
}

//UpdateCommonLand : ""
func (d *Daos) UpdateCommonLand(ctx *models.Context, CommonLand *models.CommonLand) error {
	selector := bson.M{"_id": CommonLand.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": CommonLand, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLAND).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterCommonLand : ""
func (d *Daos) FilterCommonLand(ctx *models.Context, CommonLandfilter *models.CommonLandFilter, pagination *models.Pagination) ([]models.RefCommonLand, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if CommonLandfilter != nil {

		if len(CommonLandfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": CommonLandfilter.ActiveStatus}})
		}
		if len(CommonLandfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": CommonLandfilter.Status}})
		}
		//Regex
		if CommonLandfilter.Regex.Type != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: CommonLandfilter.Regex.Type, Options: "xi"}})
		}
		if CommonLandfilter.Regex.KhasraNumber != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: CommonLandfilter.Regex.KhasraNumber, Options: "xi"}})
		}
		if CommonLandfilter.Regex.ParcelNumber != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: CommonLandfilter.Regex.ParcelNumber, Options: "xi"}})
		}
		if CommonLandfilter.Regex.Ownership != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: CommonLandfilter.Regex.Ownership, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLAND).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("CommonLand query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMONLAND).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var CommonLands []models.RefCommonLand
	if err = cursor.All(context.TODO(), &CommonLands); err != nil {
		return nil, err
	}
	return CommonLands, nil
}

//EnableCommonLand :""
func (d *Daos) EnableCommonLand(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMONLANDSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMONLAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCommonLand :""
func (d *Daos) DisableCommonLand(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMONLANDSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMONLAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCommonLand :""
func (d *Daos) DeleteCommonLand(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMONLANDSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMONLAND).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
