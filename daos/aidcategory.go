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

func (d *Daos) SaveAidCategory(ctx *models.Context, aidCategory *models.AidCategory) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONAIDCATEGORY).InsertOne(ctx.CTX, aidCategory)
	if err != nil {
		return err
	}
	aidCategory.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleAidCategory(ctx *models.Context, UniqueID string) (*models.RefAidCategory, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAIDCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var aidCategorys []models.RefAidCategory
	var aidCategory *models.RefAidCategory
	if err = cursor.All(ctx.CTX, &aidCategorys); err != nil {
		return nil, err
	}
	if len(aidCategorys) > 0 {
		aidCategory = &aidCategorys[0]
	}
	return aidCategory, nil
}

func (d *Daos) UpdateAidCategory(ctx *models.Context, aidCategory *models.AidCategory) error {

	selector := bson.M{"_id": aidCategory.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": aidCategory}
	_, err := ctx.DB.Collection(constants.COLLECTIONAIDCATEGORY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterAidCategory(ctx *models.Context, aidCategoryfilter *models.AidCategoryFilter, pagination *models.Pagination) ([]models.RefAidCategory, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if aidCategoryfilter != nil {

		if len(aidCategoryfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": aidCategoryfilter.Status}})
		}
		//SearchBox
		if aidCategoryfilter.SearchBox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: aidCategoryfilter.SearchBox.Name, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if aidCategoryfilter != nil {
		if aidCategoryfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{aidCategoryfilter.SortBy: aidCategoryfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONAIDCATEGORY).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("aidCategory query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONAIDCATEGORY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var aidCategorys []models.RefAidCategory
	if err = cursor.All(context.TODO(), &aidCategorys); err != nil {
		return nil, err
	}
	return aidCategorys, nil
}

func (d *Daos) EnableAidCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.AIDCATEGORYSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONAIDCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableAidCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.AIDCATEGORYSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONAIDCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteAidCategory(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.AIDCATEGORYSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONAIDCATEGORY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
