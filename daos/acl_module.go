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

//SaveModule :""
func (d *Daos) SaveModule(ctx *models.Context, Module *models.Module) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONMODULE).InsertOne(ctx.CTX, Module)
	return err
}

//GetSingleModule : ""
func (d *Daos) GetSingleModule(ctx *models.Context, UniqueID string) (*models.RefModule, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMODULE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var modules []models.RefModule
	var Module *models.RefModule
	if err = cursor.All(ctx.CTX, &modules); err != nil {
		return nil, err
	}
	if len(modules) > 0 {
		Module = &modules[0]
	}
	return Module, nil
}

//UpdateModule : ""
func (d *Daos) UpdateModule(ctx *models.Context, Module *models.Module) error {
	selector := bson.M{"uniqueId": Module.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Module, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMODULE).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterModule : ""
func (d *Daos) FilterModule(ctx *models.Context, modulefilter *models.ModuleFilter, pagination *models.Pagination) ([]models.RefModule, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if modulefilter != nil {

		if len(modulefilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": modulefilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMODULE).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Module query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMODULE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var modules []models.RefModule
	if err = cursor.All(context.TODO(), &modules); err != nil {
		return nil, err
	}
	return modules, nil
}

//EnableModule :""
func (d *Daos) EnableModule(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERMODULESTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMODULE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableModule :""
func (d *Daos) DisableModule(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERMODULESTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMODULE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteModule :""
func (d *Daos) DeleteModule(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERMODULESTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMODULE).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleModuleUserType : ""
func (d *Daos) GetSingleModuleUserType(ctx *models.Context, userTypeID string) (*models.UserTypeModuleAccess, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": userTypeID}})
	mainPipeline = append(mainPipeline, bson.M{"$lookup": bson.M{
		"from": constants.COLLECTIONMODULE,
		"as":   "modules",
		"let":  bson.M{"userTypeId": "$uniqueId"},
		"pipeline": []bson.M{
			bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
				bson.M{"$eq": []string{"$status", "Active"}},
			}}}},
			bson.M{"$lookup": bson.M{
				"from": constants.COLLECTIONACLUSERTYPEMODULE,
				"as":   "access",
				"let":  bson.M{"userTypeId": "$$userTypeId", "moduleId": "$uniqueId"},
				"pipeline": []bson.M{
					bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
						bson.M{"$eq": []string{"$userTypeId", "$$userTypeId"}},
						bson.M{"$eq": []string{"$moduleId", "$$moduleId"}},
					},
					}}},
				},
			},
			},
			bson.M{"$addFields": bson.M{"access": bson.M{"$arrayElemAt": []interface{}{"$access", 0}}}},
		},
	}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Module query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var modules []models.UserTypeModuleAccess
	var Module *models.UserTypeModuleAccess
	if err = cursor.All(ctx.CTX, &modules); err != nil {
		return nil, err
	}
	if len(modules) > 0 {
		Module = &modules[0]
	}
	return Module, nil
}
