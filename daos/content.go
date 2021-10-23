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

func (d *Daos) SaveContent(ctx *models.Context, content *models.Content) error {

	res, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).InsertOne(ctx.CTX, content)
	if err != nil {
		return err
	}
	content.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (d *Daos) GetSingleContent(ctx *models.Context, UniqueID string) (*models.RefContent, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisation", "_id", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "projects", "_id", "ref.projects", "ref.projects")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTOPIC, "topic", "_id", "ref.topic", "ref.topic")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBTOPIC, "subTopic", "_id", "ref.subTopic", "ref.subTopic")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCROPSEASON, "cropseason", "_id", "ref.cropseason", "ref.cropseason")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var contents []models.RefContent
	var content *models.RefContent
	if err = cursor.All(ctx.CTX, &contents); err != nil {
		return nil, err
	}
	if len(contents) > 0 {
		content = &contents[0]
	}
	return content, nil
}

func (d *Daos) UpdateContent(ctx *models.Context, content *models.Content) error {

	selector := bson.M{"_id": content.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": content}
	_, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

func (d *Daos) FilterContent(ctx *models.Context, contentfilter *models.ContentFilter, pagination *models.Pagination) ([]models.RefContent, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if contentfilter != nil {

		if len(contentfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": contentfilter.Status}})
		}
		//Regex
		if contentfilter.SearchBox.Content != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: contentfilter.SearchBox.Content, Options: "xi"}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if contentfilter != nil {
		if contentfilter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{contentfilter.SortBy: contentfilter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).CountDocuments(ctx.CTX, func() bson.M {
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
	//Lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONORGANISATION, "organisation", "_id", "ref.organisation", "ref.organisation")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "projects", "_id", "ref.projects", "ref.projects")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONTOPIC, "topic", "_id", "ref.topic", "ref.topic")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBDOMAIN, "subDomain", "_id", "ref.subDomain", "ref.subDomain")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSUBTOPIC, "subTopic", "_id", "ref.subTopic", "ref.subTopic")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONSTATE, "state", "_id", "ref.state", "ref.state")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONDISTRICT, "district", "_id", "ref.district", "ref.district")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONBLOCK, "block", "_id", "ref.block", "ref.block")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONCROPSEASON, "cropseason", "_id", "ref.season", "ref.season")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Aidlocation query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCONTENT).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Contents []models.RefContent
	if err = cursor.All(context.TODO(), &Contents); err != nil {
		return nil, err
	}
	return Contents, nil
}

func (d *Daos) EnableContent(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DisableContent(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTSTATUSDISABLE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

func (d *Daos) DeleteContent(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CONTENTSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCONTENT).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
