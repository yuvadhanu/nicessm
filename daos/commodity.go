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

//SaveCommodity :""
func (d *Daos) SaveCommodity(ctx *models.Context, commodity *models.Commodity) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITY).InsertOne(ctx.CTX, commodity)
	return err
}

//UpdateCommodity : ""
func (d *Daos) UpdateCommodity(ctx *models.Context, commodity *models.Commodity) error {

	selector := bson.M{"_id": commodity.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": commodity}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableCommodity :""
func (d *Daos) EnableCommodity(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCommodity :""
func (d *Daos) DisableCommodity(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCommodity :""
func (d *Daos) DeleteCommodity(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.COMMODITYSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCOMMODITY).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleCommodity : ""
func (d *Daos) GetSingleCommodity(ctx *models.Context, UniqueID string) (*models.RefCommodity, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYCATEGORY, "category", "_id", "ref.category", "ref.category")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYFUNCTION, "function", "_id", "ref.function", "ref.function")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDISEASE, "diseases", "_id", "ref.diseases", "ref.diseases")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONINSECT, "insects", "_id", "ref.insects", "ref.insects")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Commoditys []models.RefCommodity
	var Commodity *models.RefCommodity
	if err = cursor.All(ctx.CTX, &Commoditys); err != nil {
		return nil, err
	}
	if len(Commoditys) > 0 {
		Commodity = &Commoditys[0]
	}
	return Commodity, nil
}

//FilterCommodity : ""
func (d *Daos) FilterCommodity(ctx *models.Context, filter *models.CommodityFilter, pagination *models.Pagination) ([]models.RefCommodity, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.Classfication) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Classfication}})
		}
		if len(filter.Category) > 0 {
			query = append(query, bson.M{"category": bson.M{"$in": filter.Category}})
		}
		if len(filter.Function) > 0 {
			query = append(query, bson.M{"function": bson.M{"$in": filter.Function}})
		}
		if filter.SearchBox.CommonName != "" {
			query = append(query, bson.M{"commonName": primitive.Regex{Pattern: filter.SearchBox.CommonName, Options: "xi"}})
		}
		if filter.SearchBox.ScientificName != "" {
			query = append(query, bson.M{"scientificName": primitive.Regex{Pattern: filter.SearchBox.ScientificName, Options: "xi"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYCATEGORY, "category", "_id", "ref.category", "ref.category")...)
	if filter != nil {
		if len(filter.Classfication) > 0 {
			mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"ref.category.classification": bson.M{"$in": filter.Classfication}}})
		}
	}
	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITY).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITYFUNCTION, "function", "_id", "ref.function", "ref.function")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONDISEASE, "diseases", "_id", "ref.diseases", "ref.diseases")...)
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONINSECT, "insects", "_id", "ref.insects", "ref.insects")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Commodity query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITY).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Commoditys []models.RefCommodity
	if err = cursor.All(context.TODO(), &Commoditys); err != nil {
		return nil, err
	}
	return Commoditys, nil
}

//AddInsectsCommodity : ""
func (d *Daos) AddInsectsCommodity(ctx *models.Context, commodity *models.Commodity) error {

	selector := bson.M{"_id": commodity.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$push": bson.M{"insects": bson.M{"$each": commodity.Insects}}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//AddDieseasesCommodity : ""
func (d *Daos) AddDieseasesCommodity(ctx *models.Context, commodity *models.Commodity) error {

	selector := bson.M{"_id": commodity.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$push": bson.M{"diseases": bson.M{"$each": commodity.Diseases}}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITY).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//DeleteInsectsCommodity : ""
func (d *Daos) DeleteInsectsCommodity(ctx *models.Context, commodity *models.Commodity) error {

	selector := bson.M{"_id": commodity.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$pull": bson.M{"insects": bson.M{"$in": commodity.Insects}}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITY).UpdateMany(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//DeleteDiseasesCommodity : ""
func (d *Daos) DeleteDiseasesCommodity(ctx *models.Context, commodity *models.Commodity) error {

	selector := bson.M{"_id": commodity.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$pull": bson.M{"diseases": bson.M{"$in": commodity.Diseases}}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCOMMODITY).UpdateMany(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}
