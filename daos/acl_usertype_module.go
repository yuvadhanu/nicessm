package daos

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveACLUserTypeModuleMultiple :""
func (d *Daos) SaveACLUserTypeModuleMultiple(ctx *models.Context, modules []models.ACLUserTypeModule) error {
	for _, v := range modules {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"moduleId": v.ModuleID, "userTypeId": v.UserTypeID}
		updateData := bson.M{"$set": v}
		if _, err := ctx.DB.Collection(constants.COLLECTIONACLUSERTYPEMODULE).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

//FilterACLUserTypeModule : ""
func (d *Daos) FilterACLUserTypeModule(ctx *models.Context, filter *models.ACLUserTypeModuleFilter, pagination *models.Pagination) ([]models.RefACLUserTypeModule, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {

		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}

		if len(filter.Module) > 0 {
			query = append(query, bson.M{"moduleId": bson.M{"$in": filter.Module}})
		}
		if len(filter.UserType) > 0 {
			query = append(query, bson.M{"userTypeId": bson.M{"$in": filter.UserType}})
		}
		if len(filter.Check) > 0 {
			query = append(query, bson.M{"check": bson.M{"$in": filter.Check}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONACLUSERTYPEMODULE).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONMODULE, "moduleId", "uniqueId", "ref.module", "ref.module")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("ACLUserTypeModule query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONACLUSERTYPEMODULE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var data []models.RefACLUserTypeModule
	if err = cursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return data, nil
}
