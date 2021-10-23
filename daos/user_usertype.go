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
)

//SaveUserType :""
func (d *Daos) SaveUserType(ctx *models.Context, UserType *models.UserType) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).InsertOne(ctx.CTX, UserType)
	return err
}

//GetSingleUserType : ""
func (d *Daos) GetSingleUserType(ctx *models.Context, UniqueID string) (*models.RefUserType, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var userTypes []models.RefUserType
	var UserType *models.RefUserType
	if err = cursor.All(ctx.CTX, &userTypes); err != nil {
		return nil, err
	}
	if len(userTypes) > 0 {
		UserType = &userTypes[0]
	}
	return UserType, nil
}

//UpdateUserType : ""
func (d *Daos) UpdateUserType(ctx *models.Context, UserType *models.UserType) error {
	selector := bson.M{"uniqueId": UserType.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": UserType, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterUserType : ""
func (d *Daos) FilterUserType(ctx *models.Context, userTypefilter *models.UserTypeFilter, pagination *models.Pagination) ([]models.RefUserType, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if userTypefilter != nil {

		if len(userTypefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": userTypefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("UserType query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var userTypes []models.RefUserType
	if err = cursor.All(context.TODO(), &userTypes); err != nil {
		return nil, err
	}
	return userTypes, nil
}

//EnableUserType :""
func (d *Daos) EnableUserType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERTYPESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableUserType :""
func (d *Daos) DisableUserType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERTYPESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteUserType :""
func (d *Daos) DeleteUserType(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.USERTYPESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
