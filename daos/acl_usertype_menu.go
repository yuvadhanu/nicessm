package daos

import (
	"errors"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//SaveACLUserTypeMenuMultiple :""
func (d *Daos) SaveACLUserTypeMenuMultiple(ctx *models.Context, modules []models.ACLUserTypeMenu) error {
	for _, v := range modules {
		opts := options.Update().SetUpsert(true)
		updateQuery := bson.M{"menuId": v.MenuID, "userTypeId": v.UserTypeID}
		updateData := bson.M{"$set": v}
		if _, err := ctx.DB.Collection(constants.COLLECTIONACLUSERTYPEMENU).UpdateOne(ctx.CTX, updateQuery, updateData, opts); err != nil {
			return errors.New("Error in updating log - " + err.Error())
		}
	}
	return nil
}

//GetSingleUserTypeMenuAccess : ""
func (d *Daos) GetSingleUserTypeMenuAccess(ctx *models.Context, userTypeID, moduleID string) (*models.UserTypeMenuAccess, error) {
	mainPipeline := []bson.M{}
	mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": userTypeID}})
	mainPipeline = append(mainPipeline,
		bson.M{"$lookup": bson.M{"as": "module", "from": "aclmastermodules", "let": bson.M{"userTypeId": "$uniqueId"},
			"pipeline": []bson.M{
				bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$status", "Active"}}, bson.M{"$eq": []string{"$uniqueId", moduleID}}}}}},
				bson.M{"$lookup": bson.M{"as": "menus", "from": "aclmastermenus", "let": bson.M{"moduleId": "$uniqueId", "userTypeId": "$$userTypeId"},
					"pipeline": []bson.M{
						bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{bson.M{"$eq": []string{"$moduleId", "$$moduleId"}}}}}},
						bson.M{"$lookup": bson.M{
							"from": "aclmasterusetypemenus",
							"as":   "access",
							"let":  bson.M{"moduleId": "$$moduleId", "userTypeId": "$$userTypeId", "menuId": "$uniqueId"},
							"pipeline": []bson.M{
								bson.M{"$match": bson.M{"$expr": bson.M{"$and": []bson.M{
									bson.M{"$eq": []string{"$moduleId", "$$moduleId"}},
									bson.M{"$eq": []string{"$userTypeId", "$$userTypeId"}},
									bson.M{"$eq": []string{"$menuId", "$$menuId"}},
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
	d.Shared.BsonToJSONPrintTag("Menu query =>", mainPipeline)
	cursor, err := ctx.DB.Collection(constants.COLLECTIONDUSERTYPE).Aggregate(ctx.CTX, mainPipeline, nil)
	if err != nil {
		return nil, err
	}
	var menus []models.UserTypeMenuAccess
	var menu *models.UserTypeMenuAccess
	if err = cursor.All(ctx.CTX, &menus); err != nil {
		return nil, err
	}
	if len(menus) > 0 {
		menu = &menus[0]
	}
	return menu, nil
}
