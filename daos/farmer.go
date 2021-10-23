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

//SaveFarmer :""
func (d *Daos) SaveFarmer(ctx *models.Context, Farmer *models.Farmer) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONFARMER).InsertOne(ctx.CTX, Farmer)
	if err != nil {
		return err
	}
	Farmer.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleFarmer : ""
func (d *Daos) GetSingleFarmer(ctx *models.Context, code string) (*models.RefFarmer, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONLANGAUAGE, "languages", "_id", "ref.languages", "ref.languages")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Farmers []models.RefFarmer
	var Farmer *models.RefFarmer
	if err = cursor.All(ctx.CTX, &Farmers); err != nil {
		return nil, err
	}
	if len(Farmers) > 0 {
		Farmer = &Farmers[0]
	}
	return Farmer, nil
}

//UpdateFarmer : ""
func (d *Daos) UpdateFarmer(ctx *models.Context, Farmer *models.Farmer) error {
	selector := bson.M{"_id": Farmer.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Farmer, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterFarmer : ""
func (d *Daos) FilterFarmer(ctx *models.Context, Farmerfilter *models.FarmerFilter, pagination *models.Pagination) ([]models.RefFarmer, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Farmerfilter != nil {

		if len(Farmerfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Farmerfilter.ActiveStatus}})
		}
		if len(Farmerfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Farmerfilter.Status}})
		}
		if len(Farmerfilter.State) > 0 {
			query = append(query, bson.M{"state": bson.M{"$in": Farmerfilter.State}})
		}
		if len(Farmerfilter.District) > 0 {
			query = append(query, bson.M{"district": bson.M{"$in": Farmerfilter.District}})
		}
		if len(Farmerfilter.Block) > 0 {
			query = append(query, bson.M{"block": bson.M{"$in": Farmerfilter.Block}})
		}
		if len(Farmerfilter.GramPanchayat) > 0 {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$in": Farmerfilter.GramPanchayat}})
		}
		if len(Farmerfilter.Village) > 0 {
			query = append(query, bson.M{"village": bson.M{"$in": Farmerfilter.Village}})
		}
		//Regex
		if Farmerfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Farmerfilter.Regex.Name, Options: "xi"}})
		}
		if Farmerfilter.Regex.MobileNumber != "" {
			query = append(query, bson.M{"mobileNumber": primitive.Regex{Pattern: Farmerfilter.Regex.MobileNumber, Options: "xi"}})
		}
		if Farmerfilter.Regex.SpouseName != "" {
			query = append(query, bson.M{"spouseName": primitive.Regex{Pattern: Farmerfilter.Regex.SpouseName, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONFARMER).CountDocuments(ctx.CTX, func() bson.M {
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
	//mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONLANGAUAGE, "languages", "_id", "ref.languages", "ref.languages")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Farmer query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONFARMER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Farmers []models.RefFarmer
	if err = cursor.All(context.TODO(), &Farmers); err != nil {
		return nil, err
	}
	return Farmers, nil
}

//EnableFarmer :""
func (d *Daos) EnableFarmer(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableFarmer :""
func (d *Daos) DisableFarmer(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteFarmer :""
func (d *Daos) DeleteFarmer(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.FARMERSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONFARMER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
