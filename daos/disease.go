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
func (d *Daos) SaveDisease(ctx *models.Context, disease *models.Disease) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONDISEASE).InsertOne(ctx.CTX, disease)
	if err != nil {
		return err
	}
	disease.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleDisease : ""
func (d *Daos) GetSingleDisease(ctx *models.Context, UniqueID string) (*models.RefDisease, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISEASE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var diseases []models.RefDisease
	var disease *models.RefDisease
	if err = cursor.All(ctx.CTX, &diseases); err != nil {
		return nil, err
	}
	if len(diseases) > 0 {
		disease = &diseases[0]
	}
	return disease, nil
}

//UpdateDisease : ""
func (d *Daos) UpdateDisease(ctx *models.Context, disease *models.Disease) error {

	selector := bson.M{"_id": disease.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": disease}
	_, err := ctx.DB.Collection(constants.COLLECTIONDISEASE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterDisease : ""
func (d *Daos) FilterDisease(ctx *models.Context, diseasefilter *models.DiseaseFilter, pagination *models.Pagination) ([]models.RefDisease, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if diseasefilter != nil {

		if len(diseasefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": diseasefilter.ActiveStatus}})
		}
		if len(diseasefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": diseasefilter.Status}})
		}
		if len(diseasefilter.OmitIDs) > 0 {
			query = append(query, bson.M{"_id": bson.M{"$nin": diseasefilter.OmitIDs}})
		}
		//Regex
		if diseasefilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: diseasefilter.SearchBox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if diseasefilter != nil {
		if diseasefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{diseasefilter.SortBy: diseasefilter.SortOrder}})

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
	d.Shared.BsonToJSONPrintTag("disease query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDISEASE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var diseases []models.RefDisease
	if err = cursor.All(context.TODO(), &diseases); err != nil {
		return nil, err
	}
	return diseases, nil
}

//EnableDisease :""
func (d *Daos) EnableDisease(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISEASESTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISEASE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableDisease :""
func (d *Daos) DisableDisease(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISEASESTATUSDISABLE, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISEASE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteDisease :""
func (d *Daos) DeleteDisease(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.DISEASESTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONDISEASE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
