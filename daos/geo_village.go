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

//SaveVillage :""
func (d *Daos) SaveVillage(ctx *models.Context, village *models.Village) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONVILLAGE).InsertOne(ctx.CTX, village)
	return err
}

//GetSingleVillage : ""
func (d *Daos) GetSingleVillage(ctx *models.Context, ID string) (*models.RefVillage, error) {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "ref.gramPanchayat.block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.block.district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVILLAGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var villages []models.RefVillage
	var village *models.RefVillage
	if err = cursor.All(ctx.CTX, &villages); err != nil {
		return nil, err
	}
	if len(villages) > 0 {
		village = &villages[0]
	}
	return village, nil
}

//UpdateVillage : ""
func (d *Daos) UpdateVillage(ctx *models.Context, village *models.Village) error {

	selector := bson.M{"_id": village.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": village, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONVILLAGE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterVillage : ""
func (d *Daos) FilterVillage(ctx *models.Context, villagefilter *models.VillageFilter, pagination *models.Pagination) ([]models.RefVillage, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if villagefilter != nil {
		if len(villagefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"code": bson.M{"$in": villagefilter.ActiveStatus}})
		}
		if len(villagefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": villagefilter.Status}})
		}
		if len(villagefilter.GramPanchayat) > 0 {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$in": villagefilter.GramPanchayat}})
		}
		if villagefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: villagefilter.Regex.Name, Options: "ix"}})
		}
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	// if villagefilter != nil {
	// 	if villagefilter.SortBy != "" {
	// 		mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{villagefilter.SortBy: villagefilter.SortOrder}})

	// 	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONVILLAGE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONGRAMPANCHAYAT, "gramPanchayat", "_id", "ref.gramPanchayat", "ref.gramPanchayat")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "ref.gramPanchayat.block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.block.district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.state", "_id", "ref.state", "ref.state")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("village query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONVILLAGE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var villages []models.RefVillage
	if err = cursor.All(context.TODO(), &villages); err != nil {
		return nil, err
	}
	return villages, nil
}

//EnableVillage :""
func (d *Daos) EnableVillage(ctx *models.Context, ID string) error {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.VILLAGESTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONVILLAGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableVillage :""
func (d *Daos) DisableVillage(ctx *models.Context, ID string) error {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.VILLAGESTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONVILLAGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteVillage :""
func (d *Daos) DeleteVillage(ctx *models.Context, ID string) error {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.VILLAGESTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONVILLAGE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
