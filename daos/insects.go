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

func (d *Daos) SaveInsect(ctx *models.Context, insect *models.Insect) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONINSECT).InsertOne(ctx.CTX, insect)
	if err != nil {
		return err
	}
	insect.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleInsect(ctx *models.Context, UniqueID string) (*models.RefInsect, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONINSECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var insects []models.RefInsect
	var insect *models.RefInsect
	if err = cursor.All(ctx.CTX, &insects); err != nil {
		return nil, err
	}
	if len(insects) > 0 {
		insect = &insects[0]
	}
	return insect, nil
}

func (d *Daos) UpdateInsect(ctx *models.Context, insect *models.Insect) error {

	selector := bson.M{"_id": insect.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": insect}
	_, err := ctx.DB.Collection(constants.COLLECTIONINSECT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterInsect(ctx *models.Context, insectfilter *models.InsectFilter, pagination *models.Pagination) ([]models.RefInsect, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if insectfilter != nil {

		if len(insectfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activestatus": bson.M{"$in": insectfilter.ActiveStatus}})
		}
		if len(insectfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": insectfilter.Status}})
		}
		if len(insectfilter.OmitIDs) > 0 {
			query = append(query, bson.M{"_id": bson.M{"$nin": insectfilter.OmitIDs}})
		}
		//Regex
		if insectfilter.Searchbox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: insectfilter.Searchbox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if insectfilter != nil {
		if insectfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{insectfilter.SortBy: insectfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONINSECT).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONINSECT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var insect []models.RefInsect
	if err = cursor.All(context.TODO(), &insect); err != nil {
		return nil, err
	}
	return insect, nil
}

func (d *Daos) EnableInsect(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.INSECTSTATUSTRUE, "status": constants.INSECTSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONINSECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableInsect(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.INSECTSTATUSFALSE, "status": constants.INSECTSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONINSECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteInsect(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.INSECTSTATUSFALSE, "status": constants.INSECTSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONINSECT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
