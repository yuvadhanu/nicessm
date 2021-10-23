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

//SaveBlockCrop :""
func (d *Daos) SaveBlockCrop(ctx *models.Context, BlockCrop *models.BlockCrop) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONBLOCKCROP).InsertOne(ctx.CTX, BlockCrop)
	if err != nil {
		return err
	}
	BlockCrop.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleBlockCrop : ""
func (d *Daos) GetSingleBlockCrop(ctx *models.Context, code string) (*models.RefBlockCrop, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)
	// mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONAGROECOLOGICALZONE, "agroEcologicalZones", "_id", "ref.agroEcologicalZones", "ref.agroEcologicalZones")...)
	// //Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBLOCKCROP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var BlockCrops []models.RefBlockCrop
	var BlockCrop *models.RefBlockCrop
	if err = cursor.All(ctx.CTX, &BlockCrops); err != nil {
		return nil, err
	}
	if len(BlockCrops) > 0 {
		BlockCrop = &BlockCrops[0]
	}
	return BlockCrop, nil
}

//UpdateBlockCrop : ""
func (d *Daos) UpdateBlockCrop(ctx *models.Context, BlockCrop *models.BlockCrop) error {

	selector := bson.M{"_id": BlockCrop.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": BlockCrop, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONBLOCKCROP).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterBlockCrop : ""
func (d *Daos) FilterBlockCrop(ctx *models.Context, BlockCropfilter *models.BlockCropFilter, pagination *models.Pagination) ([]models.RefBlockCrop, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if BlockCropfilter != nil {
		if len(BlockCropfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": BlockCropfilter.ActiveStatus}})
		}
		if len(BlockCropfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": BlockCropfilter.State}})
		}
		if len(BlockCropfilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": BlockCropfilter.District}})
		}
		if len(BlockCropfilter.Block) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": BlockCropfilter.Block}})
		}
		if len(BlockCropfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": BlockCropfilter.Status}})
		}
		//Regex
		// if BlockCropfilter.Regex.Name != "" {
		// 	query = append(query, bson.M{"name": primitive.Regex{Pattern: BlockCropfilter.Regex.Name, Options: "xi"}})
		// }
	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCOMMODITY, "commodity", "_id", "ref.commodity", "ref.commodity")...)
	if BlockCropfilter.Regex.Name != "" {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"ref.commodity.commonName": primitive.Regex{Pattern: BlockCropfilter.Regex.Name, Options: "xi"}}})

	}

	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONBLOCKCROP).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	//Aggregation
	d.Shared.BsonToJSONPrintTag("BlockCrop query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONBLOCKCROP).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var BlockCrops []models.RefBlockCrop
	if err = cursor.All(context.TODO(), &BlockCrops); err != nil {
		return nil, err
	}
	return BlockCrops, nil
}

//EnableBlockCrop :""
func (d *Daos) EnableBlockCrop(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.BLOCKCROPSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBLOCKCROP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableBlockCrop :""
func (d *Daos) DisableBlockCrop(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.BLOCKCROPSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBLOCKCROP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteBlockCrop :""
func (d *Daos) DeleteBlockCrop(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.BLOCKCROPSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONBLOCKCROP).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
