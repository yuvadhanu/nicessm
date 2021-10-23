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

//SaveDistrict :""
func (d *Daos) SaveDistrict(ctx *models.Context, district *models.District) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).InsertOne(ctx.CTX, district)
	return err
}

//GetSingleDistrict : ""
func (d *Daos) GetSingleDistrict(ctx *models.Context, code string) (*models.RefDistrict, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSOILTYPE, "soilTypes", "_id", "ref.soilTypes", "ref.soilTypes")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONAGROECOLOGICALZONE, "agroEcologicalZones", "_id", "ref.agroEcologicalZones", "ref.agroEcologicalZones")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districts []models.RefDistrict
	var district *models.RefDistrict
	if err = cursor.All(ctx.CTX, &districts); err != nil {
		return nil, err
	}
	if len(districts) > 0 {
		district = &districts[0]
	}
	return district, nil
}

//UpdateDistrict : ""
func (d *Daos) UpdateDistrict(ctx *models.Context, district *models.District) error {

	selector := bson.M{"_id": district.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": district, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDistrict : ""
func (d *Daos) FilterDistrict(ctx *models.Context, districtfilter *models.DistrictFilter, pagination *models.Pagination) ([]models.RefDistrict, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if districtfilter != nil {
		if len(districtfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": districtfilter.ActiveStatus}})
		}
		if len(districtfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": districtfilter.State}})
		}
		if len(districtfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": districtfilter.Status}})
		}
		//Regex
		if districtfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: districtfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSOILTYPE, "soilTypes", "_id", "ref.soilTypes", "ref.soilTypes")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONAGROECOLOGICALZONE, "agroEcologicalZones", "_id", "ref.agroEcologicalZones", "ref.agroEcologicalZones")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("district query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISTRICT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var districts []models.RefDistrict
	if err = cursor.All(context.TODO(), &districts); err != nil {
		return nil, err
	}
	return districts, nil
}

//EnableDistrict :""
func (d *Daos) EnableDistrict(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDistrict :""
func (d *Daos) DisableDistrict(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDistrict :""
func (d *Daos) DeleteDistrict(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISTRICTSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISTRICT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
