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

//SaveState :""
func (d *Daos) SaveState(ctx *models.Context, state *models.State) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSTATE).InsertOne(ctx.CTX, state)
	if err != nil {
		return err
	}
	state.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleState : ""
func (d *Daos) GetSingleState(ctx *models.Context, code string) (*models.RefState, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONLANGAUAGE, "languages", "_id", "ref.languages", "ref.languages")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.RefState
	var state *models.RefState
	if err = cursor.All(ctx.CTX, &states); err != nil {
		return nil, err
	}
	if len(states) > 0 {
		state = &states[0]
	}
	return state, nil
}

//UpdateState : ""
func (d *Daos) UpdateState(ctx *models.Context, state *models.State) error {
	selector := bson.M{"_id": state.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": state, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSTATE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterState : ""
func (d *Daos) FilterState(ctx *models.Context, statefilter *models.StateFilter, pagination *models.Pagination) ([]models.RefState, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if statefilter != nil {

		if len(statefilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": statefilter.ActiveStatus}})
		}
		if len(statefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": statefilter.Status}})
		}
		//Regex
		if statefilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: statefilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSTATE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONLANGAUAGE, "languages", "_id", "ref.languages", "ref.languages")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("state query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSTATE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var states []models.RefState
	if err = cursor.All(context.TODO(), &states); err != nil {
		return nil, err
	}
	return states, nil
}

//EnableState :""
func (d *Daos) EnableState(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATESTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableState :""
func (d *Daos) DisableState(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATESTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteState :""
func (d *Daos) DeleteState(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.STATESTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSTATE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
