package daos

import (
	"errors"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveACLUserTypeFeatureMultiple :""
func (d *Daos) SaveACLUserTypeFeatureMultiple(ctx *models.Context, modules []models.ACLUserTypeFeature) error {
	for _, v := range modules {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"featureId": v.FeatureID, "userTypeId": v.UserTypeID}
		updateData := bson.M{"$set": v}
		if _, err := ctx.DB.Collection(constants.COLLECTIONACLUSERTYPEFEATURE).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

//GetSingleUserTypeFeatureAccess : ""
func (d *Daos) GetSingleUserTypeFeatureAccess(ctx *models.Context, userTypeID, moduleID string) (*models.UserTypeFeatureAccess, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": userTypeID}})
	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{"as": "module", "from": "aclmastermodules", "let": bson.M{"userTypeId": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$status", "Active"}}, bson.M{"$eq": []string{"$uniqueId", moduleID}}}}}},
				bson.M{"$lookup": bson.M{"as": "features", "from": "aclmasterfeatures", "let": bson.M{"moduleId": "$uniqueId", "userTypeId": "$$userTypeId"},
					"pipeline": []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$moduleId", "$$moduleId"}}}}}},
						bson.M{"$lookup": bson.M{
							"from": "aclmasterusetypefeatures",
							"as":   "access",
							"let":  bson.M{"moduleId": "$$moduleId", "userTypeId": "$$userTypeId", "featureId": "$uniqueId"},
							"pipeline": []bson.M{
								bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
									bson.M{"$eq": []string{"$moduleId", "$$moduleId"}},
									bson.M{"$eq": []string{"$userTypeId", "$$userTypeId"}},
									bson.M{"$eq": []string{"$featureId", "$$featureId"}},
								}},
								}},
							},
						}},
						bson.M{"$addFields": bson.M{"access": bson.M{"$arrayElemAt": []interface{}{"$access", 0}}}},
					}}},
			},
		}},
	)
	mainPipeline = append(mainPipeline, bson.M{"$addFields": bson.M{"module": bson.M{"$arrayElemAt": []interface{}{"$module", 0}}}})

	//Aggregation
	d.Shared.BsonToJSONPrintTag("Feature query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var features []models.UserTypeFeatureAccess
	var feature *models.UserTypeFeatureAccess
	if err = cursor.All(ctx.CTX, &features); err != nil {
		return nil, err
	}
	if len(features) > 0 {
		feature = &features[0]
	}
	return feature, nil
}
