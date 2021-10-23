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

func (d *Daos) SaveAgroEcologicalZone(ctx *models.Context, agroEcologicalZone *models.AgroEcologicalZone) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONAGROECOLOGICALZONE).InsertOne(ctx.CTX, agroEcologicalZone)
	if err != nil {
		return err
	}
	agroEcologicalZone.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleAgroEcologicalZone(ctx *models.Context, UniqueID string) (*models.RefAgroEcologicalZone, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAGROECOLOGICALZONE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var agroEcologicalZones []models.RefAgroEcologicalZone
	var agroEcologicalZone *models.RefAgroEcologicalZone
	if err = cursor.All(ctx.CTX, &agroEcologicalZones); err != nil {
		return nil, err
	}
	if len(agroEcologicalZones) > 0 {
		agroEcologicalZone = &agroEcologicalZones[0]
	}
	return agroEcologicalZone, nil
}

func (d *Daos) UpdateAgroEcologicalZone(ctx *models.Context, agroEcologicalZone *models.AgroEcologicalZone) error {

	selector := bson.M{"_id": agroEcologicalZone.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": agroEcologicalZone}
	_, err := ctx.DB.Collection(constants.COLLECTIONAGROECOLOGICALZONE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterAgroEcologicalZone(ctx *models.Context, agroEcologicalZonefilter *models.AgroEcologicalZoneFilter, pagination *models.Pagination) ([]models.RefAgroEcologicalZone, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if agroEcologicalZonefilter != nil {

		if len(agroEcologicalZonefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activestatus": bson.M{"$in": agroEcologicalZonefilter.ActiveStatus}})
		}
		if len(agroEcologicalZonefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": agroEcologicalZonefilter.Status}})
		}
		//Regex
		if agroEcologicalZonefilter.Searchbox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: agroEcologicalZonefilter.Searchbox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if agroEcologicalZonefilter != nil {
		if agroEcologicalZonefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{agroEcologicalZonefilter.SortBy: agroEcologicalZonefilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONAGROECOLOGICALZONE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAGROECOLOGICALZONE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var agroEcologicalZone []models.RefAgroEcologicalZone
	if err = cursor.All(context.TODO(), &agroEcologicalZone); err != nil {
		return nil, err
	}
	return agroEcologicalZone, nil
}

func (d *Daos) EnableAgroEcologicalZone(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.AGROECOLOGICALZONESTATUSTRUE, "status": constants.AGROECOLOGICALZONESTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONAGROECOLOGICALZONE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableAgroEcologicalZone(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.AGROECOLOGICALZONESTATUSFALSE, "status": constants.AGROECOLOGICALZONESTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONAGROECOLOGICALZONE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteAgroEcologicalZone(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.AGROECOLOGICALZONESTATUSFALSE, "status": constants.AGROECOLOGICALZONESTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONAGROECOLOGICALZONE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
