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

//SaveLiveStockVaccination :""
func (d *Daos) SaveLiveStockVaccination(ctx *models.Context, LiveStockVaccination *models.LiveStockVaccination) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONLIVESTOCKVACCINATION).InsertOne(ctx.CTX, LiveStockVaccination)
	if err != nil {
		return err
	}
	LiveStockVaccination.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleLiveStockVaccination : ""
func (d *Daos) GetSingleLiveStockVaccination(ctx *models.Context, code string) (*models.RefLiveStockVaccination, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.liveStock", "ref.liveStock")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDISEASE, "disease", "_id", "ref.disease", "ref.disease")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVACCINE, "vaccine", "_id", "ref.vaccine", "ref.vaccine")...)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLIVESTOCKVACCINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var LiveStockVaccinations []models.RefLiveStockVaccination
	var LiveStockVaccination *models.RefLiveStockVaccination
	if err = cursor.All(ctx.CTX, &LiveStockVaccinations); err != nil {
		return nil, err
	}
	if len(LiveStockVaccinations) > 0 {
		LiveStockVaccination = &LiveStockVaccinations[0]
	}
	return LiveStockVaccination, nil
}

//UpdateLiveStockVaccination : ""
func (d *Daos) UpdateLiveStockVaccination(ctx *models.Context, LiveStockVaccination *models.LiveStockVaccination) error {

	selector := bson.M{"_id": LiveStockVaccination.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": LiveStockVaccination, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONLIVESTOCKVACCINATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterLiveStockVaccination : ""
func (d *Daos) FilterLiveStockVaccination(ctx *models.Context, LiveStockVaccinationfilter *models.LiveStockVaccinationFilter, pagination *models.Pagination) ([]models.RefLiveStockVaccination, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if LiveStockVaccinationfilter != nil {
		if len(LiveStockVaccinationfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": LiveStockVaccinationfilter.ActiveStatus}})
		}
		if len(LiveStockVaccinationfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": LiveStockVaccinationfilter.State}})
		}
		if len(LiveStockVaccinationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": LiveStockVaccinationfilter.Status}})
		}
		if len(LiveStockVaccinationfilter.Diseases) > 0 {
			query = append(query, bson.M{"diseases": bson.M{"$in": LiveStockVaccinationfilter.Diseases}})
		}
		if len(LiveStockVaccinationfilter.LiveStocks) > 0 {
			query = append(query, bson.M{"liveStocks": bson.M{"$in": LiveStockVaccinationfilter.LiveStocks}})
		}
		if len(LiveStockVaccinationfilter.Booster) > 0 {
			query = append(query, bson.M{"booster": bson.M{"$in": LiveStockVaccinationfilter.Booster}})
		}

		if len(LiveStockVaccinationfilter.TimeOfVaccination) > 0 {
			query = append(query, bson.M{"timeOfVaccination": bson.M{"$in": LiveStockVaccinationfilter.TimeOfVaccination}})
		}
		//Regex
		if LiveStockVaccinationfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: LiveStockVaccinationfilter.Regex.Name, Options: "xi"}})
		}
		if LiveStockVaccinationfilter.Regex.BoosterDose != "" {
			query = append(query, bson.M{"boosterDose": primitive.Regex{Pattern: LiveStockVaccinationfilter.Regex.BoosterDose, Options: "xi"}})
		}
		if LiveStockVaccinationfilter.Regex.BoosterTime != "" {
			query = append(query, bson.M{"boosterTime": primitive.Regex{Pattern: LiveStockVaccinationfilter.Regex.BoosterTime, Options: "xi"}})
		}
		if LiveStockVaccinationfilter.Regex.Dose != "" {
			query = append(query, bson.M{"dose": primitive.Regex{Pattern: LiveStockVaccinationfilter.Regex.Dose, Options: "xi"}})
		}
		if LiveStockVaccinationfilter.Regex.Immunity != "" {
			query = append(query, bson.M{"immunity": primitive.Regex{Pattern: LiveStockVaccinationfilter.Regex.Immunity, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLIVESTOCKVACCINATION).CountDocuments(ctx.CTX, func() bson.M {
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
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.liveStock", "ref.liveStock")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDISEASE, "diseases", "_id", "ref.diseases", "ref.diseases")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONVACCINE, "vaccine", "_id", "ref.vaccine", "ref.vaccine")...)
	// //Aggregation
	d.Shared.BsonToJSONPrintTag("LiveStockVaccination query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLIVESTOCKVACCINATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var LiveStockVaccinations []models.RefLiveStockVaccination
	if err = cursor.All(context.TODO(), &LiveStockVaccinations); err != nil {
		return nil, err
	}
	return LiveStockVaccinations, nil
}

//EnableLiveStockVaccination :""
func (d *Daos) EnableLiveStockVaccination(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.LIVESTOCKVACCINATIONSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLIVESTOCKVACCINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableLiveStockVaccination :""
func (d *Daos) DisableLiveStockVaccination(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.LIVESTOCKVACCINATIONSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLIVESTOCKVACCINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteLiveStockVaccination :""
func (d *Daos) DeleteLiveStockVaccination(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.LIVESTOCKVACCINATIONSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLIVESTOCKVACCINATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
