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

//SaveGramPanchayat :""
func (d *Daos) SaveGramPanchayat(ctx *models.Context, grampanchayat *models.GramPanchayat) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).InsertOne(ctx.CTX, grampanchayat)
	return err
}

//GetSingleGramPanchayat : ""
func (d *Daos) GetSingleGramPanchayat(ctx *models.Context, ID string) (*models.RefGramPanchayat, error) {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.block.district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var grampanchayats []models.RefGramPanchayat
	var grampanchayat *models.RefGramPanchayat
	if err = cursor.All(ctx.CTX, &grampanchayats); err != nil {
		return nil, err
	}
	if len(grampanchayats) > 0 {
		grampanchayat = &grampanchayats[0]
	}
	return grampanchayat, nil
}

//UpdateGramPanchayat : ""
func (d *Daos) UpdateGramPanchayat(ctx *models.Context, grampanchayat *models.GramPanchayat) error {
	selector := bson.M{"_id": grampanchayat.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": grampanchayat, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterGramPanchayat : ""
func (d *Daos) FilterGramPanchayat(ctx *models.Context, grampanchayatfilter *models.GramPanchayatFilter, pagination *models.Pagination) ([]models.RefGramPanchayat, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if grampanchayatfilter != nil {
		if len(grampanchayatfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": grampanchayatfilter.ActiveStatus}})
		}
		if len(grampanchayatfilter.Block) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": grampanchayatfilter.Block}})
		}
		if len(grampanchayatfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": grampanchayatfilter.Status}})
		}
		//Regex
		if grampanchayatfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: grampanchayatfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "ref.block.district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "ref.district.state", "_id", "ref.state", "ref.state")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("grampanchayat query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var grampanchayats []models.RefGramPanchayat
	if err = cursor.All(context.TODO(), &grampanchayats); err != nil {
		return nil, err
	}
	return grampanchayats, nil
}

//EnableGramPanchayat :""
func (d *Daos) EnableGramPanchayat(ctx *models.Context, ID string) error {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.GRAMPANCHAYATSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableGramPanchayat :""
func (d *Daos) DisableGramPanchayat(ctx *models.Context, ID string) error {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.GRAMPANCHAYATSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteGramPanchayat :""
func (d *Daos) DeleteGramPanchayat(ctx *models.Context, ID string) error {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.GRAMPANCHAYATSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONGRAMPANCHAYAT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
