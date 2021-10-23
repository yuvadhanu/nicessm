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

//SaveSelfRegister :""
func (d *Daos) SaveSelfRegister(ctx *models.Context, selfregister *models.User) error {
	_, err := ctx.DB.Collection(constants.COLLECTIONSELFREGISTER).InsertOne(ctx.CTX, selfregister)
	return err
}

//GetSingleSelfRegister : ""
func (d *Daos) GetSingleSelfRegister(ctx *models.Context, UserName string) (*models.RefUser, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"userName": UserName}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSELFREGISTER, "managerId", "userName", "ref.manager", "ref.manager")...)

	d.Shared.BsonToJSONPrintTag("user query =>", mainPipeline)
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSELFREGISTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var selfregisters []models.RefUser
	var selfregister *models.RefUser
	if err = cursor.All(ctx.CTX, &selfregisters); err != nil {
		return nil, err
	}
	if len(selfregisters) > 0 {
		selfregister = &selfregisters[0]
	}
	return selfregister, nil
}

//UpdateSelfRegister : ""
func (d *Daos) UpdateSelfRegister(ctx *models.Context, selfregister *models.User) error {
	selector := bson.M{"userName": selfregister.UserName}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": selfregister}
	_, err := ctx.DB.Collection(constants.COLLECTIONSELFREGISTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterSelfRegister : ""
func (d *Daos) FilterSelfRegister(ctx *models.Context, selfregisterfilter *models.UserFilter, pagination *models.Pagination) ([]models.RefUser, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if selfregisterfilter != nil {
		if len(selfregisterfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": selfregisterfilter.Status}})
		}
		if len(selfregisterfilter.Manager) > 0 {
			query = append(query, bson.M{"managerId": bson.M{"$in": selfregisterfilter.Manager}})
		}
		if len(selfregisterfilter.Type) > 0 {
			query = append(query, bson.M{"type": bson.M{"$in": selfregisterfilter.Type}})
		}
		if len(selfregisterfilter.OmitID) > 0 {
			query = append(query, bson.M{"userName": bson.M{"$nin": selfregisterfilter.OmitID}})
		}
		if len(selfregisterfilter.OrganisationID) > 0 {
			query = append(query, bson.M{"organisationId": bson.M{"$in": selfregisterfilter.OrganisationID}})
		}

		//Regex
		if selfregisterfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: selfregisterfilter.Regex.Name, Options: "xi"}})
		}
		if selfregisterfilter.Regex.Contact != "" {
			query = append(query, bson.M{"mobile": primitive.Regex{Pattern: selfregisterfilter.Regex.Contact, Options: "xi"}})
		}
		if selfregisterfilter.Regex.UserName != "" {
			query = append(query, bson.M{"userName": primitive.Regex{Pattern: selfregisterfilter.Regex.UserName, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSELFREGISTER).CountDocuments(ctx.CTX, func() bson.M {
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

	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisationId", "uniqueId", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSELFREGISTER, "managerId", "userName", "ref.manager", "ref.manager")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("selfregister query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSELFREGISTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var selfregisters []models.RefUser
	if err = cursor.All(context.TODO(), &selfregisters); err != nil {
		return nil, err
	}
	return selfregisters, nil
}

func (d *Daos) ApprovedSelfRegister(ctx *models.Context, UserName string) error {
	query := bson.M{"userName": UserName}
	update := bson.M{"$set": bson.M{"status": "Approved"}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSELFREGISTER).UpdateOne(ctx.CTX, query, update)

	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) RejectSelfRegister(ctx *models.Context, UserName string) error {
	query := bson.M{"userName": UserName}
	update := bson.M{"$set": bson.M{"status": "Rejected"}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSELFREGISTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
