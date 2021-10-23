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

//SaveLanguage :""
func (d *Daos) SaveLanguage(ctx *models.Context, language *models.Languages) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGE).InsertOne(ctx.CTX, language)
	if err != nil {
		return err
	}
	language.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleLanguage : ""
func (d *Daos) GetSingleLanguage(ctx *models.Context, UniqueID string) (*models.RefLanguage, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var languages []models.RefLanguage
	var language *models.RefLanguage
	if err = cursor.All(ctx.CTX, &languages); err != nil {
		return nil, err
	}
	if len(languages) > 0 {
		language = &languages[0]
	}
	return language, nil
}

//UpdateLanguage : ""
func (d *Daos) UpdateLanguage(ctx *models.Context, language *models.Languages) error {

	selector := bson.M{"_id": language.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": language}
	_, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterLanguage : ""
func (d *Daos) FilterLanguage(ctx *models.Context, languagefilter *models.LanguageFilter, pagination *models.Pagination) ([]models.RefLanguage, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if languagefilter != nil {

		if len(languagefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activestatus": bson.M{"$in": languagefilter.ActiveStatus}})
		}
		if len(languagefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": languagefilter.Status}})
		}
		//Regex
		if languagefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: languagefilter.Regex.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if languagefilter != nil {
		if languagefilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{languagefilter.SortBy: languagefilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGE).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONLANGAUAGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var languages []models.RefLanguage
	if err = cursor.All(context.TODO(), &languages); err != nil {
		return nil, err
	}
	return languages, nil
}

//EnableLanguage :""
func (d *Daos) EnableLanguage(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.LANGAUAGESTATUSTRUE, "status": constants.LANGAUAGESTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANGAUAGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableLanguage :""
func (d *Daos) DisableLanguage(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.LANGAUAGESTATUSFALSE, "status": constants.LANGAUAGESTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANGAUAGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteLanguage :""
func (d *Daos) DeleteLanguage(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"activestatus": constants.LANGAUAGESTATUSFALSE, "status": constants.LANGAUAGESTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONLANGAUAGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
