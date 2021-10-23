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

//SaveSoilType :""
func (d *Daos) SaveSoilType(ctx *models.Context, soiltype *models.SoilType) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSOILTYPE).InsertOne(ctx.CTX, soiltype)
	if err != nil {
		return err
	}
	soiltype.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//UpdateSoilType : ""
func (d *Daos) UpdateSoilType(ctx *models.Context, soiltype *models.SoilType) error {

	selector := bson.M{"_id": soiltype.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": soiltype}
	_, err := ctx.DB.Collection(constants.COLLECTIONSOILTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableSoilType :""
func (d *Daos) EnableSoilType(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SOILTYPESTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSOILTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableSoilType :""
func (d *Daos) DisableSoilType(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SOILTYPESTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSOILTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteSoilType :""
func (d *Daos) DeleteSoilType(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SOILTYPESTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSOILTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleSoilType : ""
func (d *Daos) GetSingleSoilType(ctx *models.Context, UniqueID string) (*models.RefSoilType, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOILTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var soiltypes []models.RefSoilType
	var soiltype *models.RefSoilType
	if err = cursor.All(ctx.CTX, &soiltypes); err != nil {
		return nil, err
	}
	if len(soiltypes) > 0 {
		soiltype = &soiltypes[0]
	}
	return soiltype, nil
}

//FilterSoilType : ""
func (d *Daos) FilterSoilType(ctx *models.Context, soiltypefilter *models.SoilTypeFilter, pagination *models.Pagination) ([]models.RefSoilType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if soiltypefilter != nil {
		if len(soiltypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": soiltypefilter.Status}})
		}
		if soiltypefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: soiltypefilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if soiltypefilter != nil {
		if soiltypefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{soiltypefilter.SortBy: soiltypefilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSOILTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("soiltype query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSOILTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var soiltypes []models.RefSoilType
	if err = cursor.All(context.TODO(), &soiltypes); err != nil {
		return nil, err
	}
	return soiltypes, nil
}
