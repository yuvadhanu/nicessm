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

//SaveDisease :""
func (d *Daos) SaveDistrictWeatherData(ctx *models.Context, districtweatherdata *models.DistrictWeatherData) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDISTRICWEATHERDATA).InsertOne(ctx.CTX, districtweatherdata)
	if err != nil {
		return err
	}
	districtweatherdata.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDistrictWeatherData : ""
func (d *Daos) GetSingleDistrictWeatherData(ctx *models.Context, UniqueID string) (*models.RefDistrictWeatherData, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districtweatherdatas []models.RefDistrictWeatherData
	var districtweatherdata *models.RefDistrictWeatherData
	if err = cursor.All(ctx.CTX, &districtweatherdatas); err != nil {
		return nil, err
	}
	if len(districtweatherdatas) > 0 {
		districtweatherdata = &districtweatherdatas[0]
	}
	return districtweatherdata, nil
}

//UpdateDistrictWeatherData : ""
func (d *Daos) UpdateDistrictWeatherData(ctx *models.Context, districtweatherdata *models.DistrictWeatherData) error {

	selector := bson.M{"_id": districtweatherdata.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": districtweatherdata}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICWEATHERDATA).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDistrictWeatherData : ""
func (d *Daos) FilterDistrictWeatherData(ctx *models.Context, districtweatherdatafilter *models.DistrictWeatherDataFilter, pagination *models.Pagination) ([]models.RefDistrictWeatherData, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if districtweatherdatafilter != nil {

		if len(districtweatherdatafilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": districtweatherdatafilter.ActiveStatus}})
		}
		if len(districtweatherdatafilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": districtweatherdatafilter.Status}})
		}
		if len(districtweatherdatafilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": districtweatherdatafilter.District}})
		}
		if len(districtweatherdatafilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": districtweatherdatafilter.State}})
		}
		//Regex
		if districtweatherdatafilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: districtweatherdatafilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if districtweatherdatafilter != nil {
		if districtweatherdatafilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{districtweatherdatafilter.SortBy: districtweatherdatafilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDISTRICWEATHERDATA).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("districtweatherdata query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICWEATHERDATA).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districtweatherdatas []models.RefDistrictWeatherData
	if err = cursor.All(context.TODO(), &districtweatherdatas); err != nil {
		return nil, err
	}
	return districtweatherdatas, nil
}

//EnableDistrictWeatherData :""
func (d *Daos) EnableDistrictWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICWEATHERDATASTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDistrictWeatherData :""
func (d *Daos) DisableDistrictWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICWEATHERDATASTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDistrictWeatherData :""
func (d *Daos) DeleteDistrictWeatherData(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICWEATHERDATASTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICWEATHERDATA).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
