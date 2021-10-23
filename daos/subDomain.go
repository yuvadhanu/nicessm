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

//SaveSubDomain :""
func (d *Daos) SaveSubDomain(ctx *models.Context, SubDomain *models.SubDomain) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSUBDOMAIN).InsertOne(ctx.CTX, SubDomain)
	if err != nil {
		return err
	}
	SubDomain.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleSubDomain : ""
func (d *Daos) GetSingleSubDomain(ctx *models.Context, code string) (*models.RefSubDomain, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSUBDOMAIN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var SubDomains []models.RefSubDomain
	var SubDomain *models.RefSubDomain
	if err = cursor.All(ctx.CTX, &SubDomains); err != nil {
		return nil, err
	}
	if len(SubDomains) > 0 {
		SubDomain = &SubDomains[0]
	}
	return SubDomain, nil
}

//UpdateSubDomain : ""
func (d *Daos) UpdateSubDomain(ctx *models.Context, SubDomain *models.SubDomain) error {
	selector := bson.M{"_id": SubDomain.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": SubDomain, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSUBDOMAIN).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterSubDomain : ""
func (d *Daos) FilterSubDomain(ctx *models.Context, SubDomainfilter *models.SubDomainFilter, pagination *models.Pagination) ([]models.RefSubDomain, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if SubDomainfilter != nil {

		if len(SubDomainfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": SubDomainfilter.ActiveStatus}})
		}
		if len(SubDomainfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": SubDomainfilter.Status}})
		}
		if len(SubDomainfilter.KnowledgeDomain) > 0 {
			query = append(query, bson.M{"knowledgeDomain": bson.M{"$in": SubDomainfilter.KnowledgeDomain}})
		}
		//Regex
		if SubDomainfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: SubDomainfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSUBDOMAIN).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("SubDomain query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSUBDOMAIN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var SubDomains []models.RefSubDomain
	if err = cursor.All(context.TODO(), &SubDomains); err != nil {
		return nil, err
	}
	return SubDomains, nil
}

//EnableSubDomain :""
func (d *Daos) EnableSubDomain(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SUBDOMAINSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSUBDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableSubDomain :""
func (d *Daos) DisableSubDomain(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SUBDOMAINSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSUBDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteSubDomain :""
func (d *Daos) DeleteSubDomain(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SUBDOMAINSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSUBDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
