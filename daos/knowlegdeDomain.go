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

//SaveKnowlegdeDomain :""
func (d *Daos) SaveKnowlegdeDomain(ctx *models.Context, KnowlegdeDomain *models.KnowledgeDomain) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).InsertOne(ctx.CTX, KnowlegdeDomain)
	if err != nil {
		return err
	}
	KnowlegdeDomain.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleKnowlegdeDomain : ""
func (d *Daos) GetSingleKnowlegdeDomain(ctx *models.Context, code string) (*models.RefKnowledgeDomain, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var KnowlegdeDomains []models.RefKnowledgeDomain
	var KnowlegdeDomain *models.RefKnowledgeDomain
	if err = cursor.All(ctx.CTX, &KnowlegdeDomains); err != nil {
		return nil, err
	}
	if len(KnowlegdeDomains) > 0 {
		KnowlegdeDomain = &KnowlegdeDomains[0]
	}
	return KnowlegdeDomain, nil
}

//UpdateKnowlegdeDomain : ""
func (d *Daos) UpdateKnowlegdeDomain(ctx *models.Context, KnowlegdeDomain *models.KnowledgeDomain) error {
	selector := bson.M{"_id": KnowlegdeDomain.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": KnowlegdeDomain, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterKnowlegdeDomain : ""
func (d *Daos) FilterKnowledgeDomain(ctx *models.Context, KnowledgeDomainfilter *models.KnowledgeDomainFilter, pagination *models.Pagination) ([]models.RefKnowledgeDomain, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if KnowledgeDomainfilter != nil {

		if len(KnowledgeDomainfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": KnowledgeDomainfilter.ActiveStatus}})
		}
		if len(KnowledgeDomainfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": KnowledgeDomainfilter.Status}})
		}
		//Regex
		if KnowledgeDomainfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: KnowledgeDomainfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).CountDocuments(ctx.CTX, func() bson.M {
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
	d.Shared.BsonToJSONPrintTag("KnowlegdeDomain query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var KnowlegdeDomains []models.RefKnowledgeDomain
	if err = cursor.All(context.TODO(), &KnowlegdeDomains); err != nil {
		return nil, err
	}
	return KnowlegdeDomains, nil
}

//EnableKnowlegdeDomain :""
func (d *Daos) EnableKnowlegdeDomain(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.KNOWLEDGEDOMAINSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableKnowlegdeDomain :""
func (d *Daos) DisableKnowlegdeDomain(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.KNOWLEDGEDOMAINSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteKnowlegdeDomain :""
func (d *Daos) DeleteKnowlegdeDomain(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.KNOWLEDGEDOMAINSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
