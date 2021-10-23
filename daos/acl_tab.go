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

//SaveTab :""
func (d *Daos) SaveTab(ctx *models.Context, Tab *models.Tab) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONTAB).InsertOne(ctx.CTX, Tab)
	return err
}

//GetSingleTab : ""
func (d *Daos) GetSingleTab(ctx *models.Context, UniqueID string) (*models.RefTab, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTAB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var tabs []models.RefTab
	var Tab *models.RefTab
	if err = cursor.All(ctx.CTX, &tabs); err != nil {
		return nil, err
	}
	if len(tabs) > 0 {
		Tab = &tabs[0]
	}
	return Tab, nil
}

//UpdateTab : ""
func (d *Daos) UpdateTab(ctx *models.Context, Tab *models.Tab) error {
	selector := bson.M{"uniqueId": Tab.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Tab, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTAB).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterTab : ""
func (d *Daos) FilterTab(ctx *models.Context, tabfilter *models.TabFilter, pagination *models.Pagination) ([]models.RefTab, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if tabfilter != nil {

		if len(tabfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": tabfilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTAB).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Tab query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTAB).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var tabs []models.RefTab
	if err = cursor.All(context.TODO(), &tabs); err != nil {
		return nil, err
	}
	return tabs, nil
}

//EnableTab :""
func (d *Daos) EnableTab(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERTABSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTAB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableTab :""
func (d *Daos) DisableTab(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERTABSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTAB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteTab :""
func (d *Daos) DeleteTab(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERTABSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTAB).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
