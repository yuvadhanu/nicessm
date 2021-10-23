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

//SaveMenu :""
func (d *Daos) SaveMenu(ctx *models.Context, Menu *models.Menu) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONMENU).InsertOne(ctx.CTX, Menu)
	return err
}

//GetSingleMenu : ""
func (d *Daos) GetSingleMenu(ctx *models.Context, UniqueID string) (*models.RefMenu, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMENU).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var menus []models.RefMenu
	var Menu *models.RefMenu
	if err = cursor.All(ctx.CTX, &menus); err != nil {
		return nil, err
	}
	if len(menus) > 0 {
		Menu = &menus[0]
	}
	return Menu, nil
}

//UpdateMenu : ""
func (d *Daos) UpdateMenu(ctx *models.Context, Menu *models.Menu) error {
	selector := bson.M{"uniqueId": Menu.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Menu, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMENU).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterMenu : ""
func (d *Daos) FilterMenu(ctx *models.Context, menufilter *models.MenuFilter, pagination *models.Pagination) ([]models.RefMenu, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if menufilter != nil {

		if len(menufilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": menufilter.Status}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONMENU).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("Menu query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONMENU).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var menus []models.RefMenu
	if err = cursor.All(context.TODO(), &menus); err != nil {
		return nil, err
	}
	return menus, nil
}

//EnableMenu :""
func (d *Daos) EnableMenu(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERMENUSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMENU).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableMenu :""
func (d *Daos) DisableMenu(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERMENUSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMENU).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteMenu :""
func (d *Daos) DeleteMenu(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ACLMASTERMENUSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONMENU).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
