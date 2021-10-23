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

//SaveSubTopic :""
func (d *Daos) SaveSubTopic(ctx *models.Context, SubTopic *models.SubTopic) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONSUBTOPIC).InsertOne(ctx.CTX, SubTopic)
	if err != nil {
		return err
	}
	SubTopic.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//GetSingleSubTopic : ""
func (d *Daos) GetSingleSubTopic(ctx *models.Context, code string) (*models.RefSubTopic, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Lookups
	//Lookups
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTOPIC, "topic", "_id", "ref.topic", "ref.topic")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSUBTOPIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var SubTopics []models.RefSubTopic
	var SubTopic *models.RefSubTopic
	if err = cursor.All(ctx.CTX, &SubTopics); err != nil {
		return nil, err
	}
	if len(SubTopics) > 0 {
		SubTopic = &SubTopics[0]
	}
	return SubTopic, nil
}

//UpdateSubTopic : ""
func (d *Daos) UpdateSubTopic(ctx *models.Context, SubTopic *models.SubTopic) error {
	selector := bson.M{"_id": SubTopic.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": SubTopic, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONSUBTOPIC).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterSubTopic : ""
func (d *Daos) FilterSubTopic(ctx *models.Context, SubTopicfilter *models.SubTopicFilter, pagination *models.Pagination) ([]models.RefSubTopic, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if SubTopicfilter != nil {

		if len(SubTopicfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": SubTopicfilter.ActiveStatus}})
		}
		if len(SubTopicfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": SubTopicfilter.Status}})
		}
		if len(SubTopicfilter.Topic) > 0 {
			query = append(query, bson.M{"topic": bson.M{"$in": SubTopicfilter.Topic}})
		}
		//Regex
		if SubTopicfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: SubTopicfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONSUBTOPIC).CountDocuments(ctx.CTX, func() bson.M {
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
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTOPIC, "topic", "_id", "ref.topic", "ref.topic")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("SubTopic query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONSUBTOPIC).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var SubTopics []models.RefSubTopic
	if err = cursor.All(context.TODO(), &SubTopics); err != nil {
		return nil, err
	}
	return SubTopics, nil
}

//EnableSubTopic :""
func (d *Daos) EnableSubTopic(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SUBTOPICSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSUBTOPIC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableSubTopic :""
func (d *Daos) DisableSubTopic(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SUBTOPICSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSUBTOPIC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteSubTopic :""
func (d *Daos) DeleteSubTopic(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.SUBTOPICSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONSUBTOPIC).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
