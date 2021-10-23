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

func (d *Daos) SaveProductConfig(ctx *models.Context, productConfig *models.ProductConfig) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).InsertOne(ctx.CTX, productConfig)
	if err != nil {
		return err
	}
	productConfig.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleProductConfig(ctx *models.Context, UniqueID string) (*models.RefProductConfig, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var productConfigs []models.RefProductConfig
	var productConfig *models.RefProductConfig
	if err = cursor.All(ctx.CTX, &productConfigs); err != nil {
		return nil, err
	}
	if len(productConfigs) > 0 {
		productConfig = &productConfigs[0]
	}
	return productConfig, nil
}

func (d *Daos) UpdateProductConfig(ctx *models.Context, productConfig *models.ProductConfig) error {

	selector := bson.M{"_id": productConfig.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": productConfig}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterProductConfig(ctx *models.Context, productConfigfilter *models.ProductConfigFilter, pagination *models.Pagination) ([]models.RefProductConfig, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if productConfigfilter != nil {

		if len(productConfigfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": productConfigfilter.Status}})
		}
		//Regex
		if productConfigfilter.Searchbox.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: productConfigfilter.Searchbox.Name, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if productConfigfilter != nil {
		if productConfigfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{productConfigfilter.SortBy: productConfigfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).CountDocuments(ctx.CTX, func() bson.M {
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
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var productConfig []models.RefProductConfig
	if err = cursor.All(context.TODO(), &productConfig); err != nil {
		return nil, err
	}
	return productConfig, nil
}

func (d *Daos) EnableProductConfig(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTCONFIGSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableProductConfig(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTCONFIGSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteProductConfig(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PRODUCTCONFIGSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) SetdefaultProductConfig(ctx *models.Context, UniqueID string) error {
	filter := bson.M{
		"isdefault": bson.M{
			"$eq": true,
		},
	}
	updatemany := bson.M{"$set": bson.M{"isdefault": false}}
	_, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).UpdateMany(ctx.CTX, filter, updatemany)
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isdefault": constants.PRODUCTCONFIGSTATUSTRUE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
func (d *Daos) GetactiveProductConfig(ctx *models.Context, IsDefault bool) (*models.ProductConfig, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"isdefault": IsDefault}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPRODUCTCONFIG, "categoryId", "uniqueId", "ref.category", "ref.category")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPRODUCTCONFIG).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var products []models.ProductConfig
	var product *models.ProductConfig
	if err = cursor.All(ctx.CTX, &products); err != nil {
		return nil, err
	}
	if len(products) > 0 {
		product = &products[0]
	}
	return product, nil
}
