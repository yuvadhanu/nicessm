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

//SaveTopic :""
func (d *Daos) SaveTopic(ctx *models.Context, Topic *models.Topic) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONTOPIC).InsertOne(ctx.CTX, Topic)
	if err != nil {
		return err
	}
	Topic.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleTopic : ""
func (d *Daos) GetSingleTopic(ctx *models.Context, code string) (*models.RefTopic, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTOPIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Topics []models.RefTopic
	var Topic *models.RefTopic
	if err = cursor.All(ctx.CTX, &Topics); err != nil {
		return nil, err
	}
	if len(Topics) > 0 {
		Topic = &Topics[0]
	}
	return Topic, nil
}

//UpdateTopic : ""
func (d *Daos) UpdateTopic(ctx *models.Context, Topic *models.Topic) error {
	selector := bson.M{"_id": Topic.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Topic, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONTOPIC).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterTopic : ""
func (d *Daos) FilterTopic(ctx *models.Context, Topicfilter *models.TopicFilter, pagination *models.Pagination) ([]models.RefTopic, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Topicfilter != nil {

		if len(Topicfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Topicfilter.ActiveStatus}})
		}
		if len(Topicfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Topicfilter.Status}})
		}
		if len(Topicfilter.SubDomain) > 0 {
			query = append(query, bson.M{"subDomain": bson.M{"$in": Topicfilter.SubDomain}})
		}
		//Regex
		if Topicfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Topicfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONTOPIC).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Topic query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONTOPIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Topics []models.RefTopic
	if err = cursor.All(context.TODO(), &Topics); err != nil {
		return nil, err
	}
	return Topics, nil
}

//EnableTopic :""
func (d *Daos) EnableTopic(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.TOPICSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONTOPIC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableTopic :""
func (d *Daos) DisableTopic(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.TOPICSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONTOPIC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteTopic :""
func (d *Daos) DeleteTopic(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.TOPICSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONTOPIC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
