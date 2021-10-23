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
func (d *Daos) SaveVaccine(ctx *models.Context, vaccine *models.Vaccine) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONVACCINE).InsertOne(ctx.CTX, vaccine)
	if err != nil {
		return err
	}
	vaccine.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleVaccine(ctx *models.Context, UniqueID string) (*models.RefVaccine, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVACCINE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vaccines []models.RefVaccine
	var vaccine *models.RefVaccine
	if err = cursor.All(ctx.CTX, &vaccines); err != nil {
		return nil, err
	}
	if len(vaccines) > 0 {
		vaccine = &vaccines[0]
	}
	return vaccine, nil
}

//UpdateVaccine : ""
func (d *Daos) UpdateVaccine(ctx *models.Context, vaccine *models.Vaccine) error {

	selector := bson.M{"_id": vaccine.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": vaccine}
	_, err := ctx.DB.Collection(constants.COLLECTIONVACCINE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterVaccine : ""
func (d *Daos) FilterVaccine(ctx *models.Context, vaccinefilter *models.VaccineFilter, pagination *models.Pagination) ([]models.RefVaccine, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if vaccinefilter != nil {

		if len(vaccinefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": vaccinefilter.ActiveStatus}})
		}
		if len(vaccinefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": vaccinefilter.Status}})
		}
		//Regex
		if vaccinefilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: vaccinefilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if vaccinefilter != nil {
		if vaccinefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{vaccinefilter.SortBy: vaccinefilter.SortOrder}})

		}

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVACCINE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("vaccine query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVACCINE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var vaccines []models.RefVaccine
	if err = cursor.All(context.TODO(), &vaccines); err != nil {
		return nil, err
	}
	return vaccines, nil
}

//EnableVaccine :""
func (d *Daos) EnableVaccine(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.VACCINESTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONVACCINE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableVaccine(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.VACCINESTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONVACCINE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteVaccine :""
func (d *Daos) DeleteVaccine(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.VACCINESTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONVACCINE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
