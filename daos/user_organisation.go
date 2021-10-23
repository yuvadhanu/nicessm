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

//SaveOrganisation :""
func (d *Daos) SaveUserOrganisation(ctx *models.Context, organisation *models.UserOrganisation) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).InsertOne(ctx.CTX, organisation)
	return err
}

//GetSingleOrganisation : ""
func (d *Daos) GetSingleUserOrganisation(ctx *models.Context, UniqueID string) (*models.RefUserOrganisation, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": UniqueID}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var organisations []models.RefUserOrganisation
	var organisation *models.RefUserOrganisation
	if err = cursor.All(ctx.CTX, &organisations); err != nil {
		return nil, err
	}
	if len(organisations) > 0 {
		organisation = &organisations[0]
	}
	return organisation, nil
}

//UpdateOrganisation : ""
func (d *Daos) UpdateUserOrganisation(ctx *models.Context, organisation *models.UserOrganisation) error {
	selector := bson.M{"uniqueId": organisation.UniqueID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": organisation}
	_, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterOrganisation : ""
func (d *Daos) FilterUserOrganisation(ctx *models.Context, organisationfilter *models.UserOrganisationFilter, pagination *models.Pagination) ([]models.RefUserOrganisation, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if organisationfilter != nil {

		if len(organisationfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": organisationfilter.Status}})
		}
		//Regex
		if organisationfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: organisationfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("organisation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var organisations []models.RefUserOrganisation
	if err = cursor.All(context.TODO(), &organisations); err != nil {
		return nil, err
	}
	return organisations, nil
}

//EnableOrganisation :""
func (d *Daos) EnableUserOrganisation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ORGANISATIONOWNERSTATUSACTIVE}}
	_, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableOrganisation :""
func (d *Daos) DisableUserOrganisation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ORGANISATIONOWNERSTATUSDISABLED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteOrganisation :""
func (d *Daos) DeleteUserOrganisation(ctx *models.Context, UniqueID string) error {
	query := bson.M{"uniqueId": UniqueID}
	update := bson.M{"$set": bson.M{"status": constants.ORGANISATIONOWNERSTATUSDELETED}}
	_, err := ctx.DB.Collection(constants.COLLECTIONORGANISATION).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
