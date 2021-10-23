package daos

import "go.mongodb.org/mongo-driver/bson"

// CommonLookup :
func (d *Daos) CommonLookup(from string, localField string, foreignField string, as string, addField string) []bson.M {
	var Lookups []bson.M
	Lookups = append(Lookups,
		bson.M{"$lookup": bson.M{
			"from":         from,
			"localField":   localField,
			"foreignField": foreignField,
			"as":           as},
		},
		bson.M{"$addFields": bson.M{
			addField: bson.M{
				"$arrayElemAt": []interface{}{"$" + as, 0},
			},
		},
		})
	return Lookups
}

// CommonLookupArrayOutput :
func (d *Daos) CommonLookupArrayOutput(from string, localField string, foreignField string, as string, addField string) []bson.M {
	var Lookups []bson.M
	Lookups = append(Lookups,
		bson.M{"$lookup": bson.M{
			"from":         from,
			"localField":   localField,
			"foreignField": foreignField,
			"as":           as},
		},
	)
	return Lookups
}

// CommonLookup :
func (d *Daos) CommonLookupAdvanced(from string, let bson.M, pipeline []bson.M, as string, addField string) []bson.M {
	var Lookups []bson.M
	Lookups = append(Lookups,
		bson.M{"$lookup": bson.M{
			"from":     from,
			"let":      let,
			"pipeline": pipeline,
			"as":       as},
		},
		bson.M{"$addFields": bson.M{
			addField: bson.M{
				"$arrayElemAt": []interface{}{"$" + as, 0},
			},
		},
		})
	return Lookups
}

// CommonLookup :
func (d *Daos) CommonLookupAdvancedArray(from string, let bson.M, pipeline []bson.M, as string, addField string) []bson.M {
	var Lookups []bson.M
	Lookups = append(Lookups,
		bson.M{"$lookup": bson.M{
			"from":     from,
			"let":      let,
			"pipeline": pipeline,
			"as":       as},
		})

	return Lookups
}
