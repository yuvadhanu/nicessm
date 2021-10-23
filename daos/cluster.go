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

//SaveCluster :""
func (d *Daos) SaveCluster(ctx *models.Context, Cluster *models.Cluster) error {
	res, err := ctx.DB.Collection(constants.COLLECTIONCLUSTER).InsertOne(ctx.CTX, Cluster)
	if err != nil {
		return err
	}
	Cluster.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

//GetSingleCluster : ""
func (d *Daos) GetSingleCluster(ctx *models.Context, code string) (*models.RefCluster, error) {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, err
	}
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"_id": id}})
	//Aggregation
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCLUSTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Clusters []models.RefCluster
	var Cluster *models.RefCluster
	if err = cursor.All(ctx.CTX, &Clusters); err != nil {
		return nil, err
	}
	if len(Clusters) > 0 {
		Cluster = &Clusters[0]
	}
	return Cluster, nil
}

//UpdateCluster : ""
func (d *Daos) UpdateCluster(ctx *models.Context, Cluster *models.Cluster) error {
	selector := bson.M{"_id": Cluster.ID}
	t := time.Now()
	update := models.Updated{}
	update.On = &t
	update.By = constants.SYSTEM
	updateInterface := bson.M{"$set": Cluster, "$push": bson.M{"updated": update}}
	_, err := ctx.DB.Collection(constants.COLLECTIONCLUSTER).UpdateOne(ctx.CTX, selector, updateInterface)
	if err != nil {
		fmt.Println("Not changed", err.Error())
		return err
	}
	return err
}

//FilterCluster : ""
func (d *Daos) FilterCluster(ctx *models.Context, Clusterfilter *models.ClusterFilter, pagination *models.Pagination) ([]models.RefCluster, error) {
	mainPipeline := []bson.M{}
	query := []bson.M{}
	if Clusterfilter != nil {

		if len(Clusterfilter.ActiveStatus) > 0 {
			query = append(query, bson.M{"activeStatus": bson.M{"$in": Clusterfilter.ActiveStatus}})
		}
		if len(Clusterfilter.Status) > 0 {
			query = append(query, bson.M{"status": bson.M{"$in": Clusterfilter.Status}})
		}
		if len(Clusterfilter.GramPanchayat) > 0 {
			query = append(query, bson.M{"gramPanchayat": bson.M{"$in": Clusterfilter.GramPanchayat}})
		}
		if len(Clusterfilter.Village) > 0 {
			query = append(query, bson.M{"village": bson.M{"$in": Clusterfilter.Village}})
		}
		//Regex
		if Clusterfilter.Regex.Name != "" {
			query = append(query, bson.M{"name": primitive.Regex{Pattern: Clusterfilter.Regex.Name, Options: "xi"}})
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
		totalCount, err := ctx.DB.Collection(constants.COLLECTIONCLUSTER).CountDocuments(ctx.CTX, func() bson.M {
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
	//lookup
	mainPipeline = append(mainPipeline, d.CommonLookupArrayOutput(constants.COLLECTIONVILLAGE, "village", "_id", "ref.village", "ref.village")...)

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Cluster query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONCLUSTER).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var Clusters []models.RefCluster
	if err = cursor.All(context.TODO(), &Clusters); err != nil {
		return nil, err
	}
	return Clusters, nil
}

//EnableCluster :""
func (d *Daos) EnableCluster(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CLUSTERSTATUSACTIVE, "activeStatus": true}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCLUSTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DisableCluster :""
func (d *Daos) DisableCluster(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CLUSTERSTATUSDISABLED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCLUSTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}

//DeleteCluster :""
func (d *Daos) DeleteCluster(ctx *models.Context, code string) error {
	id, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return err
	}
	query := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": constants.CLUSTERSTATUSDELETED, "activeStatus": false}}
	_, err = ctx.DB.Collection(constants.COLLECTIONCLUSTER).UpdateOne(ctx.CTX, query, update)
	if err != nil {
		return errors.New("Not Changed" + err.Error())
	}
	return err
}
