package daos

import (
	"errors"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveACLUserTypeTabMultiple :""
func (d *Daos) SaveACLUserTypeTabMultiple(ctx *models.Context, modules []models.ACLUserTypeTab) error {
	for _, v := range modules {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"tabId": v.TabID, "userTypeId": v.UserTypeID}
		updateData := bson.M{"$set": v}
		if _, err := ctx.DB.Collection(constants.COLLECTIONACLUSERTYPETAB).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

//GetSingleUserTypeTabAccess : ""
func (d *Daos) GetSingleUserTypeTabAccess(ctx *models.Context, userTypeID, moduleID string) (*models.UserTypeTabAccess, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": userTypeID}})
	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{"as": "module", "from": "aclmastermodules", "let": bson.M{"userTypeId": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$status", "Active"}}, bson.M{"$eq": []string{"$uniqueId", moduleID}}}}}},
				bson.M{"$lookup": bson.M{"as": "tabs", "from": "aclmastertabs", "let": bson.M{"moduleId": "$uniqueId", "userTypeId": "$$userTypeId"},
					"pipeline": []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$moduleId", "$$moduleId"}}}}}},
						bson.M{"$lookup": bson.M{
							"from": "aclmasterusetypetabs",
							"as":   "access",
							"let":  bson.M{"moduleId": "$$moduleId", "userTypeId": "$$userTypeId", "tabId": "$uniqueId"},
							"pipeline": []bson.M{
								bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
									bson.M{"$eq": []string{"$moduleId", "$$moduleId"}},
									bson.M{"$eq": []string{"$userTypeId", "$$userTypeId"}},
									bson.M{"$eq": []string{"$tabId", "$$tabId"}},
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
	d.Shared.BsonToJSONPrintTag("Tab query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var tabs []models.UserTypeTabAccess
	var tab *models.UserTypeTabAccess
	if err = cursor.All(ctx.CTX, &tabs); err != nil {
		return nil, err
	}
	if len(tabs) > 0 {
		tab = &tabs[0]
	}
	return tab, nil
}
