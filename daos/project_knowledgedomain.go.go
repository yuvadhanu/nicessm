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

//SaveProjectKnowledgeDomain :""
func (d *Daos) SaveProjectKnowledgeDomain(ctx *models.Context, domain *models.ProjectKnowledgeDomain) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN).InsertOne(ctx.CTX, domain)
	if err != nil {
		return err
	}
	domain.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

//SaveMultipleProjectKnowledgeDomain :""
func (d *Daos) SaveMultipleProjectKnowledgeDomain(ctx *models.Context, projectKD []models.ProjectKnowledgeDomain) error {
	var temProjectKD []interface{}
	for _, v := range projectKD {
		temProjectKD = append(temProjectKD, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN).InsertMany(ctx.CTX, temProjectKD)
	return err
}

//UpdateProjectKnowledgeDomain : ""
func (d *Daos) UpdateProjectKnowledgeDomain(ctx *models.Context, domain *models.ProjectKnowledgeDomain) error {

	selector := bson.M{"_id": domain.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": domain}
	_, err := ctx.DB.Collection(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//EnableProjectKnowledgeDomain :""
func (d *Daos) EnableProjectKnowledgeDomain(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTKNOWLEDGEDOMAINSTATUSACTIVE}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableProjectKnowledgeDomain :""
func (d *Daos) DisableProjectKnowledgeDomain(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTKNOWLEDGEDOMAINSTATUSDISABLED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteProjectKnowledgeDomain :""
func (d *Daos) DeleteProjectKnowledgeDomain(ctx *models.Context, UniqueID string) error {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.PROJECTKNOWLEDGEDOMAINSTATUSDELETED}}
	_, err = ctx.DB.Collection(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//GetSingleProjectKnowledgeDomain : ""
func (d *Daos) GetSingleProjectKnowledgeDomain(ctx *models.Context, UniqueID string) (*models.RefProjectKnowledgeDomain, error) {
	id, err := primitive.ObjectIDFromHex(UniqueID)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "project", "_id", "ref.project", "ref.project")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)

	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ProjectKnowledgeDomains []models.RefProjectKnowledgeDomain
	var ProjectKnowledgeDomain *models.RefProjectKnowledgeDomain
	if err = cursor.All(ctx.CTX, &ProjectKnowledgeDomains); err != nil {
		return nil, err
	}
	if len(ProjectKnowledgeDomains) > 0 {
		ProjectKnowledgeDomain = &ProjectKnowledgeDomains[0]
	}
	return ProjectKnowledgeDomain, nil
}

//FilterProjectKnowledgeDomain : ""
func (d *Daos) FilterProjectKnowledgeDomain(ctx *models.Context, filter *models.ProjectKnowledgeDomainFilter, pagination *models.Pagination) ([]models.RefProjectKnowledgeDomain, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if filter != nil {
		if len(filter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": filter.Status}})
		}
		if len(filter.Project) > 0 {
			query = append(query, bson.M{"project": bson.M{"$in": filter.Project}})
		}
		if len(filter.KnowledgeDomain) > 0 {
			query = append(query, bson.M{"knowledgeDomain": bson.M{"$in": filter.KnowledgeDomain}})
		}

	}
	//Adding $match from filter
	if len(query) > 0 {
		mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"$and": query}})
	}
	if filter != nil {
		if filter.SortBy != "" {
			mainPipeline = append(mainPipeline, bson.M{"$sort": bson.M{filter.SortBy: filter.SortOrder}})
		}
	}
	//Adding pagination if necessary
	if pagination != nil {
		mainPipeline = append(mainPipeline, []bson.M{bson.M{"$skip": (pagination.PageNum - 1) * (pagination.Limit)}, bson.M{"$limit": pagination.Limit}}...)
		//Getting Total count
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN).CountDocuments(ctx.CTX, func() bson.M {
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
	// lookup
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONPROJECT, "project", "_id", "ref.project", "ref.project")...)
	mainPipeline = append(mainPipeline, d.CommonLookup(constants.COLLECTIONKNOWLEDGEDOMAIN, "knowledgeDomain", "_id", "ref.knowledgeDomain", "ref.knowledgeDomain")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("ProjectKnowledgeDomain query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONPROJECTKNOWLEDGEDOMAIN).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var ProjectKnowledgeDomains []models.RefProjectKnowledgeDomain
	if err = cursor.All(context.TODO(), &ProjectKnowledgeDomains); err != nil {
		return nil, err
	}
	return ProjectKnowledgeDomains, nil
}
